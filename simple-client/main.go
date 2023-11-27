package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/tuoaitang/kitex-demo/simple-client/kitex_gen/api"
	"github.com/tuoaitang/kitex-demo/simple-client/kitex_gen/api/echo"
)

func main() {
	client, err := echo.NewClient("hello", client.WithHostPorts("0.0.0.0:9090"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
