package ratelimiter

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	SizeOfBuffer int
	Duration     time.Duration
}

func Limit(config Config) gin.HandlerFunc {
	buffer := make(chan time.Time, config.SizeOfBuffer)
	ticker := time.NewTicker(config.Duration)
	go func() {
		for {
			<-ticker.C
			for len(buffer) > 0 {
				<-buffer
			}
		}
	}()

	return func(c *gin.Context) {
		select {
		case buffer <- time.Now():
			c.Next()
		default:
			c.AbortWithStatus(429)
		}

	}
}
