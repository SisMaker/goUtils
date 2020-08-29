# 代码元素
    关键字
    标识符
    基本类型 和 它们的字面量
    常量 变量
    操作符
    函数声明与函数调用
    代码包和包引入
    表达式，语句和简单语句
    基本流程控制语法

# 关键字 共25个分四组
    用来声明各种代码元素
        const func import package type var
    用于组合类型的字面表示
        chan interface map struct
    基本流程控制
        break case continue default else fallthrough for goto if range switch return select 
    特殊流程关键字
        defer go            


# 标识符
    一个标识符是以unioncode字母或者 _ 开头并且完全由unioncode字母和uninocode数字组成的单词
        Unicode字母是定义在Unicode标准8.0 ? 中的Lu、Ll、Lt、Lm和Lo分类中的字符。
        Unicode数字是定义在Unicode标准8.0中的Nd数字字符分类中的字符。
    所有的类型名， 变量名 常量名 跳转标签 包名 和包的引入名 都必须是标识符
    由UnionCode大写字母开头的标识符成为导出标识符    

# 基本类型和它们的字面量表示
## 基本的内置类型（go支持下列基本类型） 也称为 预声明类型
    1种布尔类型 bool
    11种中内置整数类型
        int8 uint8 int16 uint16 int32 uint32 int64 uint64 int uint uintptr
    2种浮点数类型
        float32 float64
    2种内置的复数类型
        complex64 complex128
    1种字符串
        string
    
    除了bool和string类型 其他15中内置基本类型都称为数值类型
    
    Go中有两种内置类型别名（type alias）
        byte是uint8的内置别名 我们可以将byte 和 uint8 看作同一个类型
        rune是uint32内置别名  我们可以将rune 和 uint32 看作容一个类型
    
    其中 uintptr int uint 类型的值的尺寸都依赖于具体的编译器实现 通常在64位架构上 int，uint类型的
    值是64位，在32位是则是32位     
    编译器必输保证uintptr类型的值的尺寸能够存下任意一个内存地址
     
    complex64复数值的实部和虚部都是float32 
    complex128复数值的实部和虚部都是float64
    在内存中 浮点数都是用IEEE-754格式存储
    
    尽管布尔和字符串类型分类各自只有一种类型，我们可以声明更多的自定义布尔和字符串类型

    1| // 一些类型定义声明
    2| type status bool // status和bool是两个不同的类型
    3| type MyString string // MyString和string是两个不同的类型
    4| type Id uint64 // Id和uint64是两个不同的类型
    5| type real float32 // real和float32是两个不同的类型
    6|
    7| // 一些类型别名声明
    8| type boolean = bool // boolean和bool表示同一个类型
    9| type Text = string // Text和string表示同一个类型
    10| type U8 = uint8 // U8、uint8和 byte表示同一个类型
    11| type char = rune // char、rune和int32表示同一个类型

# 零值
    每种类型都有一个零值， 零值都可以看做此类型的默认值, 相应类型的零值如下
        bool: false
        数值类型: 0 (虽然都是零 但是内存占用不同)
        字符串: ""

# 基本类型的字面量表示形式
    一个值的字面形式称为一个字面量，它表示此值在代码中的字体体现形式(和内存中的表现形式相对应)
    一个值可能有很多种字面量形式
## 布尔值的字面量形式
    true false
## 整数类型值的字面量形式
    整数类型值有四种字面量形式：
        十进制形式(decimal)  不能以 0 开头
        八进制形式(octal) ）0o 0O 0 开头
        十六进制形式(hex) 0x 0X 开头
        二进制形式(binary) 0b 0B 开头
##  浮点数类型值的字面量形式
    一个浮点数的完整字面量形式包含一个十进制整数部分， 一个小数点 一个十进制小数部分和一个整数指数部分，部分可以省略













