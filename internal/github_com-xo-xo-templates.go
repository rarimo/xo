// Code generated by 'yaegi extract gitlab.com/rarify-protocol/xo/templates'. DO NOT EDIT.

package internal

import (
	"gitlab.com/rarify-protocol/xo/templates"
	"reflect"
)

func init() {
	Symbols["gitlab.com/rarify-protocol/xo/templates/templates"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BaseFuncs":       reflect.ValueOf(templates.BaseFuncs),
		"Errors":          reflect.ValueOf(templates.Errors),
		"Flags":           reflect.ValueOf(templates.Flags),
		"For":             reflect.ValueOf(templates.For),
		"GenType":         reflect.ValueOf(templates.GenType),
		"GenTypeKey":      reflect.ValueOf(templates.GenTypeKey),
		"Out":             reflect.ValueOf(templates.Out),
		"OutKey":          reflect.ValueOf(templates.OutKey),
		"Process":         reflect.ValueOf(templates.Process),
		"Register":        reflect.ValueOf(templates.Register),
		"Src":             reflect.ValueOf(templates.Src),
		"SrcKey":          reflect.ValueOf(templates.SrcKey),
		"Suffix":          reflect.ValueOf(templates.Suffix),
		"SuffixKey":       reflect.ValueOf(templates.SuffixKey),
		"Symbols":         reflect.ValueOf(templates.Symbols),
		"SymbolsKey":      reflect.ValueOf(templates.SymbolsKey),
		"TemplateType":    reflect.ValueOf(templates.TemplateType),
		"TemplateTypeKey": reflect.ValueOf(templates.TemplateTypeKey),
		"Types":           reflect.ValueOf(templates.Types),
		"Write":           reflect.ValueOf(templates.Write),
		"WriteFiles":      reflect.ValueOf(templates.WriteFiles),
		"WriteRaw":        reflect.ValueOf(templates.WriteRaw),

		// type definitions
		"EmittedTemplate": reflect.ValueOf((*templates.EmittedTemplate)(nil)),
		"ErrPostFailed":   reflect.ValueOf((*templates.ErrPostFailed)(nil)),
		"Template":        reflect.ValueOf((*templates.Template)(nil)),
		"TemplateSet":     reflect.ValueOf((*templates.TemplateSet)(nil)),
	}
}
