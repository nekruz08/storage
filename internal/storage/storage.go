package storage

import "github.com/nekruz08/storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string,blob []byte) (*file.File,error) {
	return file.NewFile(filename,blob)
}