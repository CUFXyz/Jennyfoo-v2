package storage

import (
	"errors"
	"sync"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrNotWritten    = errors.New("data is not written")
	ErrEmptyInput    = errors.New("input empty")
	ErrAlreadyExists = errors.New("link already exists")
)

type Cache struct {
	Cache map[string]string
	Mu    sync.Mutex
}

func (c *Cache) WriteCache(token string, login string) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if login == "" || token == "" {
		return ErrNotFound
	}

	if _, ok := c.Cache[token]; ok {
		return ErrAlreadyExists
	}
	c.Cache[token] = login
	return nil
}

func (c *Cache) GetValue(token string) (string, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	email, ok := c.Cache[token]
	if !ok {
		return "", ErrNotFound
	}
	return email, nil
}

func NewCache() *Cache {
	var c Cache
	c.Cache = make(map[string]string)
	return &c
}
