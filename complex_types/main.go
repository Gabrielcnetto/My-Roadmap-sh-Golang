package main

import "fmt"

func main() {
	//complex envia 2 numeros, index 1 - O Numero real | 2 - envia o valor imaginario
	/*
		c1 := complex(5, 7)
		c2 := 8 + 27i
		fmt.Println(c2)
		cadd := c1 + c2
		fmt.Println("sum:", cadd)
	*/
	c1 := complex(5, 2)
	c2 := 5 + 5i
	fmt.Printf("Valor c1:%v | valor c2: %v", c1, c2)
	sum := c1 + c2
	//eles somam reais com reais,e imaginarios com imaginarios e retornam
	/*
	 Valor c1:(5+2i) | valor c2: (5+5i)
	Soma:(10+7i)
	*/
	fmt.Printf("\nSoma:%v \n", sum)

}
