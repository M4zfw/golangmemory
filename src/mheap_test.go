package src

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_init(t *testing.T) {
	Mheap.Init()
	fmt.Println(os.Getpid(), Mheap)
	for {
		time.Sleep(time.Hour)
	}
}
