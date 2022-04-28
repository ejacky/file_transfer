package core

import (
	"bytes"
	"context"
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

	fileInfo *messaging.FileInfo
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

	s.fileInfo = &messaging.FileInfo{
		FileHash:    req.FileHash,
		ChunkCount:  uint64(math.Ceil(float64(req.FileSize) / (5 * 1024 * 1024))),
		FileSize:    req.FileSize,
		UploadID:    "uploadID",
		ChunkSize:   5 * 1024 * 1024, // 5MB
		ChunkExists: []uint64{},
		FileName:    req.FileName,
	}
	ack = &messaging.InitAck{
		Data: s.fileInfo,
	}

	return ack, nil
}

func (s *ServerGRPC) Complete(context.Context, *messaging.InitReq) (*messaging.InitAck, error) {
	return nil, nil
}

func (s *ServerGRPC) Cancel(context.Context, *messaging.InitReq) (*messaging.InitAck, error) {
	return nil, nil
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
		fmt.Println("%v", err)
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
	s.filePath <- filePath

	if err != nil {
		err = errors.Wrapf(err,
			"failed to send status code")
		return
	}

	fmt.Println("upload received success")

	return
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
