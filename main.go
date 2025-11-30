package main

import (
	"bufio"
	"io"
	"sync"
)

type store struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewStore() *store {
	return &store{
		data: make(map[string]string),
	}

}

func (s *store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

func (s *store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]
	delete(s.data, key)
	return exists
}

func (s *store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.Unlock()
	_, ok := s.data[key]
	return ok
}

type RESPReader struct {
	reader *bufio.Reader
}

func NewRESPReader(r io.Reader) *RESPReader {
	return &RESPReader{reader: bufio.NewReader(r)}
}

func (r *RESPReader) ReadCommand() ([]string, error) {
	line, err := r.reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

}
