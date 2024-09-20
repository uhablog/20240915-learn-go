package main

import "fmt"

func main() {
	// 配列の宣言
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}

	// 配列の比較
	fmt.Println("配列の比較")
	fmt.Println(x == y)
	fmt.Println()

	// 構造体の宣言
	type person struct {
		name string
		age  int
	}

	// 構造体の初期化
	bob := person{
		name: "bob",
		age:  40,
	}
	// 構造体の値にアクセス
	fmt.Println("構造体の値にアクセス")
	fmt.Println(bob.name)
	fmt.Println(bob.age)
	fmt.Println()

	// 無名構造体
	printAnonymosStruct()
}

func printAnonymosStruct() {
	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}

	fmt.Println("無名構造体の値を出力")
	fmt.Println(pet.name)
	fmt.Println(pet.kind)
}
