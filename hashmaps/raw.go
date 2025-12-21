//package main

/*
	func main() {
		//Entrada:
		//n := 8
		var a = []int{15, -2, 2, -8, 1, 7, 10, 23}
		x := 0
		resuls
		for _, item := range a {
			x += item

		}
		fmt.Println(x)
		// Saída: 5
		// Explicação: O maior subarray com
		// a soma 0 será -2 2 -8 1 7.
	}
*/
/*
func main() {
	arr := []int{15, -2, 2, -8, 1, 7, 10, 23}
	target := 0
	for start := 0; start < len(arr); start++ {
		sum := 0
		for end := start; end < len(arr); end++ {
			sum += arr[end]
			if sum == target {
				fmt.Printf("\nValor correto:: arr[%v:%v]\n", start, end+1)
				fmt.Println("encontrada:", arr[start:end+1])
				fmt.Println("________________________")
			} else {
				fmt.Printf("\nNulo: arr[%v:%v]\n", start, end+1)
			}
		}
	}
}
*/
