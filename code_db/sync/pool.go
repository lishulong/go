package main

import (
	"bytes"
	"log"
	"sync"
)

type StrBuffer struct {
	// embed it
	Buf []byte
}

func NewStrBuffer() interface{} {
	log.Println("NewStrBuffer is called")
	return &StrBuffer{
		Buf: make([]byte, 0, 1024),
	}
}

func GetStrBuffer(pool *sync.Pool, len int) *StrBuffer {
	sb := pool.Get().(*StrBuffer)
	sb.Buf = sb.Buf[0:len]

	return sb
}

func PutBackStrBuffer(pool *sync.Pool, sb *StrBuffer) {
	// reset it first
	sb.Buf = sb.Buf[0:0]

	pool.Put(sb)
}

func main() {
	bufPool := &sync.Pool{
		New: NewStrBuffer,
	}

	strBuf := GetStrBuffer(bufPool, 100)

	b := bytes.NewBufferString("hello world!")
	b.Read(strBuf.Buf)
	log.Printf("len: %d, value: %s\n", len(strBuf.Buf), string(strBuf.Buf))

	PutBackStrBuffer(bufPool, strBuf)

	// use the same underlying byte slice
	strBuf = GetStrBuffer(bufPool, 50)
	log.Printf("len: %d, value: %s\n", len(strBuf.Buf), string(strBuf.Buf))

	b = bytes.NewBufferString("goodbye world!")
	b.Read(strBuf.Buf)
	log.Printf("len: %d, value: %s\n", len(strBuf.Buf), string(strBuf.Buf))
}
