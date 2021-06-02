package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rl "github.com/i3h/ratelimiter"
)

func main() {
	r := gin.New()
	r.Use(rl.Limit(
		rl.Config{
			SizeOfBuffer: 500,
			Duration:     time.Second,
		}))
	r.GET("/", Test)
	go r.Run(":5000")
	time.Sleep(1 * time.Second)

	pass := 0
	block := 0
	start := time.Now()

	for {
		resp, _ := http.Get("http://127.0.0.1:5000")
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		if resp.StatusCode == 200 {
			pass++
		} else {
			block++
		}

		if time.Now().After(start.Add(5 * time.Second)) {
			break
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println()
	fmt.Printf("Elapsed: %.2fs\n", elapsed.Seconds())
	fmt.Printf("Pass:    %d\n", pass)
	fmt.Printf("Block:   %d\n", block)
	fmt.Printf("QPS:     %.2f\n", float64(pass)/elapsed.Seconds())
}

func Test(c *gin.Context) {
	s := time.Now().String() + "\n"
	c.String(200, s)
}
