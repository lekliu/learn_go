package main

type A interface {
	data1(int)
}

type B interface {
	data1(int)
	data2(int)
}

type Impl_B struct {
}

func (this *Impl_B) data1(i int) {
	println(i)
}
func (this *Impl_B) data2(i int) {
	println(i)
}

type Impl_A struct {
}

func (this *Impl_A) data1(i int) {
	println(i)
}

func main() {
	// B的具体实现才可以赋值给A
	var a A = &Impl_B{}
	var b B = &Impl_A{}

}
