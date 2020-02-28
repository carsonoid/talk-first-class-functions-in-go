package main

import (
	"fmt"
	"time"
)

type Service struct {
	Name       string
	Enabled    bool
	Listen     string
	NumRetries int
	Timeout    time.Duration
}

func NewService(name string, enabled bool, listen string, numRetries int, timeout time.Duration) *Service {
	return &Service{
		Name:       name,
		Enabled:    enabled,
		Listen:     listen,
		NumRetries: numRetries,
		Timeout:    timeout,
	}
}

func (s *Service) Run() {
	fmt.Println("Listening on", s.Listen)
}

func main() {
	s := NewService("http", true, "127.0.0.0:8080", 10, time.Second*3)

	fmt.Printf("Service: %#v\n\n", s)
	s.Run()
}
