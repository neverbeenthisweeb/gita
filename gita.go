package gita

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var counterVec = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gita_http_requests",
	Help: "Total number of HTTP requests",
}, []string{"url", "status"})

var histogramVec = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "gita_http_latency",
	Help: "Latency of HTTP requests",
}, []string{"url", "status"})

func HandleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip prometheus metrics endpoint
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		startTime := time.Now()

		c.Next()

		counterVec.With(prometheus.Labels{
			"url":    c.Request.URL.String(),
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		elapsedSec := float64(time.Since(startTime)) / float64(time.Second)
		histogramVec.With(prometheus.Labels{
			"url":    c.Request.URL.String(),
			"status": strconv.Itoa(c.Writer.Status()),
		}).Observe(elapsedSec)
	}
}
