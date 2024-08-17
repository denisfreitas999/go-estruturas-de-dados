package main

import (
	"fmt"
	"sort"
)

func somaElementosArray(array[] int) int {
	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}
	return soma
}

func ordenaArrayInteiros(array[5] int) [5]int {
	slice := array[:]
	sort.Ints(slice)
	var arrayOrdenado [5]int
	copy(arrayOrdenado[:], slice)
	return arrayOrdenado
}

func inverterSlice(slice[] int) []int  {
	n := len(slice)

	for i := 0; i < n/2; i++ {
        slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
    }

	return slice;
}

func somaElementosMatriz(matriz[3][3] int) int {
	soma := 0
	for i := 0; i < 3; i++ {
		for j := 0; j  < 3; j ++ {
			soma += matriz[i][j]
		}
	}
	return soma
}

func main() {
	// Questão 01
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Questão 01:", array)

	// Questão 02
	array[2] = 100
	fmt.Println("Questão 02:", array)

	// Questão 03
	fmt.Println("Questão 03:", somaElementosArray(array[:]))

	// Questão 04
	slice := []int {array[0], array[1], array[2]}
	fmt.Println("Questão 04:", slice)

	// Questão 05
	slice_02 := []string{}
	slice_02 = append(slice_02, "Go")
	slice_02 = append(slice_02, "Python")
	slice_02 = append(slice_02, "Java")
	fmt.Println("Questão 05:", slice_02)

	// Questão 06
	slice_03 := []int{1, 2, 3, 4, 5}
	// slice[:2]: Obtém os primeiros dois elementos do slice [1, 2].
	// slice[3:]: Obtém todos os elementos após o terceiro elemento, ou seja, [4, 5].
	slice_03 = append(slice_03[:2], slice_03[3:]...)
	fmt.Println("Questão 06:", slice_03)

	// Questão 07
	slice_04 := []int{10, 20, 30, 40, 50}
	slice_05 := []int{}
	slice_05 = append(slice_04[:2], slice_05...)
	fmt.Println("Questão 07:", slice_05)
	slice_06 := slice_04[3:]
	fmt.Println("Questão 07:", slice_06)

	// Questão 08
	array2 := [5]int{5, 3, 8, 1, 2} 
	fmt.Println("Questão 08:", ordenaArrayInteiros(array2))

	// Questão 09
	slice_07 := []int {1,2,3,5,8}
	fmt.Println("Questão 09:", inverterSlice(slice_07))

	// Questão 10
	matriz := [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println("Questão 10:", somaElementosMatriz(matriz))
}
