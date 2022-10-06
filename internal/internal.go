// Package internal contains internal code for xo.
package internal

import (
	"reflect"
)

// Symbols are extracted (generated) symbols from the types package.
//
//go:generate yaegi extract github.com/goccy/go-yaml
//go:generate yaegi extract github.com/kenshaw/inflector
//go:generate yaegi extract github.com/kenshaw/snaker
//go:generate yaegi extract golang.org/x/tools/imports
//go:generate yaegi extract mvdan.cc/gofumpt/format
//
// go list ./... |grep -v internal|tail -n +2|sed -e 's%^%//go:generate yaegi extract %'
//
//go:generate yaegi extract gitlab.com/rarify-protocol/cmd
//go:generate yaegi extract gitlab.com/rarify-protocol/loader
//go:generate yaegi extract gitlab.com/rarify-protocol/models
//go:generate yaegi extract gitlab.com/rarify-protocol/templates
//go:generate yaegi extract gitlab.com/rarify-protocol/templates/createdbtpl
//go:generate yaegi extract gitlab.com/rarify-protocol/templates/dottpl
//go:generate yaegi extract gitlab.com/rarify-protocol/templates/gotpl
//go:generate yaegi extract gitlab.com/rarify-protocol/templates/jsontpl
//go:generate yaegi extract gitlab.com/rarify-protocol/templates/yamltpl
//go:generate yaegi extract gitlab.com/rarify-protocol/types
var Symbols map[string]map[string]reflect.Value = make(map[string]map[string]reflect.Value)
