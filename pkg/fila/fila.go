package main

import "fmt"

type Fila struct {
	elementos []int
}

func inicializaFila() *Fila {
	return &Fila{elementos: []int{}}
}

func (f *Fila) Enqueue(valor int) {
	f.elementos = append(f.elementos, valor)
}

func (f *Fila) Dequeue() (int, bool) {

	if len(f.elementos) == 0 {
		return 0, false
	}

	primeiroElemento := f.elementos[0]
	f.elementos = f.elementos[1:]
	return primeiroElemento, true
}

func (f *Fila) Peek () (int, bool){
	if len(f.elementos) == 0 {
		return 0, false
	}

	return f.elementos[0], true
}

func (f *Fila) IsEmpty () (bool){
	return len(f.elementos) == 0
}

func main() {
	newFila := inicializaFila()
	fmt.Println(*newFila)
	newFila.Enqueue(1)
	newFila.Enqueue(2)
	newFila.Enqueue(3)
	fmt.Println(*newFila)
	fmt.Println(newFila.Dequeue())
	fmt.Println(*newFila)
	fmt.Println(newFila.Peek())
	newFila.Enqueue(8)
	newFila.Enqueue(5)
	newFila.Enqueue(9)
	fmt.Println(*newFila)
	fmt.Println(newFila.IsEmpty())
}