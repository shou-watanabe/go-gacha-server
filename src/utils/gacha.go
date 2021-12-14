package utils

import (
	"math/rand"
)

func WeightPick(array []int) int {
	totalWeight := 0
	pick := 0

	// トータルの重み計算
	for i := 0; i < len(array); i++ {
		totalWeight += array[i]
	}

	rnd := rand.Intn(totalWeight)

	for i := 0; i < len(array); i++ {
		if rnd < array[i] {
			// 抽選対象決定
			pick = i
			break
		}
		// 次の対象を調べる
		rnd -= array[i]
	}

	return pick
}
