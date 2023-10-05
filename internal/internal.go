// Package internal contains internal code for xo.
package internal

import (
	"reflect"
)

// Symbols are extracted (generated) symbols from the types package.
//
// go list ./... |grep -v internal|tail -n +2|sed -e 's%^%//go:generate yaegi extract %'
//
//go:generate yaegi extract github.com/goccy/go-yaml
//go:generate yaegi extract github.com/kenshaw/inflector
//go:generate yaegi extract github.com/kenshaw/snaker
//go:generate yaegi extract golang.org/x/tools/imports
//go:generate yaegi extract mvdan.cc/gofumpt/format
//go:generate yaegi extract github.com/rarimo/xo/cmd
//go:generate yaegi extract github.com/rarimo/xo/loader
//go:generate yaegi extract github.com/rarimo/xo/models
//go:generate yaegi extract github.com/rarimo/xo/templates
//go:generate yaegi extract github.com/rarimo/xo/templates/createdbtpl
//go:generate yaegi extract github.com/rarimo/xo/templates/dottpl
//go:generate yaegi extract github.com/rarimo/xo/templates/gotpl
//go:generate yaegi extract github.com/rarimo/xo/templates/jsontpl
//go:generate yaegi extract github.com/rarimo/xo/templates/yamltpl
//go:generate yaegi extract github.com/rarimo/xo/types
var Symbols map[string]map[string]reflect.Value = make(map[string]map[string]reflect.Value)
