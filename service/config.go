package service

import (
	"encoding/json"
	"flag"
	"net/url"
	"os"
	"strings"
)

type CacheConfiguration struct {
	TTL              int    `json:"ttl"` // in seconds
	InvalidationPath string `json:"invalidation-path"`
	Token            string `json:"token"`
}

type Config struct {
	Origin  string             `json:"origin"`
	Methods []string           `json:"methods"`
	Cache   CacheConfiguration `json:"cache"`
	Method  map[string]string
}

func Configuration() Config {
	var filename string
	flag.StringVar(&filename, "f", "", "file name")
	flag.Parse()
	if filename == "" {
		panic("config file not provided")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	cache := Config{}
	err = json.Unmarshal(data, &cache)
	if err != nil {
		panic(err)
	}

	if cache.Origin == "" {
		panic("origin not provided")
	}

	if cache.Origin[len(cache.Origin)-1:] == "/" {
		cache.Origin = cache.Origin[:len(cache.Origin)-1]
	}

	if cache.Cache.InvalidationPath != "" {

		if cache.Cache.Token == "" {
			panic("token not provided")
		}
		if isStartWith := strings.HasPrefix(cache.Cache.InvalidationPath, "/"); isStartWith {
			cache.Cache.InvalidationPath = cache.Cache.InvalidationPath[1:]
		}
	}
	_, err = url.ParseRequestURI(cache.Origin)
	if err != nil {
		panic(err)
	}

	cache.Method = make(map[string]string)
	for index := range cache.Methods {
		currentMethod := cache.Methods[index]
		if currentMethod != "" {
			cache.Method[currentMethod] = currentMethod
		}
	}

	cache.Methods = nil

	return cache
}
