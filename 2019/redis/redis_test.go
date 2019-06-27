package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	r := InitRedis("127.0.0.1:6379", "", 2)
	n, err := r.Get("pct").Int()
	if err != nil && err.Error() != "redis: nil" {
		fmt.Println("err:", err)
		t.Fatal(err)
	}
	t.Log(n)
	n++
	r.Set("pct", n, time.Hour*24)
}
