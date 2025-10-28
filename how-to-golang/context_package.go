package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.WithValue(context.Background(), "foo", "bar")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", val)
	fmt.Println("Took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println("Context Value from \"foo\": ", val)

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	responseChan := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffThatCanBeSlow()
		responseChan <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long")
		case resp := <-responseChan:
			return resp.value, resp.err
		}
	}

}

func fetchThirdPartyStuffThatCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 120)

	return 42069, nil
}
