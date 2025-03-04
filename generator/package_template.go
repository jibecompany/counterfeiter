package generator

import (
	"strings"
	"text/template"
)

var packageFuncs = template.FuncMap{
	"ToLower":  strings.ToLower,
	"UnExport": unexport,
	"Replace":  strings.Replace,
	"Generate": func(suffix string) string { return suffix + ":generate" }, // yes, this seems insane but ensures that we can use `go generate ./...` from the main package
}

const packageTemplate string = `{{.Header}}// Code generated by counterfeiter. DO NOT EDIT.
package {{.DestinationPackage}}

import (
	{{- range $index, $import := .Imports.ByAlias}}
	{{$import}}
	{{- end}}
)

//{{Generate "go"}} go run github.com/jibecompany/counterfeiter/v6 -generate
//{{Generate "counterfeiter"}} . {{.Name}}

// {{.Name}} is a generated interface representing the exported functions
// in the {{.TargetPackage}} package.
type {{.Name}} interface {
  {{- range .Methods}}
  {{.Name}}({{.Params.AsNamedArgsWithTypes}}) {{.Returns.AsReturnSignature}}
  {{- end}}
}

type {{.Name}}Shim struct {}

{{- range .Methods}}
func (p *{{$.Name}}Shim) {{.Name}}({{.Params.AsNamedArgsWithTypes}}) {{.Returns.AsReturnSignature}} {
  {{if .Returns.HasLength}}return {{end}}{{$.TargetAlias}}.{{.Name}}({{.Params.AsNamedArgsForInvocation}})
}
{{end}}
var _ {{.Name}} = new({{.Name}}Shim)
`
