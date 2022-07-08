package main

import "fmt"

func average(xs []float64){
	var avg float64
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	fmt.Printf("平均值是：%f", avg)
}

func main() {
	average([]float64{1.0, 2.0, 3.0})
}