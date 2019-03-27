package src

import (
	"sync"
	"unsafe"
)

/*
	地址布局 golang
	[spans --- 512M][bitmap --- 16G][arena --- 512G]
	spans 8k(一页大小):8B(指针大小) 所以：1:1024   bitmap 2bit表示arena一个指针 所以： 1:32

	这里的模拟 arena区用5.12G
*/
//定义结构

const (
	_Pageshift    = 13  //32k
	_Pagetypesize = 127 //mspan多包含127个8k页 即1m

	//PAGESIZE 页大小
	PAGESIZE = 2 * 4 * 1024

	//AREANSIZE  所有可用与存数据的空间 arena域 模拟使用5.12G
	AREANSIZE = 512 * 1024 * 1024 * 1024 / 100

	//SPANSSIZE spans区域大小
	SPANSSIZE = AREANSIZE / 1024

	//BITMAPSIZE bitmap区域
	BITMAPSIZE = AREANSIZE / 32

	//ALLSIZE  总需空间（总分配空间要为页面大小的倍数）
	ALLSIZE = AREANSIZE + SPANSSIZE + BITMAPSIZE
)

//Mheap 全局唯一
var Mheap mheap

type mheap struct {
	lock       sync.Mutex //并发冲突
	allstart   uintptr    //堆区起点
	spans      []*mspan   //spans区
	bitmap     uintptr    //bitmap区起点
	arenaStart uintptr    //arena 区域
	arenaUsed  uintptr
	arenaEnd   uintptr

	//小对象
	free [_Pagetypesize]mSpanList //退回来的空间尽量重复利用
	busy [_Pagetypesize]mSpanList //正在用的，或者已经分配给了mcache的
	//大对象
	freelarge uintptr
	busylarge mSpanList
}

//Init 初始化
func (m *mheap) Init() {
	Mheap = mheap{}
	//申请一个连续的地址
	b, e := initheap(Round(ALLSIZE, PAGESIZE))
	if e != nil {
		panic(e)
	}
	//拿到地址
	Mheap.allstart = uintptr(unsafe.Pointer(&b))
}

//Round n向上取a的倍数 如  (1021,1024) 得到 1024  (1025,1024)  得到2048
func Round(n, a int) int {
	return (n + a - 1) &^ (a - 1)
}
