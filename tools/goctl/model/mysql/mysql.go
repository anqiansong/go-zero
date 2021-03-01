package mysql

import (
	"errors"
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/model/mysql/generate"
	"github.com/urfave/cli"
)

func Action(ctx *cli.Context) error {
	url := ctx.String("dsn")
	p := ctx.String("pattern")
	t := ctx.String("template")
	o := ctx.String("dir")

	if len(url) == 0 {
		return errors.New("missing url")
	}

	if len(p) == 0 {
		return errors.New("missing pattern")
	}

	abs, err := filepath.Abs(o)
	if err != nil {
		return err
	}

	return generate.Do(&generate.Context{
		DataSource: url,
		Pattern:    p,
		File:       t,
		Output:     abs,
	})
}
