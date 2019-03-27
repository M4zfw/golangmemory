package src

import "syscall"

//申请一段连续地址
func initheap(n int) ([]byte, error) {
	base, e := syscall.Mmap(-1, 0, n, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	return base, e
}

