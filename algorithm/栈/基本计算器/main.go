package main

import (
	"fmt"
	"unicode"
)

// calc 执行一个单独的运算：从 ops 中取出一个操作符，从 nums 中取出两个操作数，
// 计算结果并将其再次放入 nums 中。
func calc(nums *[]int, ops *[]string) {
	// 弹出操作符
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	// 弹出第二个操作数 (b)
	b := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]

	// 弹出第一个操作数 (a)
	a := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	}

	// 将结果推入栈
	*nums = append(*nums, result)
}

// calculate 实现了基础计算器的逻辑
func calculate(s string) int {
	var nums []int   // 存储数字（操作数）的栈
	var ops []string // 存储操作符的栈
	n := len(s)
	i := 0

	for i < n {
		char := s[i]

		// 1. 跳过空格
		if unicode.IsSpace(rune(char)) {
			i++
			continue
		}

		// 2. 解析数字
		if unicode.IsDigit(rune(char)) {
			num := 0
			for i < n && unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, num)
			continue
		}

		// 3. 处理左括号 '('
		if char == '(' {
			ops = append(ops, string(char))
			i++
			continue
		}

		// 4. 处理右括号 ')'
		if char == ')' {
			// 在栈中遇到 '(' 之前，持续计算
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				calc(&nums, &ops)
			}
			// 弹出 '('
			if len(ops) > 0 && ops[len(ops)-1] == "(" {
				ops = ops[:len(ops)-1]
			}
			i++
			continue
		}

		// 5. 处理操作符 '+' 或 '-'
		if char == '+' || char == '-' {
			// --- 开始修改 (处理一元操作符) ---
			isUnary := true
			for j := i - 1; j >= 0; j-- {
				if unicode.IsSpace(rune(s[j])) {
					continue
				}
				// 如果上一个非空格字符是 ')' 或一个数字，
				// 则当前操作符是二元操作符。
				if s[j] == ')' || unicode.IsDigit(rune(s[j])) {
					isUnary = false
				}
				break // 找到上一个非空格字符后，退出循环。
			}

			if isUnary {
				// 如果操作符是一元的，向数字栈中添加一个 0。
				// 例如: "(-2)" 变为 "(0-2)"
				nums = append(nums, 0)
			}
			// --- 结束修改 ---

			// 在栈中遇到 '(' 之前，持续计算
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				calc(&nums, &ops)
			}
			// 将当前操作符推入栈
			ops = append(ops, string(char))
			i++
			continue
		}
	}

	// 循环结束后，计算栈中剩余的操作
	for len(ops) > 0 {
		calc(&nums, &ops)
	}

	// 返回最终结果
	if len(nums) > 0 {
		return nums[0]
	}
	return 0
}

func main() {
	// 示例 1: 1 - (-2) = 3
	s1 := "1 - (   -2)"
	fmt.Printf("表达式: \"%s\" = %d\n", s1, calculate(s1))

	// 示例 2: (1 + 2) - (-3) = 6
	s2 := "(1 + 2) - (-3)"
	fmt.Printf("表达式: \"%s\" = %d\n", s2, calculate(s2))

	// 示例 3: -2 + 1 = -1
	s3 := "-2 + 1"
	fmt.Printf("表达式: \"%s\" = %d\n", s3, calculate(s3))
}
