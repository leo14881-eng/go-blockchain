package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block 是区块结构体，相当于 Java 的一个 class
type Block struct {
	Timestamp int64  // 字段首字母大写 = public（对外可见）
	Data      string // 首字母小写 = private（仅包内可见）
	PrevHash  string
	Hash      string
}

// SetHash 是 Block 的方法。(b *Block) 叫"接收者"，相当于 Java 的 this
func (b *Block) SetHash() {
	info := fmt.Sprintf("%d%s%s", b.Timestamp, b.Data, b.PrevHash)
	sum := sha256.Sum256([]byte(info)) // 算 SHA256 哈希
	b.Hash = fmt.Sprintf("%x", sum)    // %x 转十六进制字符串
}

// NewBlock 是"构造函数"（Go 没有 new 关键字构造，靠约定俗成的函数）
func NewBlock(data string, prevHash string) *Block {
	b := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}
	b.SetHash()
	return b
}
