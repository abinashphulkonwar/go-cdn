package storage

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	MinifyJson "github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
)

type Storage struct {
	Dir     string
	MetaDir string
	Cache   *cache.Cache
	Minify  *minify.M
}

type MetaData struct {
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
	Date          string `json:"Date"`
	CacheControl  string `json:"Cache-Control"`
}

const Meta = "Meta"
const Temp = "Temp"
const MetaDir = "meta"
const TempDir = "temp"

func New(Dir string, MetaDir string) *Storage {
	c := cache.New(5*time.Minute, 10*time.Minute)

	m := minify.New()
	m.AddFunc(MediaType.HTML, html.Minify)
	m.AddFunc(MediaType.CSS, css.Minify)
	m.AddFunc(MediaType.IMAGE, svg.Minify)
	m.AddFuncRegexp(MediaType.JS, js.Minify)
	m.AddFuncRegexp(MediaType.JSON, MinifyJson.Minify)

	return &Storage{
		Dir:     Dir,
		MetaDir: MetaDir,
		Cache:   c,
		Minify:  m,
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

func (s *Storage) WriteFile(key string, data []byte, mediaType string) error {
	media_data := data
	compress_data, err := s.Compress(data, mediaType)
	if err != nil {
		println("error: ", err.Error())
	} else {
		media_data = compress_data
	}
	err = s.Write(s.TempPath(key), media_data)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) DeleteFile(key string) error {

	err := os.Remove(key)
	if err != nil {

		isExist := os.IsNotExist(err)
		if isExist {
			return nil
		}
		return err
	}
	return nil
}

func (s *Storage) DeleteMeta(key string) error {

	err := s.DeleteFile(s.MetaPath(key))
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) DeleteTemp(key string) error {
	err := s.DeleteFile(s.TempPath(key))
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(key string) error {

	err := s.DeleteMeta(key)
	if err != nil {
		return err
	}

	err = s.DeleteTemp(key)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetFile(Key string) ([]byte, error) {
	path := s.TempPath(Key)

	cahce, isFound := s.GetDataFromCache(Key)
	if isFound {
		return cahce.([]byte), nil
	}

	data, err := s.Read(path)

	if err != nil || len(data) == 0 {
		return nil, err
	}
	s.Cache.Set(Key, data, cache.DefaultExpiration)
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
	if err != nil || len(buf) == 0 {
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
func (s *Storage) Compress(data []byte, mediaType string) ([]byte, error) {
	if mediaType == MediaType.HTML {
		return s.Minify.Bytes(MediaType.HTML, data)
	}
	if mediaType == MediaType.CSS {
		return s.Minify.Bytes(MediaType.CSS, data)

	}
	if mediaType == MediaType.IMAGE {
		return s.Minify.Bytes(MediaType.CSS, data)
	}

	if MediaType.JS.MatchString(mediaType) || MediaType.JSON.MatchString(mediaType) {
		return s.Minify.Bytes(mediaType, data)
	}

	return data, nil
}
