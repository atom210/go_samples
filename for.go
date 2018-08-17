package main

import "fmt"

func Max(l []int32) int32 {
	max := l[0]
	for _, v := range l {
		if max < v {
			max = v
		}
	}
	return max
}

func Min(l []int32) int32 {
	min := l[0]
	for _, v := range l {
		if min > v {
			min = v
		}
	}
	return min
}

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0.0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}

	return
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	ary := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", ary)

}
