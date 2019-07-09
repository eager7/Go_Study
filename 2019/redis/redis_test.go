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
