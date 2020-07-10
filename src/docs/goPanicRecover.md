# 作用
    pacin用来主动抛出错误
    recover用来捕捉panic抛出的错误

# 基本概念
    函数签名
    panic(i interface)
    recover() interface{}
    
    引发panic有两种情况 一种是程序主动调用panic函数，另外一种就是程序产生运行时的错误，由运行是检测并抛出
    
    发生panic后 程序会从调用panic的函数位置或发生panic的地方立即返回， 逐层向上执行函数的defer语句，
    然后逐层打印函数调用堆栈， 直到被recover捕获或运行到最外层函数而退出
    
    panic的参数是个空的接口类型interface{} 所以任意类型的变量都可以传参给panic
    panic 不但可以在函数正常流程中抛出， 在defer逻辑里也可以再次调用panic或抛出panic
    defer里面的panic能够被后续执行的defer捕获
    recover()用来捕获panic,阻止panic继续向上传递， recover()和defer一起使用，但是recover()
    只有在defer后面的函数体内被直接调用才能捕获panic终止异常， 否则返回nil， 异常继续向外传递
    
    
    
    
    
    
    
    
        