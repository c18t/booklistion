package controller

import (
	"github.com/c18t/booklistion/internal/core"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

type RootParams struct {
}

type RootController interface {
	core.Controller
	Params() *RootParams
}

type rootController struct {
	params *RootParams
}

func NewRootController(i do.Injector) (RootController, error) {
	return &rootController{
		params: &RootParams{},
	}, nil
}

func (c *rootController) Params() *RootParams {
	return c.params
}

func (c *rootController) Exec(cmd *cobra.Command, args []string) (err error) {
	err = cmd.Help()
	return
}
