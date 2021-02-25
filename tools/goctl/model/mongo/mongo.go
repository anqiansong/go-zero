package mongo

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/model/mongo/generate"
	"github.com/urfave/cli"
)

// Command provides the entry for goctl
func Command(ctx *cli.Context) error {
	tp := strings.TrimSpace(ctx.String("type"))
	c := ctx.Bool("cache")
	o := strings.TrimSpace(ctx.String("dir"))
	if len(tp) == 0 {
		return errors.New("missing type")
	}

	a, err := filepath.Abs(o)
	if err != nil {
		return err
	}

	return generate.Do(&generate.Context{
		Type:   tp,
		Cache:  c,
		Output: a,
	})
}
