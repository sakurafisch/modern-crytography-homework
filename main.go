package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func birthday(n uint32, posibility float64) (uint32, error) {
	var k uint32 = n
	var tmp float64 = 1.0
	for i := uint32(0); i < k; i++ {
		tmp = tmp * float64(n-i) / float64(n)
		if tmp <= posibility {
			fmt.Println("We need to generate", i, "numbers per experiment for a successful rate of", (float64(1.0) - tmp))
			return i, nil
		}
	}
	return 0, errors.New("error while calling func birthday")
}

func birthday_uint32_max(posibility float64) (uint32, error) {
	var uint32_max uint32 = ^uint32(0)
	return birthday(uint32_max, posibility)
}

func try_experiment(try uint32, numbers_per_experiement uint32) int {
	counter := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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

		// reset m
		m = make(map[uint32]bool)
	}
	return counter
}

func main() {
	var arr []int
	var try uint32 = 1000
	var possibilities []float64
	for i := float64(0.99); i > 0; i -= float64(0.01) {
		possibilities = append(possibilities, i)
	}
	var generated_numbers_of_experiments []uint32
	var actual_possibilities_arr []float64

	for _, v := range possibilities {
		numbers_per_experiement, err := birthday_uint32_max(v)
		if err != nil {
			fmt.Println(err)
			return
		}

		generated_numbers_of_experiments = append(generated_numbers_of_experiments, numbers_per_experiement)

		for i := 0; i < 5; i++ {
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

		// compute average hint
		average := float64(sum) / float64(len(arr))
		fmt.Println("Average hint:", average)

		// compute average possibility
		possibility := average / (float64(try))
		fmt.Println("Average possibility:", possibility)

		// save actual average possibilitity
		actual_possibilities_arr = append(actual_possibilities_arr, possibility)

		// compute variance
		variance := float64(0)
		for _, value := range arr {
			variance += (float64(value) - average) * (float64(value) - average)
		}
		variance /= float64(len(arr))
		fmt.Println("Variance:", variance)
		fmt.Println("----------------------------------------------")

		// reset arr
		arr = []int{}
	}

	// Prepare data for drawing diagram
	var generated_numbers_of_experiments_for_draw []opts.LineData
	for _, v := range generated_numbers_of_experiments {
		generated_numbers_of_experiments_for_draw = append(generated_numbers_of_experiments_for_draw, opts.LineData{Value: v})
	}

	var theoretical_possibilities_for_draw []opts.LineData
	for i := len(possibilities) - 1; i > 0; i-- {
		theoretical_possibilities_for_draw = append(theoretical_possibilities_for_draw, opts.LineData{Value: possibilities[i]})
	}

	var actual_possibilities_for_draw []opts.LineData
	for _, v := range actual_possibilities_arr {
		actual_possibilities_for_draw = append(actual_possibilities_for_draw, opts.LineData{Value: v})
	}

	// Draw the diagram
	f, _ := os.Create("lines.html")

	line_theoretical := charts.NewLine()
	line_theoretical.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Birthday paradox",
		Subtitle: "Generated numbers by theoretical possibilities",
	}))
	line_theoretical.SetXAxis(theoretical_possibilities_for_draw)
	line_theoretical.AddSeries("Theoretical value", generated_numbers_of_experiments_for_draw)
	line_theoretical.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line_theoretical.Render(f)

	line_actual := charts.NewLine()
	line_actual.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Birthday paradox",
		Subtitle: "Generated numbers by actual possibilities",
	}))
	line_actual.SetXAxis(actual_possibilities_for_draw)
	line_actual.AddSeries("Actual hint", generated_numbers_of_experiments_for_draw)
	line_actual.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line_actual.Render(f)

}
