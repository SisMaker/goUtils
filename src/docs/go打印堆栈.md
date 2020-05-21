# 打印堆栈
    fmt.Printf("%s", debug.Stack())
    debug.PrintStack()
    可以通过 debug.PrintStack() 直接打印，也可以通过 debug.Stack() 方法获取堆栈然后自己打印。