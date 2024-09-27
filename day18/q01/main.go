package main

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}
}

func main() {
	p := f()
	print(p.m)
}
