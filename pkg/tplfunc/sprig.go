package tplfunc

import (
	"maps"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var Sprig = []Func{func(funcMap template.FuncMap) {
	maps.Copy(funcMap, sprig.TxtFuncMap())
}}
