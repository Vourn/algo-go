// 栈和队列习题
package stack_queue

// 最小栈
type MinStack struct {
    min []int
    stack []int
}

func Constructor() MinStack {
    return MinStack {
	min: make([]int, 0),
	stack: make([]int, 0),
    }
}

func (this *MinStack) Push(val int) {
    min := this.GetMin()
    if val < min {
	this.min = append(this.min, val)
    } else {
	this.min = append(this.min, min)
    }
    this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
    if len(this.stack) == 0 {
	return
    }
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
	return 0
    }
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    if len(this.min) == 0 {
	return 1 << 31
    }
    return this.min[len(this.min)-1]
}

// 逆波兰式求值
func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
	return 0	
    }
    stack := make([]int, 0)
    for i := 0; i < len(tokens); i++ {
        switch tokens[i] {
            case "+","-","*","/":
                a := stack[len(stack) - 1]
                b := stack[len(stack) - 2]
                stack = stack[:len(stack) - 2]
                var num int
                switch tokens[i] {
                    case "+":
                        num = b + a
                    case "-":
                        num = b - a
                    case "*":
                        num = b * a
                    case "/":
                        num = b / a
                }
                stack = append(stack, num)
            default:
                val, _ := strconv.Atoi(tokens[i])
                stack = append(stack, val)
        }
    }
    return stack[0]
}

// 字符串解码（栈操作）
func decodeString(s string) string {
    if len(s) == 0 {
        return ""
    }
    stack := make([]byte, 0)
    numStack := make([]int, 0)
    num := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ']' {
            // 临时栈，若直接stack切片赋值，后续stack底层数组修改会影响切片值
            temp := make([]byte, 0)
            for j := len(stack) - 1; j >= 0; j-- {
                if stack[j] == '[' {
                    // 取栈顶数字
                    repeatK := numStack[len(numStack) - 1]
                    stack = stack[:j]
                    numStack = numStack[:len(numStack) - 1]
                    // 字符串重复入栈
                    for k := 0; k < repeatK; k++ {
                        for z := len(temp)-1; z >= 0; z-- {
                            stack = append(stack, temp[z])
                        } 
                    }
                    break;
                } else {
                    temp = append(temp, stack[j])
                }
            }
        } else if s[i] >= '0' && s[i] <= '9' {
            n, _ := strconv.Atoi(string(s[i]))
            num = num * 10 + n
        } else {
            if s[i] == '[' {
                numStack = append(numStack, num)
                num = 0
            }
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}
