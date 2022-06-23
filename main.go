package main

import (
	"fmt"
	"io"
	_ "log"
	"os"
	"time"
)

func CocktailSort(a []int) []int {
	n := len(a)
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
			if a[i] > a[i+1] {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
			}
		}
		//si no se mueve nada, entonces el array se ordena.
		if !swapped {
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
			if a[i] > a[i+1] {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				swapped = true
			}
		}
		start++
	}
	return a
}
func main() {
	file, err := os.Open("50000.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var perline int
	var nums []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	inicio := time.Now()
	CocktailSort(nums)
	duracion := time.Since(inicio)
	fmt.Println("Tiempo en Microsegundos: ", duracion.Microseconds())

}
