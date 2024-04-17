package interactor

import (
	"github.com/c18t/booklistion/internal/adapter/presenter"
	"github.com/c18t/booklistion/internal/usecase/port"
	"github.com/samber/do/v2"
)

type rootCommandInteractor struct {
	presenter presenter.RootCommandPresenter `do:""`
}

func NewRootCommandInteractor(i do.Injector) (port.RootCommandUseCase, error) {
	return &rootCommandInteractor{
		presenter: do.MustInvoke[presenter.RootCommandPresenter](i),
	}, nil
}

func (u *rootCommandInteractor) Handle(input *port.RootCommandUseCaseInputData) {
	output := &port.RootCommandUseCaseOutpuData{}
	output.Message = "root command called."
	u.presenter.Complete(output)
}
