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
			fmt.Println("we need to try ", i, " times for a successful rate of ", tmp)
			return i, nil
		}
	}
	return 0, errors.New("error while calling func birthday")
}

func Birthday_int() (uint32, error) {
	return birthday(uint32_max)
}

func main() {
	total_try, err := Birthday_int()
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := 0
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var try uint32 = 5000
	var itr uint32 = 0
	for ; itr < try; itr++ {
		m := make(map[uint32]bool)
		sign := false
		var i uint32 = 0
		for ; i < total_try; i++ {
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
	fmt.Println("hit ", counter, " times of ", try, " experiment")
}
