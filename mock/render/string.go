package render

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/satori/go.uuid"
	"github.com/tal-tech/go-zero/mock/resource"
	"math/rand"
	"strings"
	"time"
)

var (
	chineseNumber = []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
)

func ObjectIdHex(_ *MockContext) interface{} {
	return bson.NewObjectId().Hex()
}

func PersonName(ctx *MockContext) interface{} {
	lang := ctx.Lang()
	switch lang {
	case LangZh:
		return resource.RandPersonNameZh()
	case LangEn:
		return resource.RandPersonNameEn()
	default: // random
		if time.Now().Unix()%2 == 0 {
			return resource.RandPersonNameZh()
		}
		return resource.RandPersonNameEn()
	}
}

func ClassName(_ *MockContext) interface{} {
	return fmt.Sprintf("%v班", time.Now().UnixNano()%21)
}

func GradeName(ctx *MockContext) interface{} {
	number := time.Now().UnixNano() % 9
	if number == 0 {
		number = 1
	}
	var gradeName interface{}
	switch ctx.Lang() {
	case LangZh:
		gradeName = chineseNumber[number]
	default:
		gradeName = number
	}
	return fmt.Sprintf("%v年级", gradeName)
}

func GradeClassName(ctx *MockContext) interface{} {
	number := time.Now().UnixNano() % 9
	if number == 0 {
		number = 1
	}
	var gradeName interface{}
	switch ctx.Lang() {
	case LangZh:
		gradeName = chineseNumber[number]
	default:
		gradeName = number
	}
	return fmt.Sprintf("%v年级%v", gradeName, ClassName(ctx))
}

func Url(_ *MockContext) interface{} {
	return resource.RandomUrl()
}

func Email(_ *MockContext) interface{} {
	return resource.RandomEmail()
}

func Ipv4(_ *MockContext) interface{} {
	list := make([]string, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		vInt := r.Int31n(255)
		list = append(list, fmt.Sprintf("%v", vInt))
	}
	return strings.Join(list, ".")
}

func Ipv6(_ *MockContext) interface{} {
	list := make([]string, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 8; i++ {
		vInt := r.Int31n(65536)
		list = append(list, fmt.Sprintf("%X", vInt))
	}
	return strings.Join(list, ":")
}

// -90-90
func Latitude(_ *MockContext) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	join := make([]string, 0)
	list := r.Perm(8)
	for _, item := range list {
		join = append(join, fmt.Sprintf("%v", item))
	}
	f := strings.Join(join, "")
	if time.Now().Unix()%2 == 0 {
		return fmt.Sprintf("-%v.%v", r.Intn(90), f)
	}
	return fmt.Sprintf("%v.%v", r.Intn(90), f)
}

// -180-180
func Longitude(_ *MockContext) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	join := make([]string, 0)
	list := r.Perm(8)
	for _, item := range list {
		join = append(join, fmt.Sprintf("%v", item))
	}
	f := strings.Join(join, "")
	if time.Now().Unix()%2 == 0 {
		return fmt.Sprintf("-%v.%v", r.Intn(180), f)
	}
	return fmt.Sprintf("%v.%v", r.Intn(180), f)
}

func Uuid(_ *MockContext) interface{} {
	return uuid.NewV4().String()
}

func Now(_ *MockContext) interface{} {
	return time.Now()
}

func Date(_ *MockContext) interface{} {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%v-%v-%v", y, int(m), d)
}

func Unix(_ *MockContext) interface{} {
	return time.Now().Unix()
}

func Year(_ *MockContext) interface{} {
	return time.Now().Year()
}

func Month(_ *MockContext) interface{} {
	return int(time.Now().Month())
}

func Day(_ *MockContext) interface{} {
	return time.Now().Day()
}

func Hour(_ *MockContext) interface{} {
	return time.Now().Hour()
}

func Minute(_ *MockContext) interface{} {
	return time.Now().Minute()
}

func Second(_ *MockContext) interface{} {
	return time.Now().Second()
}

func Age(_ *MockContext) interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(100)
}

func Gender(ctx *MockContext) interface{} {
	lang := ctx.Lang()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(10)
	if randomInt%2 == 0 {
		if lang == LangZh {
			return "男"
		}
		return "Male"
	}
	if lang == LangZh {
		return "女"
	}
	return "Female"

}
