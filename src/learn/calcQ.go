package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func calcQuality(min, max, equipQuality, continueMatNumb int32, jewelQualityMap map[int32]int32) int32 {
	base := make(map[int32]float32)
	for i := min; i <= max; i++ {
		if i == equipQuality {
			base[i] = float32(jewelQualityMap[i]) + float32(continueMatNumb)
		} else {
			base[i] = float32(jewelQualityMap[i])
		}
	}
	var reduce float32 = 0.5
	var sum float32 = 0
	var sumWeight float32 = 0.0
	tem := make(map[int32]float32)
	for i := min; i <= max; i++ {
		sum += base[i]
		tem[i] = sum * float32(max-i) * reduce
		sumWeight += base[i] + tem[i]
		base[i] = base[i] + tem[i]
	}
	randRatio := rand.Float32() * sumWeight
	var temRation float32 = 0.0
	fmt.Printf("allllll    :%v\n", sumWeight)
	for i := min; i <= max; i++ {
		fmt.Printf("oneeeee     : %v %v ** %v \n", i, base[i], base[i]/sumWeight)
	}

	for i := min; i <= max; i++ {
		temRation += base[i]
		if randRatio < temRation {
			return i
		}
	}
	return min
}

func CeilToInt32(v float32) int32 {
	return int32(math.Ceil(float64(v)))
}

func main() {
	q := calcQuality(1, 4, 2, 4, map[int32]int32{1: 1, 3: 1, 4: 1})
	fmt.Printf("get eeeeeeeeeee  %v", q)
	lastCost := CeilToInt32(float32(500 * 100.0 / (100 + 1)))
	println("IMY********************", lastCost)
	fmt.Printf("IMY******************** %v \n", strings.Split("5:6:7", ":"))
	fmt.Printf("IMY******************** %v \n", strings.Split("8", ":"))
	fmt.Printf("IMY******************** %v \n", int64(math.Ceil(float64(50)*float64(7197/3600.0))))
	var duration int32 = 7197
	A := float64(float64(duration) / 3600.0)
	fmt.Printf("IMY********************%v \n", A)

	fmap := make(map[int32]float64)
	fmap[2] = 2.00
	fmap[3] = 2.01
	fmt.Printf("IMY*********%v %v \n", math.Ceil(fmap[1]), math.Ceil(fmap[1]))
	fmt.Printf("IMY*********%v %v \n", math.Ceil(fmap[2]), math.Ceil(fmap[2]))
	fmt.Printf("IMY*********%v %v \n", math.Ceil(fmap[3]), math.Ceil(fmap[3]))

	vvv := make(map[int]map[int]int)

	if _, ok := vvv[1]; !ok {
		vvv[1] = make(map[int]int)
	}
	vvv[1][1] = 1
	fmt.Printf("IMY*********%v %v \n", math.Ceil(fmap[3]), math.Ceil(fmap[3]))

	var nonInitSlice []int
	nonInitSlice = append(nonInitSlice, 111)
	nonInitSlice = append(nonInitSlice, 2222)
	fmt.Printf("IMY************** %v", nonInitSlice)

}
