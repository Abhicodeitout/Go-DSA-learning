//Problem statement https://www.codechef.com/problems/LECANDY

//Solution:
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, N, C, sum int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &N)
		fmt.Fscan(reader, &C)
		sum = 0
		for j := 0; j < N; j++ {
			var candy int
			fmt.Fscan(reader, &candy)
			sum += candy
		}
		if sum <= C {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
	writer.Flush()
}
