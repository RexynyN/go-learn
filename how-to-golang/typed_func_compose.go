package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Server struct {
	filenameTransformFunc TransformFunc
}

type TransformFunc func(string) string

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("New Filename: ", newFilename)
	return nil
}

// sha1
// prefix with GG_
// hmac
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	return hex.EncodeToString(hash[:])
}

func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {
	s := &Server{
		filenameTransformFunc: prefixFilename("FINAL_"),
	}

	s.handleRequest("breno.mp4")
}
