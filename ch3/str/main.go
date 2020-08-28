package main

import "fmt"

func main() {

	s := "hello world，你好."
	fmt.Println(len(s))
	fmt.Println(s[0], s[7]) // 104 111

	// 非ASCII字符的UTF-8需要多个字节,所以无法用下标来访问字符串
	// fmt.Print(s[len(s)]) //下标越界

	// 生成新字符串
	s1 := s[0:5]
	s2 := s[0:5] + "-a."
	fmt.Println(s, s1, s2, &s, &s1, &s2)

	// 比较:按字节进行
	s3 := "abc"
	s4 := "bc"
	fmt.Println(s3 == s4, s3 > s4, s3 < s4)

	// 字符串值无法改变，字节序列永不变
	// &s5一直没有改变？
	s5 := "left"
	fmt.Println(&s5)
	s6 := s5
	fmt.Println(&s5, &s6)
	s5 += ",right"
	fmt.Println(s5, s6, &s5, &s6)
	s5 = ",right" + s5
	fmt.Println(s5, s6, &s5, &s6)
	// s[0] = 'L' //s[0]无法赋值，编译错误

	// 多个字符串可以安全的共享内存?内存地址不同是同一地址吗？
	// linux验证一下
	s7 := "hello world"
	hello := s7[:5]
	world := s7[7:]
	fmt.Print(&s7, &hello, &world)

	// 转义 \(从左到右)
	// 字符串字面量(string literal)
	// 转义仅在双引号的字节序列中有用
	// 反引号的字节序列中转义字符序列不起作用
	// \a \b \f \n \r \t \v \' \" \\
	// \xhh

	// Unicode(utf-32)
	// ASCII US-ASCII 美国信息交换标准码
	// Unicode：unicode标准数字可以用int32表示(4byte),将文字符号的序列表示成int32的值序列，
	// 这种表示方式叫utf-32或UCS-4，每个码点的长度相同，都是4byte。
	// 大多计算机可读文本是ASCII，只需要1byte，常用的也少于65556个，用2byte就可以。
	// Unicode编码的缺点是浪费空间。

	// utf-8
	// 以字节为单位对Unicode码点变长编码。utf-8是Unicode的一中标准。
	// 每个文字符号用1~4个字节表示。一个文字符号的首字节高位指明还有多少字节。
	// 0xxxxxxx                            0~127          ASCII(1byte)
	// 110xxxxx 10xxxxxx                   128~2047       <128(2byte)
	// 1110xxxx 10xxxxxx 10xxxxxx          2048~65535     <2048(3byte)
	// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536~0x10ffff <65535(4byte)
	// 变长字符串无法按照下标直接访问第n个字符。

	//

}
