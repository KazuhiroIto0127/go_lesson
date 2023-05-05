package main

import "fmt"

type Server struct {
	Host string
	Port int
	TLS bool
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTLS(tls bool) ServerOption {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(opts ...ServerOption) *Server {
	// デフォルト
	server := &Server{
		Host: "localhost",
		Port: 8080,
		TLS: false,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func main() {
	server1 := NewServer()
	fmt.Println(server1)

	server2 := NewServer(WithHost("example.com"), WithPort(443), WithTLS(true))
	fmt.Println(server2)
}
