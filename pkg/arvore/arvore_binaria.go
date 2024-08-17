package main

import "fmt"

// Estrutura de um nó
type Node struct {
	valor    int
	esquerda *Node
	direita  *Node
}

// Estrutura da árvore
type Tree struct {
	raiz *Node
}

// Método para inserir um valor na árvore
func (t *Tree) Inserir(valor int) {
	novoNo := &Node{valor: valor}
	if t.raiz == nil {
		t.raiz = novoNo
	} else {
		inserirNo(t.raiz, novoNo)
	}
}

func inserirNo(atual, novoNo *Node) {
	if novoNo.valor < atual.valor {
		if atual.esquerda == nil {
			atual.esquerda = novoNo
		} else {
			inserirNo(atual.esquerda, novoNo)
		}
	} else {
		if atual.direita == nil {
			atual.direita = novoNo
		} else {
			inserirNo(atual.direita, novoNo)
		}
	}
}

// Método para buscar um valor na árvore
func (t *Tree) Buscar(valor int) *Node {
	return buscarNo(t.raiz, valor)
}

func buscarNo(atual *Node, valor int) *Node {
	if atual == nil || atual.valor == valor {
		return atual
	}
	if valor < atual.valor {
		return buscarNo(atual.esquerda, valor)
	}
	return buscarNo(atual.direita, valor)
}

// Método para remover um valor da árvore
func (t *Tree) Remover(valor int) {
	t.raiz = removerNo(t.raiz, valor)
}

func removerNo(atual *Node, valor int) *Node {
	if atual == nil {
		return nil
	}
	if valor < atual.valor {
		atual.esquerda = removerNo(atual.esquerda, valor)
		return atual
	}
	if valor > atual.valor {
		atual.direita = removerNo(atual.direita, valor)
		return atual
	}

	// Se o nó a ser removido tem dois filhos
	if atual.esquerda != nil && atual.direita != nil {
		substituto := buscarMinimo(atual.direita)
		atual.valor = substituto.valor
		atual.direita = removerNo(atual.direita, substituto.valor)
		return atual
	}

	// Se o nó a ser removido tem um ou nenhum filho
	if atual.esquerda != nil {
		return atual.esquerda
	}
	return atual.direita
}

func buscarMinimo(atual *Node) *Node {
	for atual.esquerda != nil {
		atual = atual.esquerda
	}
	return atual
}

// Método para percorrer a árvore em ordem
func (t *Tree) EmOrdem() {
	emOrdem(t.raiz)
}

func emOrdem(atual *Node) {
	if atual != nil {
		emOrdem(atual.esquerda)
		fmt.Println(atual.valor)
		emOrdem(atual.direita)
	}
}

func main() {
	// Criando uma nova árvore
	arvore := &Tree{}

	// Inserindo valores na árvore
	arvore.Inserir(5)
	arvore.Inserir(3)
	arvore.Inserir(7)
	arvore.Inserir(2)
	arvore.Inserir(4)
	arvore.Inserir(6)
	arvore.Inserir(8)

	// Imprimindo a árvore em ordem
	fmt.Println("Árvore em ordem:")
	arvore.EmOrdem()

	// Buscando um valor na árvore
	fmt.Println("\nBuscando valor 4:")
	no := arvore.Buscar(4)
	if no != nil {
		fmt.Printf("Valor %d encontrado na árvore.\n", no.valor)
	} else {
		fmt.Println("Valor não encontrado na árvore.")
	}

	// Removendo um valor da árvore
	fmt.Println("\nRemovendo valor 5 (raiz):")
	arvore.Remover(5)
	arvore.EmOrdem()

	// Removendo uma folha
	fmt.Println("\nRemovendo valor 8 (folha):")
	arvore.Remover(8)
	arvore.EmOrdem()

	// Removendo um nó com um filho
	fmt.Println("\nRemovendo valor 6 (nó com um filho):")
	arvore.Remover(6)
	arvore.EmOrdem()
}
