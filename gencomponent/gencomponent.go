package main

import (
	"bytes"
	"go/ast"
	"go/build"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Component struct {
	Name  string
	Props []*types.Var
	State []*types.Var
}

func main() {
	wd, _ := os.Getwd()

	buildPkg, err := build.Import(os.Args[1], wd, 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	files := make([]*ast.File, len(buildPkg.GoFiles))
	for i, name := range buildPkg.GoFiles {
		file, err := parser.ParseFile(fset, filepath.Join(buildPkg.Dir, name), nil, 0)
		if err != nil {
			panic(err)
		}
		files[i] = file
	}

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	conf := types.Config{
		Importer: importer.Default(),
	}
	pkg, err := conf.Check(buildPkg.ImportPath, fset, files, &info)
	if err != nil {
		panic(err)
	}

	var components []*Component
	imports := map[string]struct{}{
		"github.com/neelance/dom":               struct{}{},
		"github.com/neelance/dom/componentutil": struct{}{},
		buildPkg.ImportPath:                     struct{}{},
	}
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		if s, ok := obj.Type().Underlying().(*types.Struct); ok {
			comp := &Component{
				Name: name,
			}
			for i := 0; i < s.NumFields(); i++ {
				f := s.Field(i)
				if f.Type().String() == "github.com/neelance/dom.Instance" {
					continue
				}

				switch f.Exported() {
				case true:
					comp.Props = append(comp.Props, f)
				case false:
					comp.State = append(comp.State, f)
				}

				types.TypeString(f.Type(), func(pkg *types.Package) string {
					imports[pkg.Path()] = struct{}{}
					return ""
				})
			}
			components = append(components, comp)
		}
	}

	data := struct {
		Imports    map[string]struct{}
		PkgName    string
		Components []*Component
	}{
		Imports:    imports,
		PkgName:    buildPkg.Name,
		Components: components,
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("impl.gen.go", formatted, 0666); err != nil {
		panic(err)
	}
}

var tmpl = template.Must(template.New("").Delims("<<<", ">>>").Funcs(template.FuncMap{
	"c":  capitalize,
	"uc": uncapitalize,
	"type": func(typ types.Type) string {
		return types.TypeString(typ, func(pkg *types.Package) string {
			return pkg.Name()
		})
	},
}).Parse(`<<<$t := .>>>// GENERATED, DO NOT CHANGE

package impl

import (
	<<<range $k, $v := .Imports>>>
		"<<<$k>>>"
	<<<end>>>
)

func init() {<<<range .Components>>>
	<<<$t.PkgName>>>.Reconcile<<<.Name>>> = reconcile<<<.Name>>><<<end>>>
}

<<<range .Components>>>
  <<<$c := .>>>

  type <<<.Name>>>Impl struct {
		componentutil.Core
		componentutil.EmptyLifecycle
    <<<range .Props>>>
      <<<.Name>>> <<<.Type | type>>><<<end>>>
    <<<range .State>>>
      <<<.Name>>> <<<.Type | type>>><<<end>>>
  }

	func (c *<<<.Name>>>Impl) applyProps(spec *spec.<<<.Name>>>) {<<<range .Props>>>
		  c.<<<.Name | c>>> = spec.<<<.Name | c>>><<<end>>>
		c.DoRender()
	}

  func reconcile<<<.Name>>>(newSpec *<<<$t.PkgName>>>.<<<.Name>>>, oldSpec dom.Spec) {
  	if oldSpec, ok := oldSpec.(*<<<$t.PkgName>>>.<<<.Name>>>); ok {
  		newSpec.Instance = oldSpec.Instance
  		newSpec.Instance.(*<<<.Name>>>Impl).applyProps(newSpec)
  		return
  	}

  	inst := &<<<.Name>>>Impl{}
  	inst.Lifecycle = inst
  	newSpec.Instance = inst
		inst.ComponentWillMount()
		inst.applyProps(newSpec)
		inst.ComponentDidMount()
  }
<<<end>>>
`))

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func uncapitalize(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
