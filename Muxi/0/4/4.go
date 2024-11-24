package main

import "fmt"

// 定义sjn结构体
type sjn struct {
	m_name        string
	m_appearance  int
	m_voice       int
	m_personality int
}

var sjns = []sjn{}

// 输入信息
func input() {

	var (
		name        string
		appearance  int
		voice       int
		personality int
	)

	for i := 0; i < 6; i++ {
		fmt.Scanln(&name, &appearance, &voice, &personality)
		sjns[i] = sjn{name, appearance, voice, personality}
	}
}

// 进行排序
type sorts func(s []sjn)

//?都是一类

func sort(s []sjn) {
	sort_app(s)
}

func sort_app(s []sjn) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6-i; j++ {
			if s[j].m_appearance < s[j+1].m_appearance {
				s[j].m_appearance, s[j+1].m_appearance = s[j].m_appearance, s[j+1].m_appearance
			}
			if s[j].m_appearance == s[j+1].m_appearance {
				sort_voice(s)
			}
		}
	}
}

func sort_voice(s []sjn) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6-i; j++ {
			if s[j].m_voice < s[j+1].m_voice {
				s[j].m_voice, s[j+1].m_voice = s[j].m_voice, s[j+1].m_voice
			}
			if s[j].m_voice == s[j+1].m_voice {
				sort_per(s)
			}
		}
	}
}

func sort_per(s []sjn) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6-i; j++ {
			if s[j].m_personality < s[j+1].m_personality {
				s[j].m_personality, s[j+1].m_personality = s[j].m_personality, s[j+1].m_personality
			}
		}
	}
}

// 输出信息
func output() {
	for i := 0; i < 6; i++ {
		fmt.Println(sjns[i])
	}
}
func main() {
	input()
	sort(sjns)
	output()
}
