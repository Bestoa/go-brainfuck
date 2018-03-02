package main
import (
    "fmt"
)

func main() {
    var code string
    _, err := fmt.Scanln(&code)
    if err != nil {
        fmt.Println(err)
        return
    }
    interpreter(code)
}

func interpreter(code string) {
    var cp, sp, jsp, end int = 0, 0, 0, len(code)
    var stack [1024]byte
    var jump_stack [512]int
    for cp < end {
        switch code[cp] {
        case '>': sp++
        case '<': sp--
        case '+': stack[sp]++
        case '-': stack[sp]--
        case '.': fmt.Printf("%c", stack[sp])
        case ',': fmt.Scanf("%c", &stack[sp])
        case '[':
            if stack[sp] == 0 {
                cp++
                var need_paried_bracket = 0
                var find = false
                for cp < end {
                    switch code[cp] {
                    case '[': need_paried_bracket++
                    case ']':
                        if need_paried_bracket == 0 {
                            find = true
                        } else {
                            need_paried_bracket--
                        }
                    }
                    if find {
                        break
                    }
                    cp++
                }
                if !find {
                    fmt.Println("Unpaired bracket!")
                    return
                }
            } else {
                jump_stack[jsp] = cp
                jsp++
            }
        case ']':
            if stack[sp] != 0 {
                cp = jump_stack[jsp-1]
            } else {
                jsp--
            }
        default:
            fmt.Println("Unexcept char, ascii:", code[cp])
        }
        //fmt.Printf("%c, %d\n", code[cp], jsp)
        cp++
    }
}
