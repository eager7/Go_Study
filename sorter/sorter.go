package main

import (
    "fmt"
    "flag"
    "bufio"
    "io"
    "os"
    "strconv"
    "algorithm/bubblesort"
    "algorithm/qsort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main(){
    println("this is a sort test")
  
    flag.Parse()

    fmt.Printf("The flag input is wrong, num is %d, please input again\n",flag.NArg())

    fmt.Println(*infile)
    if infile != nil{
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    } else {
    	fmt.Println("please input the infile")
    	return
    }

    values,err := readValues(*infile)
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println("Read values: ", values)
    	t1 := time.Now()
    	switch *algorithm{
    		case "qsort":
    			qsort.QuickSort(values)
    		case "bubblesort":
    			bubblesort.BubbleSort(values)
    		default:
    			fmt.Println("unknow method of sort")
    	}
    	t2 := time.Now()
    	fmt.Println("sort algorithm", *algorithm, "costs", t2.Sub(t1), "to complete")
    	writeValues(values, *outfile)
    }
}

func readValues(infile string)(values []int, err error){
    file, err := os.Open(infile)
    if err != nil{
        fmt.Println("Failed to open the input file ", infile)
        return
    }
    defer file.Close()

    br := bufio.NewReader(file)

    values = make([]int ,0)

    for{
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil{
            if err1 != io.EOF{
                err = err1
            }
            break
        }
        if isPrefix{
            fmt.Println("A too long line, seems unexpected.")
            return
        }

        str := string(line)
        value, err1 := strconv.Atoi(str)
        if err1 != nil {
            err = err1
            return
        }
        values = append(values,value)
    }
    return
}

func writeValues(values []int, outfile string) error{
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create outfile:", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}