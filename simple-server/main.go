package main

import (
	api "example/kitex_gen/api/echo"
	"log"
)

func main() {
	//port := "9090"
	//address, err := net.ResolveTCPAddr("tcp", ":"+port)
	//if err != nil {
	//	log.Panicf("Bind port:%v failed, error: %v", port, err)
	//}
	//svr := api.NewServer(new(EchoImpl), server.WithServiceAddr(address))

	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
