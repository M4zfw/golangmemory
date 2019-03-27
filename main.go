package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {

	fmt.Println("hello gomm ", os.Getpid())

	const n = 10 * 1024 * 1024 * 1024 //10个g
	
	if e != nil {
		panic(e)
	}
	base[0] = '1'
	fmt.Println(base[0])
	for {
		time.Sleep(time.Hour)
	}
}
