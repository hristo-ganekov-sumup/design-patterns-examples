module github.com/hristo-ganekov-sumup/design-patterns-examples

go 1.16

replace (
	github.com/hristo-ganekov-sumup/design-patterns-examples/singleton => ./internal/singleton
	github.com/hristo-ganekov-sumup/design-patterns-examples/visitor => ./internal/visitor
)

require (
	github.com/hristo-ganekov-sumup/design-patterns-examples/singleton v0.0.0-00010101000000-000000000000
	github.com/hristo-ganekov-sumup/design-patterns-examples/visitor v0.0.0-00010101000000-000000000000
)
