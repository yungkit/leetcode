package main

import (
	"fmt"
	"sync"
)

// 三个协程交替打印 cat dog fish
var repeatCount = 100

func main() {
	// wg 用来防止主协程提前先退出
	wg := &sync.WaitGroup{}
	wg.Add(3)

	chCat := make(chan struct{}, 1)
	defer close(chCat)
	chDog := make(chan struct{}, 1)
	defer close(chDog)
	chFish := make(chan struct{}, 1)
	defer close(chFish)

	go printAnimal(chCat, chDog, "cat", wg)
	go printAnimal(chDog, chFish, "dog", wg)
	go printAnimal(chFish, chCat, "fish", wg)
	chCat <- struct{}{}
	wg.Wait()
}

// wg 需要传指针
func printAnimal(in, out chan struct{}, s string, wg *sync.WaitGroup) {
	count := 0
	for {
		<-in
		count++
		fmt.Println(s)
		out <- struct{}{}

		if count >= repeatCount {
			wg.Done()
			return
		}

	}
}
