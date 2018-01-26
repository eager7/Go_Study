package MyDict

import (
	"testing"
	"github.com/eager7/go/mLog"
)

func TestDict(t *testing.T) {
	dict := New()

	dict.Add(1, 2)
	if value, err := dict.Query(1);err != nil{
		t.Error("Query Failed")
	}else {
		mLog.Debug.Println(value)
	}

	dict.Add("a", "b")
	if value, err := dict.Query("a");err != nil{
		t.Error("Query Failed")
	}else {
		mLog.Debug.Println(value)
	}

	dict.Add(3, "abc")
	if value, err := dict.Query(3);err != nil{
		t.Error("Query Failed")
	}else {
		mLog.Debug.Println(value)
	}

	dict.Update(3, "abcd")
	if value, err := dict.Query(3);err != nil{
		t.Error("Query Failed")
	}else {
		mLog.Debug.Println(value)
	}

	if value, err := dict.Query("b");err != nil{
		mLog.Error.Println(err.Error())
	}else {
		mLog.Debug.Println(value)
	}
}

func BenchmarkDict(b *testing.B) {
	dict := New()
	for i:=0;i<100;i++{
		dict.Add(i, i+1)
	}
}