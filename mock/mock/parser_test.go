package mock

import (
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/mock/render"
	"strconv"
	"testing"
)

type (
	Person struct {
		Name   string `mock:"personName(zh)"`
		Gender string `mock:"gender(zh)"`
	}
	AliasInt int
	AliasMap map[string][]Person

	MockData struct {
		RandomBool      bool             `mock:"randomBool"`
		SpecBool        bool             `mock:"spec={true}"`
		BoolTrue        bool             `mock:"bool=true"`
		SpecString      string           `mock:"spec={I am spec string}"`
		SpecInt         int64            `mock:"spec={123}"`
		AliasInt        AliasInt         `mock:"randomInt(1)"`
		SpecFloat       float64          `mock:"spec={123.2}"`
		SliceInt        []int64          `mock:"slice(2,rangeInt[10,20])"`
		RangeInt        int64            `mock:"rangeInt[90,110]"`
		RandomInt       int64            `mock:"randomInt(8)"`
		Map             map[string]*int  `mock:"map(3,objectIdHex,rangeInt[10,20])"`
		ObjectIdHex     string           `mock:"objectIdHex"`
		SlicePersonName []string         `mock:"slice(10,personName(zh))"`
		ClassName       string           `mock:"className"`
		GradeName       string           `mock:"gradeName(zh)"`
		GradeClassName  string           `mock:"gradeClassName(zh)"`
		Url             string           `mock:"url"`
		Image           string           `mock:"<<url>>{/xx.png}"`
		Sum             int64            `mock:"sum({1},{2},{3})"`
		Email           string           `mock:"email"`
		Phone           string           `mock:"phone"`
		Mobile          string           `mock:"mobile"`
		Ipv4            string           `mock:"ipv4"`
		Ipv6            string           `mock:"ipv6"`
		Latitude        string           `mock:"latitude"`
		Longitude       string           `mock:"longitude"`
		UUID            string           `mock:"uuid"`
		Now             string           `mock:"now"`
		Date            string           `mock:"date"`
		Unix            string           `mock:"unix"`
		Year            string           `mock:"year"`
		Month           string           `mock:"month"`
		Day             string           `mock:"day"`
		Hour            string           `mock:"hour"`
		Minute          string           `mock:"minute"`
		Second          string           `mock:"second"`
		Age             string           `mock:"age"`
		Province        string           `mock:"province"`
		City            string           `mock:"city"`
		District        string           `mock:"district"`
		Address         string           `mock:"address"`
		Provinces       []string         `mock:"slice(3,province)"`
		Cities          []string         `mock:"slice(3,city)"`
		Addresses       []string         `mock:"slice(5,address)"`
		StringPointer   []*string        `mock:"slice(5,province)"`
		PointInt        *int             `mock:"randomInt(12)"`
		PointFloat      *float64         `mock:"randomFloat(5)"`
		MapPointer      map[int]*float64 `mock:"map(3,randomInt(3),randomFloat(5))"`
		PointerStruct   *Person
		NormalStruct    Person
		AliasMap        AliasMap     `mock:"map(3,province,slice(5,-))"`
		ArrayPointer    [2][]*Person `mock:"slice(slice(3,-))"`
		ArrayStruct     [1][]Person  `mock:"slice(slice(2,-))"`
	}
)

func TestParser_Parse(t *testing.T) {
	data := &MockData{}
	mocker := With(data)
	mocker.Register(sum)
	logx.Must(mocker.Mock())
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func sum(ctx *render.MockContext) interface{} {
	args := ctx.Args()
	var sum int64 = 0
	for _, arg := range args {
		vInt, _ := strconv.ParseInt(fmt.Sprintf("%v", arg), 10, 64)
		sum += vInt
	}
	return sum
}
