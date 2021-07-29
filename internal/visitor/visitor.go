package visitor

type visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}