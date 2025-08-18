package helper

import "html/template"

// 在 main 函数之前添加这些函数
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetTplFuncs() template.FuncMap {
	return template.FuncMap{
		"seq": func(start, end int) []int {
			var s []int
			for i := start; i <= end; i++ {
				s = append(s, i)
			}
			return s
		},
		"add": add,
		"sub": sub,
		"mul": mul,
		"div": div,
		"max": max,
		"min": min,
	}
}
