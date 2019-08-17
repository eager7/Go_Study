package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	r := InitRedis("127.0.0.1:6379", "", 2)
	s, e := r.Get("test").Result()
	fmt.Println("s:", s, "e:", e)

	n, err := r.Get("pct").Int()
	if err != nil && err.Error() != "redis: nil" {
		fmt.Println("err:", err)
		t.Fatal(err)
	}
	t.Log(n)
	n++
	r.Set("pct", n, time.Second*2)
	for i := 0; i < 3; i++ {
		t.Log(r.Get("pct"))
		time.Sleep(time.Second)
	}
	t.Log(r.Get("pct"))

}

func TestSet(t *testing.T) {
	r := InitRedis("127.0.0.1:6379", "", 2)
	r.HSet("key", "1", "1")//往集合里添加数据，key是集合最外层key，filed是集合内部record的key，value是record的value
	if r.HGet("key", "1").Val() != "1" {
		t.Fatal("must be 1")
	}
	r.HSet("key", "1", "2")
	if r.HGet("key", "1").Val() != "2" {
		t.Fatal("must be 1")
	}
	if err := r.HDel("key", "1").Err();err!=nil{
		t.Fatal(err)
	}
	fmt.Println("ret:", r.HGet("key", "1").Val())
}