package main

import (
	api "example/kitex_gen/api/echo"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
)

func main() {
	port := "9988"
	address, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		log.Panicf("Bind port:%v failed, error: %v", port, err)
	}
	svr := api.NewServer(new(EchoImpl), server.WithServiceAddr(address))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
