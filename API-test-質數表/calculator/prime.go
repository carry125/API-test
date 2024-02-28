package calculator

func IsPrime(numberInterval int) bool {
	if numberInterval < 2 {
		return false
	}
	for i := 2; i < numberInterval; i++ {
		if numberInterval%i == 0 {
			return false
		}
	}
	return true
}
func Calculate(numStart int, numEnd int) []int {
	var numberGroup []int
	if numStart > numEnd {
		numStart, numEnd = Swap(numStart, numEnd)
	}
	// 進行計算
	for i := numStart; i <= numEnd; i++ {
		if IsPrime(i) {
			numberGroup = append(numberGroup, i)
		}
	}
	return numberGroup
}
func Swap(numStart int, numEnd int) (start int, end int) {
	temp := numStart
	numStart = numEnd
	numEnd = temp

	start = numStart
	end = numEnd
	return start, end
}
func SearchInterval(cacheprime interface{}, start int, end int) []int {
	primeTable := cacheprime.([]int) //移置PRIME
	var primesInRange []int
	for _, prime := range primeTable {
		if prime >= start && prime <= end {
			primesInRange = append(primesInRange, prime)
		}
	}
	return primesInRange
}
func ComparePrime(input int, primeslice []int) bool { //確認每個start到end的數，484cache裡存的質數
	for _, value := range primeslice {
		if input == value {
			return true
		}
	}
	return false
}
