package database

import (
	"test/calculator"
	"test/server"
)

type Prime struct {
	Value int
}

func QueryDBPrimesInRange(start int, end int) ([]int, error) { //從資料表輸出
	var primesInRange []int
	prime := calculator.Calculate(start, end)

	for _, solution := range prime {
		if solution >= start && solution <= end {
			primesInRange = append(primesInRange, solution)
		}
	}
	if len(primesInRange) == 0 {
		// 如果沒有找到質數，返回一個空的切片
		return []int{}, nil
	}

	return primesInRange, nil
}
func UpdateDBPrimeRange(start int, end int) error { //達成新增的問題，如果有就不動，如果沒有就儲存
	primeInRange := calculator.Calculate(start, end)

	var existingValues []int
	if err := server.DateBaseConnection.Table("primes").
		Pluck("value", &existingValues).Error; err != nil {
		return err
	}
	existingValuesMap := make(map[int]bool) //map放置int與確認用的標記符
	for _, tableValue := range existingValues {
		existingValuesMap[tableValue] = true
	}
	for _, confirmPrime := range primeInRange {
		if !existingValuesMap[confirmPrime] { //如果沒有
			if err := server.DateBaseConnection.Table("primes").
				Create(&Prime{Value: confirmPrime}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
