package main

import (
	"fmt"
	"github.com/hristo-ganekov-sumup/design-patterns-examples/singleton"
	"github.com/hristo-ganekov-sumup/design-patterns-examples/visitor"
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
}