package main // 可执行程序的入口包必须叫 main

import (
	"fmt"
	"time"
) // 导入标准库的格式化包

func main() { // main 函数是程序入口

	start := time.Now() // 记录开始时刻。Java: long start = System.nanoTime();

	//genesis := NewBlock("创世区块", "")
	//fmt.Printf("数据: %s\n哈希: %s\n", genesis.Data, genesis.Hash)
	bc := NewBlockchain() // 创建链（含创世块）

	//for i := 1; i <= 100000; i++ {
	//	bc.AddBlock(fmt.Sprintf("交易 %d", i)) //加 10 万个块，故意让它慢下来
	//}
	//
	//elapsed := time.Since(start) // 从 start 到现在的耗时。等于 time.Now().Sub(start)
	//fmt.Printf("程序耗时: %s\n", elapsed)
	//fmt.Printf("程序耗时(纳秒): %d ns\n", elapsed.Nanoseconds()) // 这个能看到真实值
	//
	bc.AddBlock("张三转给李四 5个 BTC") //转账测试
	bc.AddBlock("李四转给王五 10个BTC") //转账测试

	bc.PrintChain()
	bc.Blocks[1].Hash = "000000000000000000000sdhfhsdhfasgdasgjfgajdgjagdj"
	bc.IsValid()
	fmt.Printf("总耗时: %s\n", time.Since(start))

}
