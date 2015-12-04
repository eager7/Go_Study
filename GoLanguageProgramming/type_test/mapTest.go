package main

type PersonInfo struct{
    ID          string
    Name        string
    Address     string
}
func main(){
    println("this is a map test program")
    
    //var personDB map[string] PersonInfo
    personDB := make (map[string] PersonInfo)

    personDB["1"] = PersonInfo{"111", "Pan", "1003"}
    personDB["2"] = PersonInfo{"222", "Li", "1004"}

    println("finding the 1")
    person, ok := personDB["1"]
    if ok{
        println("find ok, the person info is:",person.ID, person.Name, person.Address)
    }else{
        println("can't find the person")
    }

    delete(personDB, "1")
    println("finding the 1")
    person2, ok := personDB["1"]
    if ok{
        println("find ok, the person info is:",person2.ID, person2.Name, person2.Address)
    }else{
        println("can't find the person")
    }

}
