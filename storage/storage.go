package storage

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

type Storage struct {
	Dir     string
	MetaDir string
}

const Meta = "Meta"
const Temp = "Temp"

func New(Dir string, MetaDir string) *Storage {
	return &Storage{
		Dir:     Dir,
		MetaDir: MetaDir,
	}
}

func (s *Storage) Write(key string, data []byte) error {
	err := os.WriteFile(key, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) Read(key string) ([]byte, error) {
	buf, err := os.ReadFile(key)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return buf, nil
}

func (s *Storage) WriteFile(key string, data []byte) error {
	err := s.Write(s.TempPath(key), data)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) DeleteFile() error {
	return nil
}
func (s *Storage) GetFile(Key string) ([]byte, error) {
	return s.Read(s.TempPath(Key))
}

func (s *Storage) SetMetaData(key string, data []byte) error {
	err := s.Write(s.MetaPath(key), data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMetaData(Key string) ([]byte, error) {
	buf, err := s.Read(s.MetaPath(Key))
	return buf, err
}

func (s *Storage) TempPath(key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	path := s.Dir + hashString
	return path
}
func (s *Storage) MetaPath(key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	path := s.MetaDir + hashString
	return path
}
