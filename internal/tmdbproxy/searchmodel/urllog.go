package searchmodel

import (
	"github.com/gin-gonic/gin"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cachestruct"
	"github.com/twmb/murmur3"
)

type UrlLogCache = cachestruct.Cache[uint64, UrlLog]

func (o UrlLog) New(c *gin.Context) UrlLog {
	c.ShouldBindQuery(&o)
	o.Path = c.Request.URL.Path
	o.RawQuery = c.Request.URL.RawQuery
	o.QueryChanged = false
	return o
}

func (o UrlLog) pk() uint64 {
	h := murmur3.New64()
	h.Write([]byte(o.Path + o.RawQuery))
	return h.Sum64()
}

func (o UrlLog) UpdateOne(cache *UrlLogCache) error {
	o.Counter += 1
	return cache.Set(o.pk(), o)
}

func (o UrlLog) FindOne(cache *UrlLogCache) (UrlLog, error) {
	if cache.Has(o.pk()) {
		err := cache.Get(o.pk(), &o)
		return o, err
	} else {
		return o, nil
	}
}

func (o UrlLog) DeleteMany(cache *UrlLogCache) error {
	items, err := cache.Items()
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.TmdbStatus == o.TmdbStatus {
			cache.Delete(item.pk())
		}
	}
	return nil
}
