package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"regexp"
	"unicode"
)

var eventNamePattern = regexp.MustCompile(`^[a-z].*Event$`)

func parse(fileName string, r io.ReadSeeker) (*EventEmitterParams, error) {
	file, err := parser.ParseFile(token.NewFileSet(), fileName, r, 0)
	if err != nil {
		return nil, err
	}

	r.Seek(0, 0)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	fileContent := string(data)

	params := &EventEmitterParams{
		Package: file.Name.Name,
	}
	for _, decl := range file.Decls {
		switch it := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range it.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				t, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				name := ts.Name.Name
				if !eventNamePattern.MatchString(name) {
					continue
				}
				e := parseStructType(fileContent, name, t)
				params.Events = append(params.Events, e)
			}
		}
	}
	for _, importSpec := range file.Imports {
		var name string
		if importSpec.Name != nil {
			name = importSpec.Name.Name
		}
		i := &Import{
			Name: name,
			Path: importSpec.Path.Value,
		}
		params.Imports = append(params.Imports, i)
	}
	return params, nil
}

func parseStructType(fileContent, name string, t *ast.StructType) *Event {
	var params []*EventParam
	for _, field := range t.Fields.List {
		var names []string
		for _, name := range field.Names {
			if !unicode.IsUpper(rune(name.Name[0])) {
				names = append(names, name.Name)
			}
		}
		if len(names) == 0 {
			continue
		}
		typeName := fileContent[field.Type.Pos()-1 : field.Type.End()-1]
		params = append(params, &EventParam{
			Names: names,
			Type:  typeName,
		})
	}
	return &Event{
		Name:   name,
		Params: params,
	}
}
