package goUtils

import (
	"reflect"
	"unsafe"
)

// string与[]byte的转换
// Golang中，string从设计上是不可变（immutable）的。因此,string和[]byte的类型转换，都是产生一份新的副本。
//
// func Example() {
// 	s := "Hello,world"
// 	b := []byte(s)
// }
// 如果确定转换的string/[]byte不会被修改，可以进行直接的转换，这样不会生成原有变量的副本。新的变量共享底层的数据指针。

func StrToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToStr(b []byte) string {
	sliceHeader
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
