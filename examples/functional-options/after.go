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

type ServiceOption func(*Service)

func NewService(name string, opts ...ServiceOption) *Service {
	s := &Service{
		Name: name,
	}

	// call all the service options
	for _, o := range opts {
		o(s)
	}

	return s
}

func WithEnabled(s *Service) {
	s.Enabled = true
}

func WithListener(l string) ServiceOption {
	return func(s *Service) {
		s.Listen = l
	}
}

func WithRetries(r int) ServiceOption {
	return func(s *Service) {
		s.NumRetries = r
	}
}

func WithTimeout(t time.Duration) ServiceOption {
	return func(s *Service) {
		s.Timeout = t
	}
}

func (s *Service) Run() {
	fmt.Println("Listening on", s.Listen)
}

func main() {
	s := NewService("http",
		WithEnabled,
		WithListener("127.0.0.0:8080"),
		WithRetries(10),
		WithTimeout(time.Second*3),
	)

	fmt.Printf("Service: %#v\n\n", s)
	s.Run()
}
