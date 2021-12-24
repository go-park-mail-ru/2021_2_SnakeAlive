package repository

import "sync"

type ShareLinkCacheStorage struct {
	storage map[string]int
	mu      *sync.Mutex
}

var links = ShareLinkCacheStorage{
	map[string]int{},
	&sync.Mutex{},
}
