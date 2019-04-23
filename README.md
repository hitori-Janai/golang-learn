# Golang

> [Go语言圣经](https://docs.hacknode.org/gopl-zh/index.html "霸气的名字")

## 入门
* 变量初始化
    ```golang
    s := "" //短变量初始化,只能用于函数内
    var s string //默认初始化零值 为""
    var s = ""
    var s string = ""
    ```
* fmt.Printf
    ```
    %d          十进制整数
    %x, %o, %b  十六进制，八进制，二进制整数。
    %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
    %t          布尔：true或false
    %c          字符（rune） (Unicode码点)
    %s          字符串
    %q          带双引号的字符串"abc"或带单引号的字符'c'
    %v          变量的自然形式（natural format）
    %T          变量的类型
    %%          字面上的百分号标志（无操作数）
    ```
    