package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("ci-buildinfo")
	fmt.Println("Current time:", time.Now())
	fmt.Println("Go compiler:", runtime.Version())
	fmt.Println("OS info:", runtime.GOOS, runtime.GOARCH)
}
