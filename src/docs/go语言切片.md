# Go语言切片
    Go语言切片是对数组的抽象
    Go数组的长度不可改变 在特定场景中这样的集合就不在适用 而切片即使一个长度不固定 可以最佳元素 追加是可能使切片容量增大
    
## 定义切片
    你可以声明一个未指定大小的数组来定义切片
    var identifier []type
    切片不需要说明长度
    也可以使用make()函数来创建切片
    var slice1 []type = make([]type, len)
    也可以简写
    clice1 := make([]type, len)
    也可以指定容量， 其中capacity为可选参数
    make([]T, length, capacity)
    这里len是数组的长度并且也是切片的初始长度    