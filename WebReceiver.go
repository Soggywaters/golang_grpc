package Personal

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
clientCaller
)

const (
	portNumber = ":8000"
	httpType = "tcp"
)

type servers struct{}

func main () {
	// Information to listen on
	listen, error := net.Listen(httpType, portNumber)

	// If statement to check for errors and disploy a fatal log if so
	if error != nil {
		log.Fatalf("Failed: %v", error)
	}

	server := grpc.NewServer()
	clientCaller.RegisterGreeterServer(server, &servers{})

	reflection.Register(server)

	if error := server.Serve(listen); error != nil {
		log.Fatalf("Failed: %v", error)
	}

}
