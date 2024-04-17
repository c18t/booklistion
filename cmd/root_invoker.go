package cmd

import (
	"github.com/c18t/booklistion/internal/adapter/controller"
	"github.com/c18t/booklistion/internal/core"
	"github.com/c18t/booklistion/internal/inject"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

func createRootCommand() core.RunEFunc {
	cmd, err := do.Invoke[controller.RootController](inject.Injector)
	cobra.CheckErr(err)
	return cmd.Exec
}
