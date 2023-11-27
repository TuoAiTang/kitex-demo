package main

import (
	"context"
	api "example/kitex_gen/api"
	"log"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	log.Printf("receive message: %s", req.Message)
	resp = &api.Response{Message: req.Message}
	return
}
