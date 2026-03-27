package main

import "fmt"

// ===================== ARREGLO =====================
var arreglo = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func recorrerFor() {
	lon := len(arreglo)
	for i := 0; i < lon; i++ {
		fmt.Println(arreglo[i])
	}
}

func recorrerForEach() {
	for i, val := range arreglo {
		fmt.Println(i, val)
	}
}

func impares() {
	for i := range arreglo {
		if arreglo[i]%2 != 0 {
			arreglo[i] = 0
		}
	}
	fmt.Println("Arreglo con impares en 0:")
	recorrerFor()
}

func multIndice() {
	for i := range arreglo {
		arreglo[i] = arreglo[i] * i
	}
	fmt.Println("Arreglo multiplicado por índice:")
	recorrerFor()
}

func busqueda() {
	valor := 4
	encontrado := false

	for i := range arreglo {
		if arreglo[i] == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("Valor encontrado")
	} else {
		fmt.Println("El valor NO se encuentra en el arreglo")
	}
}

// ===================== MATRIZ =====================
var matriz = [3][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

func imprimirTabla() {
	for i := 0; i < 3; i++ {
		fila := ""
		for j := 0; j < 3; j++ {
			fila += fmt.Sprintf("%d ", matriz[i][j])
		}
		fmt.Println(fila)
	}
}

func sumar() {
	suma := 0
	for i := 0; i < 3; i++ {
		sumafila := 0
		for j := 0; j < 3; j++ {
			sumafila += matriz[i][j]
		}
		suma += sumafila
	}
	fmt.Println("La suma total de la matriz es:", suma)
}

func intercambiar(f1, f2 int) {
	for j := 0; j < 3; j++ {
		matriz[f1][j], matriz[f2][j] = matriz[f2][j], matriz[f1][j]
	}
	fmt.Println("Matriz después de intercambiar filas:")
	imprimirTabla()
}

// ===================== MAIN =====================
func main() {

	// ARREGLO
	fmt.Println("Recorrido con for:")
	recorrerFor()

	fmt.Println("\nRecorrido con forEach:")
	recorrerForEach()

	fmt.Println()
	impares()

	fmt.Println()
	multIndice()

	fmt.Println()
	busqueda()

	// MATRIZ
	fmt.Println("\nMatriz original:")
	imprimirTabla()

	fmt.Println()
	sumar()

	fmt.Println()
	intercambiar(0, 2)
}
