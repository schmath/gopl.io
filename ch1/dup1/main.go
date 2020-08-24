// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输入、输出、排序、文本处理
// fmt格式化输出和扫描输入
// main包和main函数:可执行程序的入口函数.
// import 使用其他包,告诉编译器代码依赖其他包
// func func_name(args) return val
// 分号:多个语句或声明出现在同一行,语句后面要写分号.
// 换行:函数名后面的{必须于函数名同一行;x+y换行符在操作符后面
// 格式化代码:gofmt,可以配置文件保存时自动调用gofmt.
// goimport (10.7)
// 输入:文件\网络连接\其他程序的输出\键盘输入\命令行参数(启动时args)
// os:平台无关的方式与操作系统打交道
// slice[m:n]:从第m个元素到n-1个元素。缺省值是m缺省值是0，n缺省值是len(s)
func main() {

	// 字符串
	str1 := "abcedfg"
	fmt.Println(str1[:])
	fmt.Println(str1[:3])

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
