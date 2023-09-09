package main

import "fmt"

type Server struct {
	Host string
	Port int
}

type ServerBuilder struct {
	server *Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{server: &Server{}}
}

func (b *ServerBuilder) WithHost(host string) *ServerBuilder {
	b.server.Host = host
	return b
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.server.Port = port
	return b
}

func (b *ServerBuilder) Build() *Server {
	return b.server
}

func main() {
	server := NewServerBuilder().
		WithHost("localhost").
		WithPort(8080).
		Build()

	fmt.Printf("Server: %+v\n", server)
}
