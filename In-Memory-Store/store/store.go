package store

import "time"

type Storage interface {
	Set(key string, value string, ttl time.Duration)
	Get(key string) (string, bool)
	Delete(key string)
	Clear()
}
