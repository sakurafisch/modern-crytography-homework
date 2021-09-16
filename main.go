package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var uint32_max uint32 = ^uint32(0)

func birthday(n uint32) (uint32, error) {
	var k uint32 = n
	var tmp float64 = 1.0
	var i uint32 = 0
	for ; i < k; i++ {
		tmp = tmp * float64(n-i) / float64(n)
		if tmp <= 0.5 {
			// fmt.Println(tmp)
			// fmt.Println(i)
			fmt.Println("we need to try", i, "times for a successful rate of", tmp)
			return i, nil
		}
	}
	return 0, errors.New("error while calling func birthday")
}

func birthday_int() (uint32, error) {
	return birthday(uint32_max)
}

func try_experiment(try uint32, numbers_per_experiement uint32) int {
	counter := 0
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var itr uint32 = 0
	for ; itr < try; itr++ {
		m := make(map[uint32]bool)
		sign := false
		var i uint32 = 0
		for ; i < numbers_per_experiement; i++ {
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
	numbers_per_experiement, err := birthday_int()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 30; i++ {
		counter := try_experiment(try, numbers_per_experiement)

		fmt.Println("hit", counter, "times of", try, "experiments")
		arr = append(arr, counter)
	}
	fmt.Println("  index		  try		  hit")
	fmt.Println("----------------------------------------------")

	for index, counter := range arr {
		if index+1 < 10 {
			fmt.Println("  ", index+1, "            ", try, "            ", counter)
		} else {
			fmt.Println(" ", index+1, "            ", try, "            ", counter)
		}
	}

}
