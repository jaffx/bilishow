package main

import (
	"fmt"
	"time"
)

func wait1() {
	time.Sleep(2 * time.Second)
	wait2()
}
func wait2() {
	time.Sleep(3 * time.Second)
}

func main() {
	t1 := time.Now().Unix()
	go wait1()
	wait2()
	t2 := time.Now().Unix()
	// time.Sleep(3)
	fmt.Print(t2 - t1)
}
