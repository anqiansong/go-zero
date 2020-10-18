package render

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RangeInt(ctx *MockContext) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := ctx.Range()[0]
	max := ctx.Range()[1]
	return r.Int63n(max-min) + min
}

func RandomInt(ctx *MockContext) interface{} {
	length := ctx.Length()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := r.Perm(int(length))
	builder := new(strings.Builder)
	for index, item := range list {
		if index == 0 && item == 0 {
			ctx.r = [2]int64{1, 9}
			v := RangeInt(ctx)
			builder.WriteString(fmt.Sprintf("%v", v))
			continue
		}
		builder.WriteString(fmt.Sprintf("%v", item))
	}
	vInt, _ := strconv.ParseInt(builder.String(), 10, 64)
	return vInt
}

func RandomFloat(ctx *MockContext) interface{} {
	length := ctx.Length()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := r.Perm(int(length))
	builder := new(strings.Builder)
	dotIndex := r.Intn(int(length) - 1)
	if dotIndex == 0 {
		dotIndex = 1
	}
	for index, item := range list {
		if index == dotIndex {
			builder.WriteString(fmt.Sprintf(".%v", item))
			continue
		}
		if index == 0 && item == 0 {
			ctx.r = [2]int64{1, 9}
			v := RangeInt(ctx)
			builder.WriteString(fmt.Sprintf("%v", v))
			continue
		}
		builder.WriteString(fmt.Sprintf("%v", item))
	}
	n := json.Number(builder.String())
	vFloat, _ := n.Float64()
	return vFloat
}
