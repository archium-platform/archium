package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func GetMemoryUsage() map[string]any {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	alloc := bToMb(m.Alloc)
	totalAlloc := bToMb(m.TotalAlloc)
	sys := bToMb(m.Sys)
	numGC := bToMb(uint64(m.NumGC))

	return map[string]any{
		"alloc":      alloc,
		"totalAlloc": totalAlloc,
		"sys":        sys,
		"numGC":      numGC,
	}
}

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		memStats := GetMemoryUsage()
		ctx.IndentedJSON(200, gin.H{
			"status": "ok",
			"mem":    memStats,
		})
	})

	r.Run()
}
