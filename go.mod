module github.com/hristo-ganekov-sumup/design-patterns-examples

go 1.16

replace (
	github.com/hristo-ganekov-sumup/design-patterns-examples/internal/singleton => ./internal/singleton
	github.com/hristo-ganekov-sumup/design-patterns-examples/internal/visitor => ./internal/visitor
	github.com/hristo-ganekov-sumup/design-patterns-examples/internal/simplefactory => ./internal/simplefactory
)

require (
	github.com/hristo-ganekov-sumup/design-patterns-examples/internal/singleton v0.0.0-00010101000000-000000000000
	github.com/hristo-ganekov-sumup/design-patterns-examples/internal/visitor v0.0.0-00010101000000-000000000000
)
