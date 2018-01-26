package storage

import (
	"testing"
	"strconv"
)

func TestStorage(t *testing.T) {
	s := New()
	if err := s.Set("pct", strconv.Itoa(30)); err != nil{
		t.Fatal(err)
	}

	if v, err := s.Get("pct"); err != nil || v != "30"{
		t.Fatal("get value err")
	}

	if err := s.Set("my", "hello world"); err != nil{
		t.Fatal(err)
	}
	if v, err := s.Get("my"); err != nil || v != "hello world"{
		t.Fatal("get value err")
	}

	if err := s.Set("bool", 1); err != nil{
		t.Fatal(err)
	}
	if v, err := s.Get("bool"); err != nil || v != strconv.Itoa(1){
		t.Fatal("get value err")
	}
}
