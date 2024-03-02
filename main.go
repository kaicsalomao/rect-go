package main

import "fmt"

// Define a struct Rectangle
type Rectangle struct {
	width, height int
}

// Calcula a area do Rectangle
func (r Rectangle) Area() int {
	return r.width * r.height
}

// Renderiza o Rectangle com a respectiva largura e altura
func (r Rectangle) Render() {
	block := "█"
	for i := 0; i <= r.width*2+2; i++ {
		fmt.Print(block)
	}
	fmt.Println()
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
	for i := 0; i <= r.width*2+2; i++ {
		fmt.Print(block)
	}
	fmt.Println()
}

// Verifica se é possivel conter um outro determinado Rectangle
func (r Rectangle) CanHold(other Rectangle) bool {
	if r.width > other.width && r.height > other.height {
		return true
	} else {
		return false
	}
}

// Renderiza um outro Rectangle dentro do mesmo, se for possivel
func (r Rectangle) RenderHold(other Rectangle) {
	if r.CanHold(other) {
		block := "█"
		for i := 0; i <= r.width*2+2; i++ {
			fmt.Print(block)
		}
		fmt.Println()
		for i := 0; i <= r.height; i++ {
			fmt.Print(block + " ")
			for j := 0; j <= r.width; j++ {
				if j == r.width {
					fmt.Println(block + " ")
				} else if j <= r.width/2+other.width/2 && j >= r.width/2-other.width/2 && i >= r.height/2-other.height/2 && i <= r.height/2+other.height/2 {
					fmt.Print("+ ")
				} else {
					fmt.Print("  ")
				}
			}
		}
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
