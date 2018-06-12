package main
import "fmt"

func fibonacci(n int) int {

    //var result []int
	result := make([]int,n+1)
	//fmt.Println("result",result)

		result[0] = 0
		result[1] = 1
		for i:=2 ; i<= n;i++ {

			result[i] = result[i-1]+result[i-2]

		}

	return result[n]

}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    fmt.Println(fibonacci(n))
}