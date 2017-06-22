package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	echo1()
	echo2()
	echo3()
	echo4()
	echoBonus()
	fmt.Printf("%.20f elapsed (total)\n\n", time.Since(start).Seconds())
}

func echo1() {
	//start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	//fmt.Printf("%.8f elapsed (echo1)\n\n", time.Since(start).Seconds())
}

func echo2 () {
	//start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	//fmt.Printf("%.8f elapsed (echo2)\n\n", time.Since(start).Seconds())
}

func echo3() {
	//start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	fmt.Println(strings.Join(os.Args[0:]," "))
	//fmt.Printf("%.8f elapsed (echo3)\n\n", time.Since(start).Seconds())
}

func echoBonus() {
	//start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	fmt.Println(os.Args[0:])
	//fmt.Printf("%.8f elapsed (echoBonus)\n\n", time.Since(start).Seconds())
}

func echo4() {
	//start := time.Now()
	//duration := time.Second
	//time.Sleep(duration)
	for i, arg := range os.Args {
		fmt.Printf("%d %s\n",i,arg)
	}
	//fmt.Printf("%.8f elapsed (echo4)\n\n", time.Since(start).Seconds())
}