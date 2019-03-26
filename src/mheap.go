package src

import "sync"

/*
	地址布局
	[spans --- 512M][bitmap --- 16G][arena --- 512G]
*/
//定义结构
type mheap struct {
	lock sync.Mutex //并发冲突

	spans []*mspan //spans区

	bitmap uintptr //bitmap区起点

	arenaStart uintptr //arena 区域
	arenaUsed  uintptr
	arenaEnd   uintptr
	
}
