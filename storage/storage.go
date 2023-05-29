package storage

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type Storage struct {
	Dir     string
	MetaDir string
	Cache   *cache.Cache
}

type MetaData struct {
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
	Date          string `json:"Date"`
	CacheControl  string `json:"Cache-Control"`
}

const Meta = "Meta"
const Temp = "Temp"

func New(Dir string, MetaDir string) *Storage {
	c := cache.New(5*time.Minute, 10*time.Minute)

	return &Storage{
		Dir:     Dir,
		MetaDir: MetaDir,
		Cache:   c,
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
	path := s.TempPath(Key)

	cacheKey := Temp + path

	cahce, isFound := s.GetDataFromCache(cacheKey)
	if isFound {
		return cahce.([]byte), nil
	}

	data, err := s.Read(path)

	if err != nil {
		return nil, err
	}
	s.Cache.Set(cacheKey, data, cache.DefaultExpiration)
	return data, nil
}

func (s *Storage) SetMetaData(key string, data []byte) error {
	err := s.Write(s.MetaPath(key), data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMetaData(Key string) (MetaData, bool, error) {
	path := s.MetaPath(Key)
	metaKey := Meta + path
	metaData := MetaData{}
	data, isFound := s.GetDataFromCache(metaKey)
	if isFound {
		return data.(MetaData), true, nil
	}
	buf, err := s.Read(path)
	if err != nil {
		return metaData, false, err
	}

	err = json.Unmarshal(buf, &metaData)

	if err != nil {
		return metaData, false, err
	}

	s.Cache.Set(metaKey, metaData, cache.DefaultExpiration)

	return metaData, true, err
}

func (s *Storage) GetDataFromCache(key string) (interface{}, bool) {
	data, found := s.Cache.Get(key)
	if found {
		return data, true
	}
	return data, false
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
