package generate

import (
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/model/mongo/template"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

// Context defines the model generating what they needs
type Context struct {
	Type   string
	Cache  bool
	Output string
}

// Do executes model template and output the result into the specified file path
func Do(ctx *Context) error {
	err := generateModel(ctx)
	if err != nil {
		return err
	}

	return generateError(ctx)
}

func generateModel(ctx *Context) error {
	text, err := util.LoadTemplate(category, modelTemplateFile, template.Text)
	if err != nil {
		return err
	}

	output := filepath.Join(ctx.Output, strings.ToLower(ctx.Type+"model.go"))

	return util.With("model").Parse(text).GoFmt(true).SaveTo(ctx, output, false)
}

func generateError(ctx *Context) error {
	text, err := util.LoadTemplate(category, errTemplateFile, template.Error)
	if err != nil {
		return err
	}

	output := filepath.Join(ctx.Output, "error.go")

	return util.With("error").Parse(text).GoFmt(true).SaveTo(ctx, output, false)
}
