package core

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

// ClientH2 provides the implementation of a file
// uploader that streams data via an HTTP2-enabled
// connection.
type ClientH2 struct {
	client  *http.Client
	address string
}

type ClientH2Config struct {
	RootCertificate string
	Address         string
}

func NewClientH2(cfg ClientH2Config) (c ClientH2, err error) {
	if cfg.Address == "" {
		err = errors.Errorf("Address must be non-empty")
		return
	}

	if cfg.RootCertificate == "" {
		err = errors.Errorf("RootCertificate must be specified")
		return
	}

	cert, err := ioutil.ReadFile(cfg.RootCertificate)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to read root certificate")
		return
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(cert)
	if !ok {
		err = errors.Errorf(
			"failed to root certificate %s to cert pool",
			cfg.RootCertificate)
		return
	}

	c.client = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
	}

	c.address = cfg.Address

	return
}

func (c *ClientH2) UploadFile(ctx context.Context, r *http.Request) (stats Stats, err error) {
	return Stats{}, nil

}

func (c *ClientH2) Close() {
	return
}
