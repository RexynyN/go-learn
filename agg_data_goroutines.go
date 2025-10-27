package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Old Code
// likes := fetchUserLikes(userN'ame)
// match := fetchUserMatch(userName)
// fmt.Println("Likes: ", likes)
// fmt.Println("Match: ", match)

func main() {
	start := time.Now()

	userName := fetchUser()
	respChan := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go fetchUserLikes(userName, respChan, wg)
	go fetchUserMatch(userName, respChan, wg)

	wg.Wait() // Block until the two goroutines are finished
	close(respChan)

	for resp := range respChan {
		fmt.Println("Response: ", resp)
	}

	fmt.Println("Time Took", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Breno"
}

func fetchUserLikes(userName string, respChan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 233)

	respChan <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respChan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 170)

	respChan <- "Kayla"
	wg.Done()
}

func assertInt(value any) (int, error) {
	intValue, ok := value.(int)
	if ok == false {
		return 0, errors.New("not an int value")
	}

	return intValue, nil
}
