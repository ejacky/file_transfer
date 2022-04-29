package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"file_transfer/messaging"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "google.golang.org/grpc/encoding/gzip"
)

// ClientGRPC provides the implementation of a file
// uploader that streams chunks via protobuf-encoded
// messages.
type ClientGRPC struct {
	logger    zerolog.Logger
	conn      *grpc.ClientConn
	Client    messaging.GuploadServiceClient
	chunkSize int
}

type ClientGRPCConfig struct {
	Address         string
	ChunkSize       int
	RootCertificate string
	Compress        bool
}

func NewClientGRPC(cfg ClientGRPCConfig) (c ClientGRPC, err error) {
	var (
		grpcOpts  = []grpc.DialOption{}
		grpcCreds credentials.TransportCredentials
	)

	if cfg.Address == "" {
		err = errors.Errorf("address must be specified")
		return
	}

	if cfg.Compress {
		grpcOpts = append(grpcOpts,
			grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")))
	}

	if cfg.RootCertificate != "" {
		grpcCreds, err = credentials.NewClientTLSFromFile(cfg.RootCertificate, "localhost")
		if err != nil {
			err = errors.Wrapf(err,
				"failed to create grpc tls client via root-cert %s",
				cfg.RootCertificate)
			return
		}

		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(grpcCreds))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	switch {
	case cfg.ChunkSize == 0:
		err = errors.Errorf("ChunkSize must be specified")
		return
	case cfg.ChunkSize > (1 << 22):
		err = errors.Errorf("ChunkSize must be < than 4MB")
		return
	default:
		c.chunkSize = cfg.ChunkSize
	}

	c.logger = zerolog.New(os.Stdout).
		With().
		Str("from", "client").
		Logger()

	c.conn, err = grpc.Dial(cfg.Address, grpcOpts...)
	fmt.Println("grpc address,", cfg.Address)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to start grpc connection with address %s",
			cfg.Address)
		return
	}

	c.Client = messaging.NewGuploadServiceClient(c.conn)

	return
}

func (c *ClientGRPC) UploadFile(ctx context.Context, r *http.Request) (stats Stats, err error) {
	var (
		writing = true
		buf     []byte
		n       int

		status *messaging.UploadStatus
	)

	stream, err := c.Client.Upload(ctx)
	if err != nil {
		err = errors.Wrap(err, "failed to create upload stream for file")
		return
	}

	index, err := strconv.Atoi(r.FormValue("index"))
	if err != nil {
		log.Fatal(err)
		return
	}
	err = stream.Send(&messaging.Chunk{
		Data: &messaging.Chunk_Info{
			Info: &messaging.ChunkInfo{
				UploadID:   r.FormValue("uploadid"),
				ChunkIndex: uint64(index),
				CheckHash:  r.FormValue("chkhash"),
			},
		},
	})
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	defer stream.CloseSend()

	stats.StartedAt = time.Now()
	buf = make([]byte, c.chunkSize)
	for writing {
		n, err = r.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil

				// 未读出字节，跳出 for 循环
				if n == 0 {
					continue
				}
			} else {
				err = errors.Wrapf(err,
					"errored while copying from file to buf")
				return
			}

		}

		err = stream.Send(&messaging.Chunk{
			Data: &messaging.Chunk_Content{
				Content: buf[:n],
			},
		})
		if err != nil {
			err = errors.Wrapf(err,
				"failed to send chunk via stream")
			return
		}
	}

	stats.FinishedAt = time.Now()

	status, err = stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to receive upstream status response")
		return
	}

	if status.Code != messaging.UploadStatusCode_Ok {
		err = errors.Errorf(
			"upload failed - msg: %s",
			status.Message)
		return
	}

	fmt.Println("grpc upload success")

	return
}

func (c *ClientGRPC) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
