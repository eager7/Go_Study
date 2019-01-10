package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"sync"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			wg.Add(1)
			go func() {
				defer wg.Done()
				ret, err := runCmd(fmt.Sprintf("cd %s; git pull", file.Name()))
				if err != nil {
					fmt.Println("run cmd error:", err)
				}
				fmt.Println(ret)
			}()
		}
	}
	wg.Wait()
	fmt.Println("done...")
}

func runCmd(shell string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", shell)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("exec ", cmd.Args, "failed, ", err.Error())
		fmt.Println(string(out))
		return "", err
	}
	return string(out), err
}
