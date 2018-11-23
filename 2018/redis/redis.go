package main

import "github.com/eager7/go/log"
import (
	"github.com/eager7/go/errors"
	"github.com/gomodule/redigo/redis"
)

func main() {
	log.L.Debug("redis test")

	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	errors.CheckErrorPanic(err)

	//BOOL
	BOOL(c)
	//INTS
	INTS(c)
	//Scan

}

func BOOL(c redis.Conn) {
	log.L.Debug(c.Do("SET", "foo", 1))
	log.L.Debug(redis.Bool(c.Do("EXISTS", "foo")))
}

func INTS(c redis.Conn) {
	log.L.Notice(c.Do("SADD", "int_map", 1, 2, 3, 4, 5, 6, 7, 8, 9, 0))
	log.L.Notice(redis.Ints(c.Do("SMEMBERS", "int_map")))
}

func Scan(c redis.Conn) {
	log.L.Debug(c.Send("HMSET", "album:1", "title", "Red", "rating", 5))
	log.L.Debug(c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1))
	log.L.Debug(c.Send("HMSET", "album:3", "title", "Beat"))
	log.L.Debug(c.Send("LPUSH", "albums", "1"))
	log.L.Debug(c.Send("LPUSH", "albums", "2"))
	log.L.Debug(c.Send("LPUSH", "albums", "3"))
	values, err := redis.Values(c.Do("SORT", "albums", "BY", "album:*->rating", "GET", "album:*->title", "GET", "album:*->rating"))
	errors.CheckErrorPanic(err)
	log.L.Info(values)
	var title string
	rating := -1
	for len(values) > 0 {
		values, err = redis.Scan(values, &title, &rating)
		errors.CheckErrorPanic(err)
		log.L.Warn(rating)
		log.L.Debug(title, rating)
	}
}