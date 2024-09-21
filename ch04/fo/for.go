package main

import "fmt"

func main() {
	fmt.Println("標準的なfor")
	basic()
	fmt.Println()
	fmt.Println("条件のみのfor")
	only()
	fmt.Println()
	fmt.Println("無限ループ(breakで抜ける)")
	infinity()
	fmt.Println()
	fmt.Println("rangeループ(slice)")
	forRangeSlice()
	fmt.Println()
	fmt.Println("rangeループ(map)")
	forRangeMap()
}

func basic() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func only() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func infinity() {
	i := 1
	for {
		fmt.Println("hello")
		i = i + 1
		if i > 10 {
			break
		}
	}
}

func forRangeSlice() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		// i(index)にはインデックスが、v(value)には値が入る
		fmt.Println(i, v)
	}
}

func forRangeMap() {
	uniqueName := map[string]bool{"ハナコ": true, "太郎": true, "一郎": true}
	for k, v := range uniqueName {
		// k(key)にはキーが、vには値が入る
		fmt.Println(k, v)
	}
}
