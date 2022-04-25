package core

import (
	"golang.org/x/net/context"
	"net/http"
)

// Client interface defines the methods that a client
// that desired to upload a given file (`f`) to a server
// should implement.
type Client interface {

	// UploadFile takes the contents of a file (`f`) and
	// uploads to a server.
	// The context should be used in order to cancel uploads
	// if needed or provide special metadata.
	UploadFile(ctx context.Context,  r *http.Request) (stats Stats, err error)

	// Closes releases resources associated with the
	// instantiation of the client.
	Close()
}
