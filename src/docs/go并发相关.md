# goroutine特性
    go的执行是非阻塞的, 不会等待
    go后面的函数返回值会被忽略。
    调度器不能保证多个goroutine的执行次序
    没有父子goroutine的概念 所有的goroutine是平等的被调度和执行的
    Go程序在执行的时候会单独为main函数创建一个goroutine, 遇到其他go关键字时再去创建其他的goroutine
    Go没有暴露goroutine Id给用户， 所以不能在一个goroutine里面显式的地操作另外一个gorontine,
    不过runtine包提供了一些函数访问和设置gorontine的相关信息