package main

func main(){
    println("slice 2 test")
    mySlice := make([]int, 5, 10)
    println("len of mySlice:", len(mySlice))
    println("cap of mySlice:", cap(mySlice))
    
    for _, v := range mySlice{
        println(v)
    }

    mySlice = append(mySlice, 1, 2)
    println("len of mySlice:", len(mySlice))
    println("cap of mySlice:", cap(mySlice))
    for _, v := range mySlice{
        println(v)
    }
    
    mySlice2 := []int{8,9,10}
    mySlice = append(mySlice,mySlice2...)
    println("len of mySlice:", len(mySlice))
    println("cap of mySlice:", cap(mySlice))
    for _, v := range mySlice{
        println(v)
    }
}
