package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string
	n, _ := fmt.Scanln(&a, &b, &c)
	str := []string{}
	str[0] = a
	str[1] = b
	str[2] = c

	for i := 0; i < n; i++ { //缩写还原
		strings.Replace(str[i], "i`m", "i am", -1)
		strings.Replace(str[i], "i`ll", "i will", -1)
	}

	for i := 0; i < n; i++ { //大小写转换
		ch := []byte(str[i])
		l := len(ch)
		for j := 0; j < l; j++ {
			if ch[j] < 90 {
				ch[j] += 32
			}
		}
	}

	//好词比较

	//n := []int{}
	hushtable := [26]int{0} //存放a - z的出现次数

	//从末尾遍历每个单词，如果是好词就继续
	for i := 0; i < n; i++ {
		ch := []byte(str[i])
		l := len(ch)

		for j := l - 1; j >= 0; j-- { //遍历每个单词的每个字母

			if ch[j] != ch[j-1] {
				//n[i] = ch[j] - 97 //记录好词在表中的序列数
				break
			} else {
				hushtable[ch[j]-97]++
			}

		}
	}

	max := 0
	for i := 0; i < 26; i++ {
		if hushtable[i] > max {
			max = i //找到好词字母对应下标
		}
	}

	//然后根据下标返回到对应的好词
}
