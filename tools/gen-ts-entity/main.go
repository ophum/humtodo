package main

import (
	"flag"
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

var (
	dest     string
	loadPath string
)

func init() {
	flag.StringVar(&dest, "dest", "./gen", "generate destination")
	flag.StringVar(&loadPath, "load-path", "./...", "load path")
	flag.Parse()
}

func main() {
	fset := token.NewFileSet()
	config := &packages.Config{
		Fset: fset,
		Mode: packages.LoadFiles | packages.LoadSyntax,
	}

	pkgs, err := packages.Load(config, loadPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, pkg := range pkgs {
		templateImports := []TemplateImport{}
		templateEntities := []TemplateEntity{}
		for _, f := range pkg.Syntax {
			comments := ast.NewCommentMap(fset, f, f.Comments)
			for n, c := range comments {
				isGen := false
				for _, cc := range c {
					if strings.Contains(cc.Text(), "+gen-ts-entity") {
						isGen = true
						break
					}
				}
				if !isGen {
					continue
				}

				switch n.(type) {
				case *ast.GenDecl:
					specs := n.(*ast.GenDecl).Specs
					if len(specs) > 0 {
						spec := specs[0]
						switch spec.(type) {
						case *ast.TypeSpec:
							typeSpec := spec.(*ast.TypeSpec)
							switch typeSpec.Type.(type) {
							case *ast.StructType:
								structType := typeSpec.Type.(*ast.StructType)
								templateEntity := TemplateEntity{
									Name:   typeSpec.Name.Name,
									Fields: []TemplateEntityField{},
								}
								for _, field := range structType.Fields.List {
									ident := &ast.Ident{}
									fieldType := ""
									switch field.Type.(type) {
									case *ast.SelectorExpr:
										ident = field.Type.(*ast.SelectorExpr).Sel
										fieldType = ident.Name
									case *ast.Ident:
										ident = field.Type.(*ast.Ident)
										fieldType = ident.Name
									case *ast.ArrayType:
										ident = field.Type.(*ast.ArrayType).Elt.(*ast.Ident)
										fieldType = ident.Name + "[]"
									}

									for _, name := range field.Names {
										// TODO: このあたり切り出したい
										n := name.Name
										t := transTypeToTS(fieldType)
										if f, ok := name.Obj.Decl.(*ast.Field); ok && f.Tag != nil {
											ref := reflect.StructTag(f.Tag.Value[1 : len(f.Tag.Value)-1])

											// json tag
											// 1つめをフィールド名,2つめにomitemptyがある場合はoptionalにする
											if d, ok := ref.Lookup("json"); ok {
												tmp := strings.Split(d, ",")
												n = tmp[0]

												if len(tmp) > 1 {
													for _, option := range tmp {
														switch option {
														case "omitempty":
															n += "?"
														}
													}
												}
											}

											// ts-type tag
											// 指定した型にする
											if d, ok := ref.Lookup("ts-type"); ok {
												t = d
											}

											// ts-import tag
											// 指定されたパスでimportさせる
											if d, ok := ref.Lookup("ts-import"); ok {
												isNew := true
												existsIndex := -1
												for i, v := range templateImports {
													if v.Path == d {
														isNew = false
														existsIndex = i
														break
													}
												}

												// ダサい気がする
												if isNew {
													templateImports = append(templateImports, TemplateImport{
														Names: []string{t},
														Path:  d,
													})
												} else {
													isAppend := true
													for _, name := range templateImports[existsIndex].Names {
														if name == t {
															isAppend = false
															break
														}
													}

													if isAppend {
														templateImports[existsIndex].Names = append(templateImports[existsIndex].Names, t)
													}
												}
											}
										}
										templateEntity.Fields = append(templateEntity.Fields, TemplateEntityField{
											Name: n,
											Type: t,
										})
									}

								}
								templateEntities = append(templateEntities, templateEntity)
							}
						}
					}
				}
			}
		}

		if len(templateEntities) == 0 {
			continue
		}

		t, err := template.New("").Parse(tsTemplate)
		if err != nil {
			log.Fatal(err.Error())
		}
		dirPath := filepath.Join(dest, pkg.Name)
		os.MkdirAll(dirPath, 0755)
		d, err := os.Create(filepath.Join(dirPath, pkg.Name+".ts"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer d.Close()

		args := TemplateArgs{
			Imports:  templateImports,
			Entities: templateEntities,
		}
		if err := t.Execute(d, args); err != nil {
			log.Fatal(err.Error())
		}

	}
}

func transTypeToTS(t string) string {
	transMap := map[string]string{
		"int":     "number",
		"int32":   "number",
		"int64":   "number",
		"float32": "number",
		"float64": "number",
	}

	tt, is := transMap[t]
	if is {
		return tt
	}
	return t
}

type TemplateEntityField struct {
	Name string
	Type string
}

type TemplateEntity struct {
	Name   string
	Fields []TemplateEntityField
}

type TemplateImport struct {
	Names []string
	Path  string
}

type TemplateArgs struct {
	Entities []TemplateEntity
	Imports  []TemplateImport
}

var tsTemplate = `
{{- range $v := .Imports}}
import {
	{{- range $n := $v.Names }}
	{{ $n }},
	{{- end}}
} from '{{ $v.Path }}';
{{- end }}

{{- range $v := .Entities }}
export interface {{ $v.Name }} {
{{- range $vv := $v.Fields }}
	{{ $vv.Name }}: {{$vv.Type }};
{{- end }}
};
{{ end }}
`
