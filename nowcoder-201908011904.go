// 字符集合
// 输入一个字符串，求出该字符串包含的字符集合

// 输入描述:
// 每组数据输入一个字符串，字符串最大长度为100，且只包含字母，不可能为空串，区分大小写。

// 输出描述:
// 每组数据一行，按字符串原有的字符顺序，输出字符集合，即重复出现并靠后的字母不输出。

// 输入例子1:
// abcqweracb

// 输出例子1:
// abcqwer
package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var s string
		n, _ := fmt.Scan(&s)
		if n == 0 {
			break
		}

		var s1 string
		for _, r := range s {
			if -1 == strings.IndexRune(s1, r) {
				s1 += string(r)
			}
		}
		fmt.Println(s1)
	}
}
