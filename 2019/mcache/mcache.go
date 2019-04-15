package mcache

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Cache struct {
	mem *memcache.Client
}

func Initialize(addr ...string) *Cache {
	return &Cache{mem: memcache.New(addr...)}
}

func (c *Cache) Set(key string, value []byte, expiration int32) error {
	return c.mem.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Flags:      0,
		Expiration: expiration,
	})
}

func (c *Cache) Get(key string) ([]byte, error) {
	if item, err := c.mem.Get(key); err != nil {
		return nil, err
	} else {
		return item.Value, nil
	}
}

func (c *Cache) MiddleWare() gin.HandlerFunc {
	return c.CacheHandle
}

func (c *Cache) CacheHandle(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		return
	}
	key := HttpStore(ctx)
	value, err := c.Get(key)
	if err == nil { //hit the cache
		ctx.AbortWithStatusJSON(http.StatusOK, value)
		return
	}
	//请求前
	ctx.Next() //处理请求
	//请求后
	if ctx.Writer.Status() != http.StatusOK {
		return
	}
	fmt.Println("save data into mem cache")
}

func Md5(message []byte) (tmp string) {
	md5Ctx := md5.New()
	md5Ctx.Write(message)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func HttpStore(ctx *gin.Context) string {
	bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
	urlBytes := []byte(ctx.Request.URL.String())
	return Md5(append(bodyBytes, urlBytes...))
}
