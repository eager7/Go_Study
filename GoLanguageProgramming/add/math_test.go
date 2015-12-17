package mathMethod

import "testing"

func TestAdd(t *testing.T) {
    println("test add")

    ret := MyAdd(1, 2)
    println(ret)
    if ret != 3 {
        println("add err")
    }
}
