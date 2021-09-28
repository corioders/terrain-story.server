package hello

import (
	"context"
	"net/http"
)

type Controller struct{}

func NewController() (*Controller, error) {
	return &Controller{}, nil
}

func (c *Controller) Hello(ctx context.Context, rw http.ResponseWriter, r *http.Request) error {
	_, err := rw.Write([]byte("Hello"))
	if err != nil {
		return err
	}

	return nil
}
