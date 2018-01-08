// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package codegen

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	tmpl "text/template"

	"github.com/pkg/errors"
	"github.com/uber/zanzibar/codegen/template_bundle"
)

// AssetProvider provides access to template assets
type AssetProvider interface {
	// AssetNames returns a list of named assets available
	AssetNames() []string
	// Asset returns the bytes for a provided asset name
	Asset(string) ([]byte, error)
}

type defaultAssetCollection struct{}

func (*defaultAssetCollection) AssetNames() []string {
	return templates.AssetNames()
}

func (*defaultAssetCollection) Asset(assetName string) ([]byte, error) {
	return templates.Asset(assetName)
}

var defaultFuncMap = tmpl.FuncMap{
	"lower":         strings.ToLower,
	"title":         strings.Title,
	"fullTypeName":  fullTypeName,
	"camel":         camelCase,
	"split":         strings.Split,
	"dec":           decrement,
	"basePath":      filepath.Base,
	"pascal":        pascalCase,
	"jsonMarshal":   jsonMarshal,
	"isPointerType": isPointerType,
	"unref":         unref,
	"lintAcronym":   LintAcronym,
}

func fullTypeName(typeName, packageName string) string {
	if typeName == "" || strings.Contains(typeName, ".") {
		return typeName
	}
	return packageName + "." + typeName
}

func decrement(num int) int {
	return num - 1
}

func jsonMarshal(jsonObj map[string]interface{}) (string, error) {
	str, err := json.Marshal(jsonObj)
	return string(str), err
}

func isPointerType(t string) bool {
	return strings.HasPrefix(t, "*")
}

func unref(t string) string {
	if strings.HasPrefix(t, "*") {
		return strings.TrimLeft(t, "*")
	}
	return t
}

// Template generates code for edge gateway clients and edgegateway endpoints.
type Template struct {
	template *tmpl.Template
}

// NewDefaultTemplate creates a bundle of templates.
func NewDefaultTemplate() (*Template, error) {
	return NewTemplate(
		&defaultAssetCollection{},
		defaultFuncMap,
	)
}

// NewTemplate returns a template helper for the provided asset collection
func NewTemplate(
	assetProvider AssetProvider,
	functionMap tmpl.FuncMap,
) (*Template, error) {
	t := tmpl.New("main").Funcs(functionMap)
	for _, file := range assetProvider.AssetNames() {
		fileContent, err := assetProvider.Asset(file)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"Could not read bin data for template %s",
				file,
			)
		}
		if _, err := t.New(file).Parse(string(fileContent)); err != nil {
			return nil, errors.Wrapf(err, "Could not parse template %s", file)
		}
	}
	return &Template{
		template: t,
	}, nil
}

// ExecTemplate executes the named templated, returning the generated bytes
func (t *Template) ExecTemplate(
	tplName string,
	tplData interface{},
	pkgHelper *PackageHelper,
) (ret []byte, rErr error) {
	defer func() {
		if r := recover(); r != nil {
			rErr = errors.Errorf(
				"panic when generating %s template in %s: %+v",
				tplName, pkgHelper.packageRoot, r,
			)
		}
	}()
	tplBuffer := bytes.NewBuffer(nil)
	if _, err := io.WriteString(tplBuffer, "// Code generated by zanzibar \n"+
		"// @generated \n \n"); err != nil {
		rErr = errors.Wrapf(err, "failed to write to file: %s", err)
		return
	}
	if _, err := io.WriteString(tplBuffer, pkgHelper.copyrightHeader); err != nil {
		rErr = errors.Wrapf(err, "failed to write to file: %s", err)
		return
	}
	if _, err := io.WriteString(tplBuffer, "\n\n"); err != nil {
		rErr = errors.Wrapf(err, "failed to write to file: %s", err)
		return
	}
	if err := t.template.ExecuteTemplate(
		tplBuffer,
		tplName,
		tplData,
	); err != nil {
		rErr = errors.Wrapf(
			err,
			"Error generating template %s",
			tplName,
		)
	}
	ret = tplBuffer.Bytes()
	return
}

func openFileOrCreate(file string) (*os.File, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			return nil, errors.Wrapf(
				err, "could not make directory: %s", file,
			)
		}
	}
	return os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
}
