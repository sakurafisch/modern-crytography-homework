package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func birthday(n uint32) (uint32, error) {
	var k uint32 = n
	var tmp float64 = 1.0
	for i := uint32(0); i < k; i++ {
		tmp = tmp * float64(n-i) / float64(n)
		if tmp <= 0.5 {
			fmt.Println("We need to try", i, "times for a successful rate of", tmp)
			return i, nil
		}
	}
	return 0, errors.New("error while calling func birthday")
}

func birthday_uint32_max() (uint32, error) {
	var uint32_max uint32 = ^uint32(0)
	return birthday(uint32_max)
}

func try_experiment(try uint32, numbers_per_experiement uint32) int {
	counter := 0
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for itr := uint32(0); itr < try; itr++ {
		m := make(map[uint32]bool)
		sign := false
		for i := uint32(0); i < numbers_per_experiement; i++ {
			tmp := r.Uint32()
			if m[tmp] == true {
				sign = true
				counter++
				break
			}
			m[tmp] = true
		}
		if sign {
			sign = false
		}
		m = make(map[uint32]bool)
	}
	return counter
}

func main() {
	var arr []int
	var try uint32 = 5000
	numbers_per_experiement, err := birthday_uint32_max()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 30; i++ {
		counter := try_experiment(try, numbers_per_experiement)

		fmt.Println("hit", counter, "times of", try, "experiments")
		arr = append(arr, counter)
	}
	fmt.Println("----------------------------------------------")

	fmt.Println("  index            try            hit")
	fmt.Println("----------------------------------------------")
	sum := 0
	for index, counter := range arr {
		sum += counter
		if index+1 < 10 {
			fmt.Println("  ", index+1, "            ", try, "              ", counter)
		} else {
			fmt.Println(" ", index+1, "            ", try, "            ", counter)
		}
	}
	fmt.Println("----------------------------------------------")

	// 计算平均值
	average := float64(sum) / float64(len(arr))
	fmt.Println("Average hint:", average)

	// 计算平均概率
	posibility := average / (float64(try))
	fmt.Println("Average posibility:", posibility)

	// 计算方差
	variance := float64(0)
	for _, value := range arr {
		variance += (float64(value) - average) * (float64(value) - average)
	}
	variance /= float64(len(arr))
	fmt.Println("Variance:", variance)
}
