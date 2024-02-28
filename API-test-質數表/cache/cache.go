package cache

import (
	"sync"
	"test/calculator"
)

/*
var (

	cacheObject *cache.Cache

)

	func init() {
		cacheObject = cache.New(5*time.Minute, 10*time.Minute)
		cacheObject.Set("primes", []int{}, cache.DefaultExpiration) //init一個存在的cache
	}

	func GetPrimeCacheSlice() []int {
		primeCache, _ := cacheObject.Get("primes")
		primeCacheSlice := primeCache.([]int)
		return primeCacheSlice
	}

func CalculateAndCachePrimeRange(start int, end int) ([]int, error) { //沒有排序

		var outputInterval []int //輸出的答案
		primeCacheSlice := GetPrimeCacheSlice()
		for i := start; i <= end; i++ {
			if calculator.ComparePrime(i, primeCacheSlice) {
				outputInterval = append(outputInterval, i) //跟放在cache裡面的比，是質數就放到輸出
			} else if calculator.IsPrime(i) { //如果沒放在cache李，但他是質數，可解決第一次放入的問題
				outputInterval = append(outputInterval, i)   //更新output
				primeCacheSlice = append(primeCacheSlice, i) //更新cache
				//database.UpdateDBPrimeRange(start, end)      //cache miss
			}
		}
		cacheObject.Set("primes", primeCacheSlice, cache.DefaultExpiration) //最後儲存至cacheObject
		return outputInterval, nil
	}
*/
var (
	primeCache sync.Map
)

func init() {
	primeCache.Store("primes", []int{})
}

func GetPrimeCacheSlice() []int {
	primeCacheSlice, _ := primeCache.Load("primes")
	return primeCacheSlice.([]int)
}

func CalculateAndCachePrimeRange(start int, end int) ([]int, error) {
	var outputInterval []int
	primeCacheSlice := GetPrimeCacheSlice()

	for i := start; i <= end; i++ {
		if calculator.ComparePrime(i, primeCacheSlice) {
			outputInterval = append(outputInterval, i)
		} else if calculator.IsPrime(i) {
			outputInterval = append(outputInterval, i)
			primeCacheSlice = append(primeCacheSlice, i)
		}
	}
	//fmt.Println(primeCacheSlice)
	primeCache.Store("primes", primeCacheSlice)
	return outputInterval, nil
}

/*func mergeAndRemoveDuplicates(existing []int, newValues []int) []int { //關鍵
	existingMap := make(map[int]bool)
	merged := make([]int, 0, len(existing)+len(newValues))

	// 将 existing 切片中的元素添加到哈希集合中
	for _, val := range existing {
		existingMap[val] = true
		merged = append(merged, val)
	}

	// 将 newValues 切片中的不在哈希集合中的元素添加到 merged 切片中
	for _, val := range newValues {
		if !existingMap[val] {
			existingMap[val] = true
			merged = append(merged, val)
		}
	}

	return merged
}
func calculateAndGetInterface(start, end int) ([]int, []int, bool) { //当 cachedPrimes 不存在时，你试图将其强制转换为 []int 类型，
	//这会导致一个错误，因为 cachedPrimes 是 nil，而不能强制转换为切片---->解決:init一個初始cache
	primeInRange := calculator.Calculate(start, end)
	cachedPrimes, found := cacheObject.Get("primes")
	existingPrimes := cachedPrimes.([]int)
	return primeInRange, existingPrimes, found
}*/

/*func QueryCachePrimesInRange(start int, end int) ([]int, error) {
	cachedPrimes, _, found := calculateAndGetInterface(start, end)
	if !found {
		err := UpdateCachePrimeRange(start, end)
		if err != nil {
			return nil, err
		}
	}
	primesInRange := calculator.SearchInterval(cachedPrimes, start, end)
	return primesInRange, nil
}*/

// /辨認value與質數區間，有沒有可以對到的數字，全都有就不用更新，一旦缺少一個就更新
/*func UpdateCachePrimeRange(start int, end int) error {
	primeInRange, existingPrimes, _ := calculateAndGetInterface(start, end)
	updatedPrimes := mergeAndRemoveDuplicates(existingPrimes, primeInRange)
	cacheObject.Set("primes", updatedPrimes, cache.DefaultExpiration)
	return nil //找不到key
}*/
