package berliner

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type CachedClient struct {
	cache  *Cache
	client scrape.Client
}

func NewCachedClient(cache *Cache) *CachedClient {
	return &CachedClient{
		cache:  cache,
		client: scrape.NewClient(),
	}
}

func (c *CachedClient) GetPost(url string) (content.Post, error) {
	cachedPost, err := c.cache.Get(url)
	if err == nil {
		return *cachedPost, nil
	}
	post, err := c.client.GetPost(url)
	if err != nil {
		return content.Post{}, err
	}
	c.cache.Put(url, &post)
	return post, nil
}

// TODO: switch to binary encoding for smaller space
// TODO: cache eviction when cache reaches max size
type Cache struct {
	db    *bolt.DB
	store *store
}

func NewCache(path string) (*Cache, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Cache{
		db:    db,
		store: NewStore(db, []byte("posts")),
	}, nil
}

func (c *Cache) Get(permalink string) (*content.Post, error) {
	post := &content.Post{}
	err := c.store.Get([]byte(permalink), post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (c *Cache) Put(permalink string, post *content.Post) error {
	err := c.store.Put([]byte(permalink), post)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Close() error {
	return c.db.Close()
}

type store struct {
	db     *bolt.DB
	bucket []byte
}

func NewStore(db *bolt.DB, bucket []byte) *store {
	return &store{
		db:     db,
		bucket: bucket,
	}
}

func (s *store) Put(key []byte, post *content.Post) error {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(post); err != nil {
		return err
	}
	return s.db.Update(func(tx *bolt.Tx) error {
		objects, err := tx.CreateBucketIfNotExists(s.bucket)
		if err != nil {
			return err
		}
		objects.Put(key, buf.Bytes())
		return nil
	})
}

func (s *store) Get(key []byte, post *content.Post) error {
	buf := bytes.NewBuffer(nil)
	err := s.db.Update(func(tx *bolt.Tx) error {
		objects := tx.Bucket(s.bucket)
		if objects == nil {
			return errors.New("post not found")
		}
		data := objects.Get(key)
		if data == nil {
			return errors.New("post not found")
		}
		buf.Write(data)
		return nil
	})

	if err != nil {
		return err
	}

	dec := gob.NewDecoder(buf)
	err = dec.Decode(post)

	return err
}
