package generate

import (
	"github.com/tal-tech/go-zero/tools/goctl/model/mongo/template"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

// Context defines the model generating what they needs
type Context struct {
	Type   string
	Time   bool
	Cache  bool
	Output string
}

// Do executes model template and output the result into the specified file path
func Do(ctx *Context) error {
	text, err := util.LoadTemplate(category, modelTemplateFile, template.Text)
	if err != nil {
		return err
	}

	return util.With("model").Parse(text).GoFmt(true).SaveTo(ctx, ctx.Output, false)
}
