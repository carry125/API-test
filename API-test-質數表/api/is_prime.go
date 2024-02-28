package api

import (
	"fmt"
	"strconv"
	"test/cache"
	"test/calculator"

	"github.com/gin-gonic/gin"
)

func CheckAndInsertPrimeValues(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	numStartStr := ctx.PostForm("numStart")
	numEndStr := ctx.PostForm("numEnd")

	numStart, err := strconv.Atoi(numStartStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid value for numStart"})
		fmt.Printf("%s", err)
		return
	}

	numEnd, err := strconv.Atoi(numEndStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid value for numEnd"})
		return
	}

	if numStart > numEnd {
		newStart, newEnd := calculator.Swap(numStart, numEnd)
		interval, _ := cache.CalculateAndCachePrimeRange(newStart, newEnd)
		ctx.JSON(200, gin.H{"prime_values": interval})
		return
	}

	if numStart <= 0 || numEnd <= 0 {
		ctx.JSON(400, gin.H{"error": "Invalid value"})
		return
	}

	if numStart == 1 && numEnd == 1 {
		ctx.JSON(400, gin.H{"error": "not prime"})
		return
	}

	interval, _ := cache.CalculateAndCachePrimeRange(numStart, numEnd)
	ctx.JSON(200, gin.H{"prime_values": interval})
}
