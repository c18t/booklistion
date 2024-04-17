package presenter

import (
	"fmt"

	"github.com/c18t/booklistion/internal/usecase/port"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

type RootCommandPresenter interface {
	Complete(output *port.RootCommandUseCaseOutpuData)
	Suspend(err error)
}

type rootCommandPresenter struct {
}

func NewRootCommandPresenter(i do.Injector) (RootCommandPresenter, error) {
	return &rootCommandPresenter{}, nil
}

func (p *rootCommandPresenter) Complete(output *port.RootCommandUseCaseOutpuData) {
	fmt.Printf("%v\n", output)
}

func (p *rootCommandPresenter) Suspend(err error) {
	cobra.CheckErr(err)
}
