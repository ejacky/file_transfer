package core

import (
	"bytes"
	"context"
	"crypto/sha1"
	"file_transfer/messaging"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	_ "google.golang.org/grpc/encoding/gzip"
)

type ServerGRPC struct {
	logger      zerolog.Logger
	server      *grpc.Server
	port        int
	certificate string
	key         string
	filePath    chan string

	destPath string

	fileInfo *messaging.FileInfo

	finishedChunks  map[string][]uint64    // [uploadId][]chunk
	finishedFiles   map[string]interface{} // [hash]struct{}{}
	hashMapUploadId map[string]string      // [hash][uploadId]
}

type ServerGRPCConfig struct {
	Certificate string
	Key         string
	Port        int
	FilePath    chan string
}

func NewServerGRPC(cfg ServerGRPCConfig) (s ServerGRPC, err error) {
	s.logger = zerolog.New(os.Stdout).
		With().
		Str("from", "server").
		Logger()

	if cfg.Port == 0 {
		err = errors.Errorf("Port must be specified")
		return
	}

	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	s.filePath = cfg.FilePath

	s.destPath = "dest"
	s.finishedChunks = make(map[string][]uint64)
	s.finishedFiles = make(map[string]interface{})
	s.hashMapUploadId = make(map[string]string)

	return
}

func (s *ServerGRPC) Listen() (err error) {
	var (
		listener  net.Listener
		grpcOpts  = []grpc.ServerOption{}
		grpcCreds credentials.TransportCredentials
	)

	fmt.Println("grpc server address, ", "localhost:"+strconv.Itoa(s.port))
	listener, err = net.Listen("tcp", "localhost:"+strconv.Itoa(s.port))

	if err != nil {
		err = errors.Wrapf(err,
			"failed to listen on port %d",
			s.port)
		return
	}

	if s.certificate != "" && s.key != "" {
		grpcCreds, err = credentials.NewServerTLSFromFile(
			s.certificate, s.key)
		if err != nil {
			err = errors.Wrapf(err,
				"failed to create tls grpc server using cert %s and key %s",
				s.certificate, s.key)
			return
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	s.server = grpc.NewServer(grpcOpts...)
	messaging.RegisterGuploadServiceServer(s.server, s)

	err = s.server.Serve(listener)
	if err != nil {
		err = errors.Wrapf(err, "errored listening for grpc connections")
		return
	}

	return
}

func (s *ServerGRPC) Init(ctx context.Context, req *messaging.InitReq) (ack *messaging.InitAck, err error) {

	var (
		chunkExists []uint64
		uploadId    string
		ok          bool
	)

	if _, ok = s.finishedFiles[req.FileHash]; ok {
		ack = &messaging.InitAck{
			Code: 1006,
			Msg:  "file exist",
		}

		return
	}

	if uploadId, ok = s.hashMapUploadId[req.FileHash]; ok {
		chunkExists = s.finishedChunks[uploadId]
	} else {
		uploadId = fmt.Sprintf("%x", time.Now().UnixNano())
	}

	s.fileInfo = &messaging.FileInfo{
		FileHash:    req.FileHash,
		ChunkCount:  uint64(math.Ceil(float64(req.FileSize) / (5 * 1024 * 1024))),
		FileSize:    req.FileSize,
		UploadID:    uploadId,
		ChunkSize:   5 * 1024 * 1024, // 5MB
		ChunkExists: chunkExists,
	}

	ack = &messaging.InitAck{
		Data: s.fileInfo,
	}

	return ack, nil
}

func (s *ServerGRPC) Upload(stream messaging.GuploadService_UploadServer) (err error) {
	data := bytes.Buffer{}
	rand.Seed(time.Now().UnixNano())

	req, err := stream.Recv()
	if err != nil {

		err = errors.Wrapf(err,
			"failed unexpectadely while reading chunks from stream")
		return err
	}
	// todo check upload folder exist
	err = os.Mkdir(s.fileInfo.UploadID, 0755)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			fmt.Printf("%v\n", err)
			return
		}
	}

	filePath := fmt.Sprintf("%s/%d", s.fileInfo.UploadID, req.GetInfo().ChunkIndex)
	log.Println(filePath)

	file, err := os.Create(filePath)
	defer file.Close()
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {

				goto END
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return err
		}

		chunk := req.GetContent()
		//fmt.Printf("%b\t%s", chunk,string(chunk))
		_, err = data.Write(chunk)
		if err != nil {
			fmt.Println("byte buffer write data error")
		}

		if err != nil {
			return fmt.Errorf("cannot create image file: %w", err)
		}

		_, err = data.WriteTo(file)
		if err != nil {
			log.Fatal(err)
			return fmt.Errorf("cannot write image to file: %w", err)
		}
	}

	s.logger.Info().Msg("upload received")

END:
	err = stream.SendAndClose(&messaging.UploadStatus{
		Message: "Upload received with success",
		Code:    messaging.UploadStatusCode_Ok,
	})

	if err != nil {
		err = errors.Wrapf(err,
			"failed to send status code")
		return
	}

	file.Close()
	chunkHash := calcSha1(filePath)
	//chunkHash := calcSha1ByFile(file)
	if chunkHash != req.GetInfo().CheckHash {
		fmt.Printf("%d check hash failure. origin hash:%s, calcSha1:%s\n", req.GetInfo().ChunkIndex, req.GetInfo().CheckHash, chunkHash)
		return
	}

	s.finishedChunks[s.fileInfo.UploadID] = append(s.finishedChunks[s.fileInfo.UploadID], req.GetInfo().ChunkIndex)

	fmt.Printf("upload received success, uploadId:%s, index:%d\n", req.GetInfo().UploadID, req.GetInfo().ChunkIndex)

	return
}

func (s *ServerGRPC) Complete(ctx context.Context, req *messaging.CompleteReq) (ack *messaging.CompleteAck, err error) {
	// 合并文件
	err = os.MkdirAll(s.destPath, 0755)
	if err != nil {
		fmt.Println("%v", err)
	}
	destFilePath := fmt.Sprintf("%s\\%s", s.destPath, req.FileName)
	_, err = os.Create(destFilePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	destFile, err := os.OpenFile(destFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("%v", err)
		return nil, err
	}

	defer destFile.Close()

	var i uint64 = 1
	chunk := make([]byte, s.fileInfo.ChunkSize)
	for ; i <= s.fileInfo.ChunkCount; i++ {
		oFile, err := os.Open(fmt.Sprintf("%s\\%d", s.fileInfo.UploadID, i))
		if err != nil {
			fmt.Printf("%v", err)
			return nil, nil
		}
		defer oFile.Close()

		for {
			n, err := oFile.Read(chunk)

			if err != nil && err != io.EOF {
				fmt.Println("%v", err)
				return nil, err
			}

			if n == 0 {
				break
			}

			_, err = destFile.Write(chunk[:n])
			if err != nil {
				fmt.Println("%v", err)
				return nil, err
			}
			destFile.Sync()
		}
	}

	fileHash := calcSha1(destFilePath)
	//fileHash := calcSha1ByFile(destFile)
	if fileHash != req.FileHash {
		ack = &messaging.CompleteAck{
			Code: -2,
			Msg:  fmt.Sprintf("%s check hash failure, originHash:%s, fileHash:%s", req.UploadID, req.FileHash, fileHash),
		}
		fmt.Printf(ack.Msg)
		return
	}

	fmt.Printf("merge file  success, uploadId:%s\n", req.UploadID)

	s.finishedFiles[req.FileHash] = struct{}{}

	s.filePath <- destFilePath

	ack = &messaging.CompleteAck{
		Code: 0,
		Msg:  fmt.Sprintf("%s upload success", req.UploadID),
	}

	// 移除 chunk 文件夹
	go func() {
		os.RemoveAll(s.fileInfo.UploadID)
	}()

	return
}

func (s *ServerGRPC) Cancel(context.Context, *messaging.InitReq) (*messaging.InitAck, error) {
	return nil, nil
}

func (s *ServerGRPC) Close() {
	if s.server != nil {
		s.server.Stop()
	}

	return
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

func calcSha1(filePath string) string {

	infile, inerr := os.Open(filePath)
	if inerr == nil {

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		return fmt.Sprintf("%x", sha1h.Sum([]byte("")))

	} else {
		fmt.Println(inerr)
		return ""
	}
}

func calcSha1ByFile(infile *os.File) string {
	infile.Seek(0, 0)

	sha1h := sha1.New()
	io.Copy(sha1h, infile)
	return fmt.Sprintf("%x", sha1h.Sum([]byte("")))
}
