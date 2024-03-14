package http

import (
	index "distributed_log/internal/log"
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	c := NewClient()

	_, err := c.InsertIndex(index.Index{
		File: "234",
	})
	if err != nil {
		fmt.Print(err)
	}
}

func Test(t *testing.T) {
	c := NewClient()
	//var wg sync.WaitGroup
	loop := 100000
	//wg.Add(loop)
	for i := 0; i < loop; i++ {
		//go func() {
		res, err := c.InsertIndex(index.Index{
			File:       fmt.Sprintf("%d", i),
			ImportTime: time.Now(),
		})
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(res)
		//time.Sleep(1 * time.Millisecond)
		//wg.Done()
		//}()

	}
	//wg.Wait()
}
