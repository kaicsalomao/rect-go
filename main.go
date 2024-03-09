package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r Rectangle) Render() {
	block := "█"
	// Rendezirar o topo do retangulo
	for i := 0; i <= r.width*2+2; i++ {
		fmt.Print(block)
	}
	fmt.Println()

	// Renderizar as laterais, baseando-se na altura
	for i := 0; i <= r.height; i++ {
		fmt.Print(block + " ")
		for j := 0; j <= r.width; j++ {
			if j == r.width {
				fmt.Println(block + " ")
			} else {
				fmt.Print("  ")
			}
		}
	}
	// Renderizar a parte inferior do retangulo
	for i := 0; i <= r.width*2+2; i++ {
		fmt.Print(block)
	}
	fmt.Println()
}

func (r Rectangle) CanHold(other Rectangle) bool {
	if r.width > other.width && r.height > other.height {
		return true
	} else {
		return false
	}
}

func (r Rectangle) RenderHold(other Rectangle) {
	if r.CanHold(other) {
		block := "█"
		// Renderizar o topo do retangulo pai
		for i := 0; i <= r.width*2+2; i++ {
			fmt.Print(block)
		}
		fmt.Println()

		// Renderizar as laterais do retangulo pai
		for i := 0; i <= r.height; i++ {
			fmt.Print(block + " ")
			for j := 0; j <= r.width; j++ {
				if j == r.width {
					fmt.Println(block + " ")
					// Logica para encontrar o centro do retangulo pai e renderizar o retangulo filho
				} else if j <= r.width/2+other.width/2 && j >= r.width/2-other.width/2 && i >= r.height/2-other.height/2 && i <= r.height/2+other.height/2 {
					fmt.Print("+ ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		// Rendizar a parte inferior
		for i := 0; i <= r.width*2+2; i++ {
			fmt.Print(block)
		}
		fmt.Println()
	} else {
		fmt.Println("Não pode conter")
	}
}
func main() {
	fmt.Println("Rect1:")
	rect1 := Rectangle{width: 7, height: 8}
	rect1.Render()
	rect2 := Rectangle{width: 2, height: 2}
	fmt.Println("Rect2: ")
	rect2.Render()
	fmt.Println("Rect1 contendo rect2: ")
	rect1.RenderHold(rect2)
}
