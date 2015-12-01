package main

import "fmt"

func main(){
    fmt.Println("Hello\n")
    
    var grade string = "B"
    var marks int = 90

    switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 70,60,50: grade = "C"
        default:grade = "D"
    }

    switch {
    case grade == "A":
        println("your grade is A, congraulation\n")
    case grade == "B":
        println("your grade is B, good job\n")
    case grade == "C", grade == "D":
        println("your grade is C,D, come on\n")
    default:
        println("your grade is F..., not good\n")
    }

}
