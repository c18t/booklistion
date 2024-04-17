---
name: 'command'
root: 'internal'
output: '.'
questions:
  name: 'enter new command name:'
  usecase:
    message: 'enter a list of subcommand names, separated by spaces or commas[\s,]:'
    initial: 'command'
---

# Variables

- command_camel: `{{ inputs.name | camel }}`
- command_pascal: `{{ inputs.name | pascal }}`
- command_snake: `{{ inputs.name | snake }}`
- subcommand_list: `{{ inputs.usecase | replace "\s" "," }}`

# `adapter/controller/{{ command_snake }}.go`

```go
package controller

import (
	"{{ go_module }}/internal/core"
	"{{ go_module }}/internal/usecase/port"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

type {{ command_pascal }}Params struct {
}

type {{ command_pascal }}Controller interface {
	core.Controller
	Params() *{{ command_pascal }}Params
}

type {{ command_camel }}Controller struct {
	bus    port.{{ command_pascal }}UseCaseBus `do:""`
	params *{{ command_pascal }}Params
}

func New{{ command_pascal }}Controller(i do.Injector) ({{ command_pascal }}Controller, error) {
	return &{{ command_snake }}Controller{
		bus:    do.MustInvoke[port.{{ command_pascal }}UseCaseBus](i),
		params: &{{ command_pascal }}Params{},
	}, nil
}

func (c *{{ command_snake }}Controller) Params() *{{ command_pascal }}Params {
	return c.params
}

func (c *{{ command_snake }}Controller) Exec(cmd *cobra.Command, args []string) (err error) {
  {{ for subcommand in subcommand_list | split ',' -}}
  {{ prefix := command_pascal + (subcommand | trim | pascal) -}}
	c.bus.Handle(&port.{{ prefix }}UseCaseInputData{})
  {{ end }}return
}

```

# `usecase/port/{{ command_snake }}.go`

```go
package port

import (
	"fmt"

	"{{ go_module }}/internal/core"
	"github.com/samber/do/v2"
)

type {{ command_pascal }}UseCaseInputData interface{}
type {{ command_pascal }}UseCaseOutputData interface{}

{{ for subcommand in (subcommand_list | split ',' )-}}
{{ prefix := command_pascal + (subcommand | trim | pascal) -}}
type {{ prefix }}UseCaseInputData struct {
	{{ command_pascal }}UseCaseInputData
}
type {{ prefix }}UseCaseOutpuData struct {
	{{ command_pascal }}UseCaseOutputData
}
type {{ prefix }}UseCase interface {
	core.UseCase
	Handle(input *{{ prefix }}UseCaseInputData)
}
{{ end }}
type {{ command_pascal }}UseCaseBus interface {
	Handle(input {{ command_pascal }}UseCaseInputData)
}
type {{ command_camel }}UseCaseBus struct {
  {{ for subcommand in (subcommand_list | split ',') -}}
	{{ prefix := command_pascal + (subcommand | trim | pascal) -}}
	{{ subcommand | trim | camel }} {{ prefix }}UseCase
  {{ end }}
}

func NewRootUseCaseBus(i do.Injector) ({{ command_pascal }}UseCaseBus, error) {
	return &{{ command_camel }}UseCaseBus{
		{{ for subcommand in (subcommand_list | split ',') -}}
		{{ prefix := command_pascal + (subcommand | trim | pascal) -}}
		{{ subcommand | trim | camel }}: do.MustInvoke[{{ prefix }}UseCase](i),
    {{ end }}}, nil
}

func (bus *{{ command_camel }}UseCaseBus) Handle(input {{ command_pascal }}UseCaseInputData) {
	switch data := input.(type) {
	{{ for subcommand in subcommand_list | split ',' -}}
	{{ prefix := command_pascal + (subcommand | trim | pascal) -}}
	case *{{ prefix }}UseCaseInputData:
		bus.{{ subcommand | trim | camel }}.Handle(data)
  {{ end }}default:
		panic(fmt.Errorf("handler for '%T' is not implemented", data))
	}
}

```

# `usecase/interactor/{{ command_snake }}.go`

```go
package interactor

import (
	"{{ go_module }}/internal/adapter/presenter"
	"{{ go_module }}/internal/usecase/port"
	"github.com/samber/do/v2"
)

{{ for subcommand in (subcommand_list | split ',') -}}
{{ prefix_pascal := command_pascal + (subcommand | trim | pascal) -}}
{{ prefix_camel := prefix_pascal | camel }}
type {{ prefix_camel }}Interactor struct {
	presenter presenter.{{ prefix_pascal }}Presenter
}

func New{{ prefix_pascal }}Interactor(i do.Injector) (port.{{ prefix_pascal }}UseCase, error) {
	return &{{ prefix_camel }}Interactor{
		presenter: do.MustInvoke[presenter.{{ prefix_pascal }}Presenter](i),
	}, nil
}

func (u *{{ prefix_camel }}Interactor) Handle(input *port.{{ prefix_pascal }}UseCaseInputData) {
	output := &port.{{ prefix_pascal }}UseCaseOutpuData{}

	u.presenter.Complete(output)
}
{{ end }}
```

# `adapter/presenter/{{ command_snake }}.go`

```go
package presenter

import (
	"{{ go_module }}/internal/usecase/port"
)

{{ for subcommand in (subcommand_list | split ',') -}}
{{ prefix_pascal := command_pascal + (subcommand | trim | pascal) -}}
{{ prefix_camel := prefix_pascal | camel }}
type {{ prefix_pascal }}Presenter interface {
	Complete(output *port.{{ prefix_pascal }}UseCaseOutpuData)
	Suspend(err error)
}

type {{ prefix_camel }}Presenter struct {
}

func New{{ prefix_pascal }}Presenter() {{ prefix_pascal }}Presenter {
	return &{{ prefix_camel }}Presenter{}
}

func (p *{{ prefix_camel }}Presenter) Complete(output *port.{{ prefix_pascal }}UseCaseOutpuData) {
}

func (p *{{ prefix_camel }}Presenter) Suspend(err error) {
}

{{ end }}
```
