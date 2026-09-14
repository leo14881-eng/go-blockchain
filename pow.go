package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// 难度：要求哈希前面有几个 0。先设 4，想变难就调大
const Difficulty = 4

// ProofOfWork 结构体：绑定一个待挖的区块
type ProofOfWork struct {
	block *Block // 要挖的那个块（小写=私有，仅本包用）
}

// NewProofOfWork：为某个区块创建一个 PoW 挖矿器
func NewProofOfWork(b *Block) *ProofOfWork {
	return &ProofOfWork{block: b}
}

// prepareData：把区块字段 + 当前 nonce 拼成一个字符串，准备算哈希
// 注意 nonce 是变量，每次试不同的值
func (pow *ProofOfWork) prepareData(nonce int) string {
	return fmt.Sprintf("%d%s%s%d",
		pow.block.Timestamp,
		pow.block.Data,
		pow.block.PrevHash,
		nonce, // 关键：nonce 参与哈希计算
	)
}

// Run：挖矿主循环。不停试 nonce，直到哈希前 Difficulty 位都是 0
// 返回：挖到的 nonce 和对应的哈希
func (pow *ProofOfWork) Run() (int, string) {
	// 目标前缀：Difficulty 个 "0" 组成的字符串，比如 "0000"
	target := strings.Repeat("0", Difficulty)

	nonce := 0
	var hash string

	for { // 无限循环（Go 里 for 不写条件就是 while(true)）
		data := pow.prepareData(nonce)     // 拼数据
		sum := sha256.Sum256([]byte(data)) // 算哈希
		hash = fmt.Sprintf("%x", sum)      // 转十六进制字符串

		// 检查哈希是否以 target 开头（前 Difficulty 位是不是都是0）
		if strings.HasPrefix(hash, target) {
			break // 达标！跳出循环
		}
		nonce++ // 不达标，换下一个 nonce 继续试
	}

	return nonce, hash
}

// Validate:验证这个块的哈希是否合法
// 用块自己存的 Nonce 重新算一次哈希,看是否和存的 Hash 一致、且达到难度要求
func (pow *ProofOfWork) Validate() bool {
	// 用块里已有的 Nonce 重新拼数据、算哈希
	data := pow.prepareData(pow.block.Nonce)
	sum := sha256.Sum256([]byte(data))
	hash := fmt.Sprintf("%x", sum)

	target := strings.Repeat("0", Difficulty)

	// 两个条件都要满足:
	// ① 重算的哈希 == 块里存的哈希(内容没被改)
	// ② 哈希达到难度要求(前 Difficulty 位是0)
	return hash == pow.block.Hash && strings.HasPrefix(hash, target)
}
