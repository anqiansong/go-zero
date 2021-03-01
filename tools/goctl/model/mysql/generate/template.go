package generate

import (
	"fmt"

	"github.com/tal-tech/go-zero/tools/goctl/model/mysql/template"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/urfave/cli"
)

const (
	category            = "mysql"
	defaultTemplateFile = "mysql-default.tpl"
	errTemplateFile     = "error.tpl"
)

var templates = map[string]string{
	defaultTemplateFile: template.DefaultTpl,
	errTemplateFile:     template.ErrorTpl,
}

func Category() string {
	return category
}

func Clean() error {
	return util.Clean(category)
}

func Templates(_ *cli.Context) error {
	return util.InitTemplates(category, templates)
}

func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return util.CreateTemplate(category, name, content)
}

func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return util.InitTemplates(category, templates)
}
