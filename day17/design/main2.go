package main

import "fmt"

// 结构体
type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

// 声明
type Option func(o *Options)

// 初始化结构体
func InitOptions1(opts ...Option) {
	options := &Options{}
	// 遍历参数,得到每一个函数
	for _, opt := range opts {
		//调用函数，在函数里，给传进去的对象赋值
		opt(options)
	}
	fmt.Printf("init Options: %v\n", options)
}

// 定义具体给每个字段赋值的方法
func WithStrOption1(str string) Option {
	return func(o *Options) {
		o.strOption1 = str
	}
}

func WithStrOption2(str string) Option {
	return func(o *Options) {
		o.strOption2 = str
	}
}
func WithStrOption3(str string) Option {
	return func(o *Options) {
		o.strOption3 = str
	}
}

func WithIntOption1(i int) Option {
	return func(o *Options) {
		o.intOption1 = i
	}
}

func WithIntOption2(i int) Option {
	return func(o *Options) {
		o.intOption2 = i
	}
}

func WithIntOption3(i int) Option {
	return func(o *Options) {
		o.intOption3 = i
	}
}

func main() {
	InitOptions1(WithStrOption1("str1"), WithStrOption2("str2"), WithStrOption3("str3"),
		WithIntOption1(1), WithIntOption2(2), WithIntOption3(3))
}
