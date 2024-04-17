package inject

import (
	"github.com/c18t/booklistion/internal/adapter/controller"
	"github.com/samber/do/v2"
)

var Injector = NewInjector()

func NewInjector() *do.RootScope {
	injector := do.New()

	// adapter/controller
	do.Provide[controller.RootController](injector, controller.NewRootController)

	return injector
}
