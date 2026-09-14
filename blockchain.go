package main

import "fmt"

// Blockchain 结构体：代表一整条链
// Java 对照：
//
//	public class Blockchain {
//	    private List<Block> blocks;
//	}
type Blockchain struct {
	Blocks []*Block // 一个装 *Block 的切片。Java: List<Block> blocks;
}

// NewBlockchain：创建一条新链，自动放入创世区块
// Java 对照：工厂方法 static Blockchain newBlockchain()
func NewBlockchain() *Blockchain {
	genesis := NewBlock("创世区块", "") // 造创世块，PrevHash 为空
	bc := &Blockchain{
		Blocks: []*Block{genesis}, // 初始化切片，里面先放一个创世块
		// Java: this.blocks = new ArrayList<>(); this.blocks.add(genesis);
	}
	return bc
}

// AddBlock：往链尾追加一个新区块
// 接收者 (bc *Blockchain) = Java 的 this
func (bc *Blockchain) AddBlock(data string) {
	// 取当前最后一个区块。 Java: Block prev = blocks.get(blocks.size()-1);
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	// 用"前一个块的哈希"当新块的 PrevHash —— 这就是"链"的关键
	newBlock := NewBlock(data, prevBlock.Hash)

	// 追加到切片末尾。 Java: blocks.add(newBlock);
	bc.Blocks = append(bc.Blocks, newBlock)
}

// PrintChain：遍历打印整条链，验证用
func (bc *Blockchain) PrintChain() {
	// for range 遍历切片。 Java: for (Block b : blocks)
	// i 是下标，block 是元素；下标用不上时可写 _ 忽略
	for i, block := range bc.Blocks {
		fmt.Printf("=== 区块 %d ===\n", i)
		fmt.Printf("数据:     %s\n", block.Data)
		fmt.Printf("前块哈希: %s\n", block.PrevHash)
		fmt.Printf("本块哈希: %s\n\n", block.Hash)
		fmt.Printf("Nonce:    %d\n\n", block.Nonce) // 新增
	}
}

// IsValid:校验整条链是否完整、未被篡改
// 返回 true = 链合法;false = 检测到篡改
func (bc *Blockchain) IsValid() bool {

	valid := true // 先假设合法,发现问题就标记为 false,但不立即退出

	// ===== 先单独校验创世区块(下标0)=====
	genesis := bc.Blocks[0]
	powGenesis := NewProofOfWork(genesis)
	if !powGenesis.Validate() {
		fmt.Printf("[篡改提醒] 创世区块(区块0)被篡改!\n")
		valid = false
	}

	// ===== 再校验后续每个块 =====
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		// 检查①:哈希是否合法
		pow := NewProofOfWork(currentBlock)
		// 哈希对不上自己 或 和前块链接断了 —— 都是被篡改
		if !pow.Validate() || currentBlock.PrevHash != prevBlock.Hash {
			fmt.Printf("[篡改提醒] 区块 %d 被篡改!\n", i)
			valid = false
		}

	}

	return valid // 全部查完再返回
}
