package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	if n := rand.Int31n(10); n == 0 { // nはelseのブロックまで有効になる
		fmt.Println("少し小さすぎます:", n)
	} else if n > 5 {
		fmt.Println("大きすぎます:", n)
	} else {
		fmt.Println("いい感じの数字です:", n)
	}
}
