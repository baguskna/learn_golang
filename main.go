// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type counter struct {
// 	sync.Mutex
// 	val int
// }

// func (c *counter) Add() {
// 	c.Lock()
// 	c.val++
// 	c.Unlock()
// }

// func (c *counter) Value() (x int) {
// 	return c.val
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var meter counter

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)

// 		go func() {
// 			for j := 0; j < 1000; j++ {
// 				meter.Add()
// 			}

// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(meter.Value())
// }
