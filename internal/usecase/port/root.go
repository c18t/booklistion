package port

import (
	"fmt"

	"github.com/c18t/booklistion/internal/core"
	"github.com/samber/do/v2"
)

type RootUseCaseInputData interface{}
type RootUseCaseOutputData interface{}

type RootCommandUseCaseInputData struct {
	RootUseCaseInputData
}
type RootCommandUseCaseOutpuData struct {
	RootUseCaseOutputData
  Message string
}
type RootCommandUseCase interface {
	core.UseCase
	Handle(input *RootCommandUseCaseInputData)
}

type RootUseCaseBus interface {
	Handle(input RootUseCaseInputData)
}
type rootUseCaseBus struct {
	command RootCommandUseCase `do:""`

}

func NewRootUseCaseBus(i do.Injector) (RootUseCaseBus, error) {
	return &rootUseCaseBus{
		command: do.MustInvoke[RootCommandUseCase](i),
		}, nil
}

func (bus *rootUseCaseBus) Handle(input RootUseCaseInputData) {
	switch data := input.(type) {
	case *RootCommandUseCaseInputData:
		bus.command.Handle(data)
	default:
		panic(fmt.Errorf("handler for '%T' is not implemented", data))
	}
}
