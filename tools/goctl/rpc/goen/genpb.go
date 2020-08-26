package gogen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	astParser "github.com/tal-tech/go-zero/tools/goctl/rpc/parser"

	"github.com/dsymonds/gotoc/parser"

	"github.com/tal-tech/go-zero/core/lang"
	"github.com/tal-tech/go-zero/tools/goctl/rpc/execx"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
)

func (g *defaultRpcGenerator) genPb() error {
	importPath, filename := filepath.Split(g.Ctx.ProtoFileSrc)
	tree, err := parser.ParseFiles([]string{filename}, []string{importPath})
	if err != nil {
		return err
	}
	if len(tree.Files) == 0 {
		return errors.New("proto ast parse failed")
	}
	file := tree.Files[0]
	if len(file.Package) == 0 {
		return errors.New("expected package, but nothing found")
	}
	targetStruct := make(map[string]lang.PlaceholderType)
	for _, item := range file.Messages {
		if len(item.Messages) > 0 {
			return fmt.Errorf(`line %v: unexpected inner message near: "%v""`, item.Messages[0].Position.Line, item.Messages[0].Name)
		}
		name := stringx.From(item.Name)
		if _, ok := targetStruct[name.Lower()]; ok {
			return fmt.Errorf("line %v: duplicate %v", item.Position.Line, name)
		}
		targetStruct[name.Lower()] = lang.Placeholder
	}

	pbPath := g.dirM[dirPb]

	protoFileName := filepath.Base(g.Ctx.ProtoFileSrc)
	err = protocGenGo(g.Ctx.GoPath, g.Ctx.ProtoFileSrc, pbPath)
	if err != nil {
		return err
	}
	pbGo := strings.TrimSuffix(protoFileName, ".proto") + ".pb.go"
	pbFile := filepath.Join(pbPath, pbGo)
	bts, err := ioutil.ReadFile(pbFile)
	if err != nil {
		return err
	}
	aspParser := astParser.NewAstParser(bts, targetStruct, g.Ctx.Console)
	ast, err := aspParser.Parse()
	if err != nil {
		return err
	}
	if len(ast.Service) == 0 {
		return fmt.Errorf("service not found")
	}
	g.ast = ast
	return nil
}

func protocGenGo(gopath, proto, target string) error {
	src := filepath.Dir(proto)
	path := filepath.Join(gopath, "bin", "protoc-gen-go")
	fmt.Println(path)
	sh := fmt.Sprintf(`export PATH=$PATH:%s
protoc -I=%s --go_out=plugins=grpc:%s %s`, path, src, target, proto)
	return execx.RunShOrBat(sh)
}
