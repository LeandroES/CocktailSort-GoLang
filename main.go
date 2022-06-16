package main

import "fmt"

func CocktailSort(a [9]int, n int) {
	swapped := true
	start := 0
	end := n - 1
	temp := 0

	for swapped == true {
		// restablecer Flag de swap al entrar en el bucle,
		//porque podría ser cierto desde una iteración anterior.
		swapped := false
		// bucle de izquierda a derecha igual que la ordenación de burbujas
		for i := start; i < end; i++ {
			if (a[i] > a[i+1]) {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
			}
		}
		//si no se mueve nada, entonces el array se ordena.
		if (!swapped) {
			break
		}
		// de lo contrario, restablece la bandera intercambiada para
		// que pueda ser utilizada en la siguiente etapa
		swapped = false
		// retroceder el punto final en uno, porque el elemento del
		// final está en su lugar correcto
		end--
		// de derecha a izquierda, haciendo la misma comparación que
		// en la etapa anterior
		for i := end - 1; i >= start; i-- {
			if (a[i] > a[i+1]) {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
			}
		}
		start++
	}

	for i := 0; i < n; i++ {
		//printf("%d ", a[i])
		fmt.Println(a[i])
	}
}

func main() {
	a := [9]int{1522425, 10, 5, 6, 165, 1, 84, 56, 1523}
	n := len(a)
	CocktailSort(a, n)
}