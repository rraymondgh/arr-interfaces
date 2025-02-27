package cachestruct

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"io"

	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	cache    *ttlcache.Cache[K, []byte]
	compress bool
}

func New[K comparable, V any](compress bool, opts ...ttlcache.Option[K, []byte]) *Cache[K, V] {
	cache := Cache[K, V]{
		cache:    ttlcache.New[K, []byte](opts...),
		compress: compress,
	}

	return &cache
}

func (c *Cache[K, V]) Has(key K) bool {
	return c.cache.Has(key)
}

func (c *Cache[K, V]) encodeJSON(value V) (*[]byte, error) {
	mbytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	if c.compress {
		var b bytes.Buffer
		w, err := flate.NewWriter(&b, flate.BestSpeed)
		if err != nil {
			return nil, err
		}
		_, err = w.Write(mbytes)
		if err != nil {
			return nil, err
		}
		w.Close()
		mbytes = b.Bytes()
	}

	return &mbytes, nil
}

func (c *Cache[K, V]) Set(key K, value V) error {
	mbytes, err := c.encodeJSON(value)
	if err != nil {
		return err
	}
	c.cache.Set(key, *mbytes, ttlcache.DefaultTTL)
	return nil
}

func (c *Cache[K, V]) decodeJSON(item *ttlcache.Item[K, []byte], value *V) error {
	var b []byte
	if c.compress {
		r := flate.NewReader(bytes.NewReader(item.Value()))
		decompressed, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		b = decompressed
	} else {
		b = item.Value()
	}

	return json.Unmarshal(b, value)
}

func (c *Cache[K, V]) Get(key K, value *V) error {
	return c.decodeJSON(c.cache.Get(key), value)
}

func (c *Cache[K, V]) Start() {
	c.cache.Start()
}

func (c *Cache[K, V]) Stop() {
	c.cache.Stop()
}

func (c *Cache[K, V]) Len() int {
	return c.cache.Len()
}

func (c *Cache[K, V]) Items() ([]V, error) {
	items := make([]V, c.Len())
	i := 0
	for _, item := range c.cache.Items() {
		err := c.decodeJSON(item, &items[i])
		if err != nil {
			return []V{}, err
		}
		i += 1
		if i >= len(items) {
			break
		}
	}
	return items, nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.cache.Delete(key)
}
