package main

import "fmt"

// Criando estrutura da pilha
type Pilha struct {
	elementos []int 
} 

// Inicializando Pilha
func NovaPilha() *Pilha{
	return &Pilha{elementos: []int{}}
}

// Operação Push
func (p *Pilha) Push (valor int) {
	p.elementos = append(p.elementos, valor)
}

func (p *Pilha) Pop () (int, bool) {
	// Pilha Vazia
	if len(p.elementos) == 0 {
		return 0, false 
	}

	topo := p.elementos[len(p.elementos) - 1]
	p.elementos = p.elementos[:len(p.elementos) - 1]
	return topo, true
}

func (p *Pilha) Peek () (int, bool) {
	// Pilha Vazia
	if len(p.elementos) == 0 {
		return 0, false 
	}

	return p.elementos[len(p.elementos) - 1], true
}

func (p *Pilha) isEmpty () bool {
	return len(p.elementos) == 0
}


func main() {
	newPilha := NovaPilha();
	newPilha.Push(1)
	fmt.Println(*newPilha)
	newPilha.Push(2)
	fmt.Println(*newPilha)
	newPilha.Push(3)
	fmt.Println(*newPilha)
	newPilha.Pop()
	fmt.Println(*newPilha)
	fmt.Println(newPilha.Peek())
	fmt.Println(newPilha.isEmpty())
}