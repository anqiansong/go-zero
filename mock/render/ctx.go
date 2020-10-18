package render

import "fmt"

const (
	LangEn     Lang = "en"
	LangZh     Lang = "zh"
	LangNumber Lang = "number"
	LangMix    Lang = "mix"
)

var (
	EmptyMockContext        = &MockContext{}
	DefaultLength    uint64 = 0
	DefaultLang             = LangZh
	DefaultRange            = [2]int64{}
)

type (
	Lang string

	Template struct {
		prev       string
		after      string
		expression string
	}

	MockContext struct {
		name   string
		lang   Lang
		length uint64
		r      [2]int64
		args   []interface{}
	}

	EmptyContext struct {
	}
	RandBoolContext struct {
	}
	BoolContext struct {
	}
	SliceContext struct {
		Length uint64
	}
	MapContext struct {
		Length uint64
	}
	SpecValueContext struct {
		Value string
	}
	SpecConstContext struct {
		PackageName string
		ConstName   string
	}
	NumberRangeContext struct {
		Name  string
		Start int64
		End   int64
	}
	NumberLengthContext struct {
		Name   string
		Length uint64
	}
	StringRangeContext struct {
		Name  string
		Start int64
		Lang  Lang
		End   int64
	}
	StringLengthContext struct {
		Name   string
		Lang   Lang
		Length uint64
	}
)

func NewDefaultTemplate(prev, after, expression string) *Template {
	return &Template{
		prev:       prev,
		after:      after,
		expression: expression,
	}
}

func (t *Template) Render(expressionExecutor func(invokeFuncName string) (interface{}, error)) (string, error) {
	expression, err := expressionExecutor(t.expression)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v%v%v", t.prev, expression, t.after), nil
}

func NewMockContext(name string, lang Lang, length uint64, r [2]int64, arg ...interface{}) *MockContext {
	return &MockContext{
		name:   name,
		lang:   lang,
		length: length,
		r:      r,
		args:   arg,
	}
}
func InitLangByString(lang string) (Lang, error) {
	if len(lang) == 0 {
		return "", nil
	}
	switch lang {
	case LangEn.String(), LangZh.String(), LangNumber.String(), LangMix.String():
		return Lang(lang), nil
	default:
		return "", fmt.Errorf("unsupport lang: %s", lang)
	}
}
func (l Lang) String() string {
	return string(l)
}
func (c *MockContext) Length() uint64 {
	return c.length
}

func (c *MockContext) Range() [2]int64 {
	return c.r
}

func (c *MockContext) Name() string {
	return c.name
}

func (c *MockContext) Lang() Lang {
	return c.lang
}

func (c *MockContext) Args() []interface{} {
	return c.args
}
