package inject

import (
	"github.com/c18t/booklistion/internal/adapter/controller"
	"github.com/c18t/booklistion/internal/adapter/presenter"
	"github.com/c18t/booklistion/internal/usecase/interactor"
	"github.com/c18t/booklistion/internal/usecase/port"
	"github.com/samber/do/v2"
)

var InjectorRoot = AddRootProvider()

func AddRootProvider() *do.RootScope {
	// adapter/controller
	do.Provide[controller.RootController](Injector, controller.NewRootController)

	// usecase/port
	do.Provide[port.RootUseCaseBus](Injector, port.NewRootUseCaseBus)

	// usecase/intractor
	do.Provide[port.RootCommandUseCase](Injector, interactor.NewRootCommandInteractor)

	// adapter/presenter
	do.Provide[presenter.RootCommandPresenter](Injector, presenter.NewRootCommandPresenter)

	return Injector
}
