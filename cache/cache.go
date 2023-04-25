package cache

import (
	"sync"
	"time"
	"github.com/krishnanshagarwal112/pokedex/models"
)

// returns a new cache struct !
func NewCache(interval time.Duration) models.Cache {
	c := models.Cache{
		Caches:make(map[string]models.CacheEntry),
		Mux: &sync.Mutex{},
	}
	go ReadLoop(&c,interval)
	return c
}

func Add(cache *models.Cache,key string,val []byte){
	cache.Mux.Lock()
	defer cache.Mux.Unlock()
	
	cache.Caches[key] = models.CacheEntry{
		CreatedAt:time.Now(),
		Val:val,
	}
}

func Get(cache *models.Cache,key string)([]byte, bool){
	cache.Mux.Lock()
	defer cache.Mux.Unlock()

	val,ok := cache.Caches[key]
	return val.Val,ok
} 

func ReadLoop(cache *models.Cache, interval time.Duration){

	c := cache.Caches
	ticker := time.NewTicker(interval)

	for range ticker.C{
		cache.Mux.Lock()
		for key,val:= range c{
			if time.Since(val.CreatedAt) > interval{
				delete(c,key)
			}
		}
		cache.Mux.Unlock()
	}
}

