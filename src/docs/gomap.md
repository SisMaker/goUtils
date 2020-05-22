# KeyValue限制
    Key一定要是是可比较类型（可以理解为支持 == 的操作）
    非法类型的Key 如果是非法的key类型，会报错：invalid map key type xxx。
    
    可比较类型	                不可比较类型
    boolean	                    slice
    numeric	                    map
    string	                    func
    pointer	
    channel	
    interface	
    包含前文类型的array和struct

# 函数传递
    Golang中是没有引用传递的，均为值传递。这意味着传递的是数据的拷贝。
    那么map本身是引用类型，作为形参或返回参数的时候，传递的是值的拷贝，而值是地址，扩容时也不会改变这个地址。    