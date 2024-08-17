package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}
type Lista struct {
	cabeca *No
}

func NovaLista() *Lista {
	return &Lista{}
}

func (l *Lista) InserirInicio (valor int) {
	novoNo := &No{valor: valor, proximo: l.cabeca}
	l.cabeca = novoNo
}

func (l *Lista) InserirFinal (valor int) {
	novoNo := &No{valor: valor}

	if l.cabeca == nil {
		l.cabeca = novoNo
	} else {
		atual := l.cabeca
		for atual.proximo != nil {
			atual = atual.proximo
		}
		atual.proximo = novoNo
	}
}

func (l *Lista) RemoverInicio () (int, bool) {
	// Lista Vazia
	if l.cabeca == nil {
		return 0, false 
	} else {
		elementoInicio := l.cabeca
		l.cabeca = l.cabeca.proximo
		return elementoInicio.valor, true
	}
}

func (l *Lista) RemoverFim () (int, bool) {
	// Lista Vazia
	if l.cabeca == nil {
		return 0, false 
	}

	// Lista tem apenas um elemento
	if l.cabeca.proximo == nil {
		valorRemovido := l.cabeca.valor
		l.cabeca = nil
		return valorRemovido, true
	}

	// Lista tem mais de um elemento
	atual := l.cabeca
	for atual.proximo.proximo != nil {
		atual = atual.proximo
	}

	// Atual é o penúltimo elemento
	valorRemovido := atual.proximo.valor
	atual.proximo = nil
	return valorRemovido, true
}

func (l *Lista) removerValor(valor int) bool {
    // Lista vazia
    if l.cabeca == nil {
        return false
    }

    // Valor é o primeiro elemento
    if l.cabeca.valor == valor {
        l.cabeca = l.cabeca.proximo
        return true
    }

    // O valor não é o primeiro elemento e ele está no meio da lista
    atual := l.cabeca
    for atual.proximo != nil && atual.proximo.valor != valor {
        atual = atual.proximo
    }

    // Se o próximo nó não for nulo, significa que o valor foi encontrado
    if atual.proximo != nil {
        atual.proximo = atual.proximo.proximo
        return true
    }

    // Valor não encontrado
    return false
}

func (l *Lista) buscarNo (valor int) (bool) {
	// Verifica se a lista está vazia
	if l.cabeca == nil {
		return false
	}

	// Verifica se o valor é o cabeça
	if l.cabeca.valor == valor {
		return true
	}

	// Verificar por toda a lista
	atual := l.cabeca
	for atual.proximo != nil {
		atual = atual.proximo
		if atual.valor == valor {
			return true
		}
	}

	return false
}

func (l *Lista) Imprimir() {
    atual := l.cabeca
    for atual != nil {
        fmt.Print(atual.valor, " ")
        atual = atual.proximo
    }
    fmt.Println(l.cabeca)
}


func main() {
	newLista := NovaLista()
	fmt.Println(newLista)
	newLista.InserirInicio(1)
	newLista.InserirFinal(2)
	newLista.InserirFinal(3)
	newLista.InserirFinal(4)
	newLista.InserirFinal(5)
	newLista.InserirFinal(6)
	newLista.InserirFinal(7)
	newLista.InserirFinal(8)
	newLista.InserirFinal(8)
	newLista.removerValor(9)
	newLista.Imprimir()
	fmt.Println(newLista.buscarNo(2))
	fmt.Println(newLista.buscarNo(25))
}