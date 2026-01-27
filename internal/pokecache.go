package internal

import "time"

type Cache struct {
	cacheEntry struct {
		createdAt time.Time
		val       []byte
	}
}
