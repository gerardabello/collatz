package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	n := big.NewInt(int64(number))

	fmt.Println(n)
	for {
		n = collatz(n)
		fmt.Println(n)
		if n.BitLen() == 1 {
			break
		}
	}

}

var one *big.Int = big.NewInt(1)
var two *big.Int = big.NewInt(2)
var three *big.Int = big.NewInt(3)

func collatz(n *big.Int) *big.Int {
	if n.Bit(0) == 0 {
		n.Div(n, two)
	} else {
		n.Mul(n, three)
		n.Add(n, one)
	}
	return n
}
