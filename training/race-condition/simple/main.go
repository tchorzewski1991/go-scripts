package main

import (
	"fmt"
	"sync"
	"time"
)

var arr []int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	arr = make([]int, 10)

	wg.Add(2)
	go addY(5)
	go addX(2)
	wg.Wait()

	//fmt.Println(arr)
}

func addX(x int) {
	for i := 0; i < 3; i++ {
		//k := i + x;
		//arr = append(arr, k)
		mu.Lock()
		{
			fmt.Printf("%v. addX - before - arr[ %v ] : %v \n", i, i, arr[i])
			arr[i] += i + x
			fmt.Printf("%v. addX - after  - arr[ %v ] : %v \n", i, i, arr[i])
		}
		mu.Unlock()
		time.Sleep(time.Duration(time.Second * 2))
	}
	wg.Done()
}

func addY(y int) {
	for i := 0; i < 3; i++ {
		//k := i + y;
		//arr = append(arr, k)
		mu.Lock()
		{
			fmt.Printf("%v. addY - before - arr[ %v ] : %v \n", i, i, arr[i])
			arr[i] += i + y
			fmt.Printf("%v. addY - after  - arr[ %v ] : %v \n", i, i, arr[i])
		}
		mu.Unlock()
		time.Sleep(time.Duration(time.Second * 1))
	}
	wg.Done()
}
