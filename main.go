package main

import (
	"fmt"
	"github.com/hristo-ganekov-sumup/design-patterns-examples/internal/singleton"
	"github.com/hristo-ganekov-sumup/design-patterns-examples/internal/visitor"
	"github.com/hristo-ganekov-sumup/design-patterns-examples/internal/simplefactory"
)


func main() {
	//Visitor
	square := &visitor.Square{Side: 2}
	circle := &visitor.Circle{Radius: 3}
	rectangle := &visitor.Rectangle{L: 2, B: 3}

	areaCalculator := &visitor.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	//Singleton
	for i := 0; i < 3; i++ {
		go singleton.GetInstance()
	}
	fmt.Scanln()

	//SimpleFactory
	ak47, _ := simplefactory.GetGun("ak47")
	musket, _ := simplefactory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g simplefactory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}