package main

import (
	"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

func main() {

	c := gen(2, 3, 4, 5)
	cSq := sq(c)

	/*
	   when i range in this manner, the n will get the integer
	   if i want the channel itself
	   i have to use
	   for index, value := range cSq , the value will be the channel ???
	*/
	if true {
		for n := range cSq {
			//fmt.Println("square of list = ", n)
			myTools.Describe(n)
		}
	}

	/* DONT WORK???
	if false {
		for _, value := range cSq {
			//fmt.Println("square of list = ", n)
			myTools.Describe(value)
		}
	}
	*/
}

func gen(alist ...int) chan int {
	out := make(chan int)
	go func() {
		for _, value := range alist {

			out <- value
		}
		close(out)
	}()
	return out
}

func sq(aCh chan int) chan int {
	outSq := make(chan int)
	go func() {
		for n := range aCh {
			outSq <- n * n //square of n
		}
		close(outSq)
	}()
	return outSq
}
