package grpcsrv

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	notify chan error
}

// New creates new GRPC server.
func New(port string) *Server {
	s := &Server{
		grpc.NewServer(),
		make(chan error, 1),
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Can't run grpc server:", err)
	}

	go func() {
		s.notify <- s.Serve(l)
		close(s.notify)
	}()

	return s
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

// Stop gracefully stops the GRPC server.
func (s *Server) Stop() {
	s.GracefulStop()
}
