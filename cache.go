package cache

import (
	"fmt"
	"test/calculator"
	"time"

	"github.com/patrickmn/go-cache"
)

// 把質數表放入快取
type Prime struct {
	Value int
}

var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}
func QueryPrimesInRange(start int, end int) ([]int, error) {
	// 檢查快取中是否有已計算的質數範圍
	cachedPrimes, found := c.Get(fmt.Sprintf("primes-%d-%d", start, end))
	if found {
		// 如果在快取中找到，直接返回已計算的質數範圍
		return cachedPrimes.([]int), nil
	}
	// 如果快取中沒有找到，則計算質數並存儲到快取中(V)
	prime := calculator.Calculate(start, end)
	var primesInRange []int
	for _, solution := range prime {
		if solution >= start && solution <= end {
			primesInRange = append(primesInRange, solution)
		}
	}
	// 將計算結果存儲到快取中，設置適當的過期時間
	c.Set(fmt.Sprintf("primes-%d-%d", start, end), primesInRange, cache.DefaultExpiration)

	return primesInRange, nil
}
func UpdatePrimeRange(start int, end int) error {
	primeInRange := calculator.Calculate(start, end)

	for _, confirmPrime := range primeInRange {
		// 檢查快取中是否已經存在該質數
		if _, found := c.Get(fmt.Sprintf("%d", confirmPrime)); !found {
			// 如果不存在，將該質數作為鍵儲存到快取中，值為 true
			c.Set(fmt.Sprintf("%d", confirmPrime), true, cache.DefaultExpiration)

			// 同時也可以將該質數插入到資料庫中，可以實現快取和資料庫的雙重保護。
			/*if err := server.DateBaseConnection.Table("primes").
				Create(&Prime{Value: confirmPrime}).Error; err != nil {
				return err
			}*/
		}
	}
	return nil
}
