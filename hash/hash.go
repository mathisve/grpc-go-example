package hash

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type Server struct {}

func (s *Server) Sha1Message(ctx context.Context, in *Message) (*Message, error) {
	h := sha1.New()
	h.Write([]byte(in.Body))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return &Message{Body: hash}, nil
}

func (s *Server) Sha265Message(ctx context.Context, in *Message) (*Message, error) {
	h := sha256.New()
	h.Write([]byte(in.Body))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return &Message{Body: hash}, nil
}

func (s *Server) Sha512Message(ctx context.Context, in *Message) (*Message, error) {
	h := sha512.New()
	h.Write([]byte(in.Body))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return &Message{Body: hash}, nil
}