package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", Test)
	go r.Run(":5000")
	time.Sleep(1 * time.Second)

	counter := 0
	start := time.Now()

	for {
		resp, _ := http.Get("http://127.0.0.1:5000")
		resp.Body.Close()
		counter++
		if time.Now().After(start.Add(5 * time.Second)) {
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println()
	fmt.Printf("Elapsed: %.2fs\n", elapsed.Seconds())
	fmt.Printf("Counter: %d\n", counter)
	fmt.Printf("QPS:     %.2f\n", float64(counter)/elapsed.Seconds())
}

func Test(c *gin.Context) {
	s := time.Now().String() + "\n"
	c.String(200, s)
}
