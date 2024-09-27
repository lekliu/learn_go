package main

// 描述 golang 中的stack和heap的区别， 分别在什么情况下会分配stack? 又在何时会分配到heap中
区别：
 栈(stack): 由编译器自动分配和释放，存变量名、各种名
 堆（heap）：在C里由程序嘈分配和释放内存，go自动了，存变量的数据值

make(XXX)  a:=3  a就在栈中，3 在堆中