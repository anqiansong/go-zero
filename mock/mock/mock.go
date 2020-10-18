package mock

import (
	"fmt"
	"github.com/tal-tech/go-zero/mock/render"
	"reflect"

	"github.com/tal-tech/go-zero/core/logx"
)

var (
	emptyValue = reflect.Value{}
	tagKey     = "mock"
)

type (
	Mocker struct {
		v      interface{}
		parser *Parser
	}
)

func With(v interface{}, opts ...InvokeFn) *Mocker {
	m := &Mocker{
		v:      v,
		parser: NewParser(NewInvokeMatcher()),
	}
	opts = append(opts, render.RandomBool, render.RangeInt, render.RandomFloat,
		render.ObjectIdHex, render.RandomInt, render.PersonName,
		render.ClassName, render.GradeName, render.GradeClassName,
		render.Url, render.Email, render.Phone, render.Mobile,
		render.Ipv4, render.Ipv6, render.Latitude, render.Longitude,
		render.Uuid, render.Now, render.Date, render.Unix, render.Year,
		render.Month, render.Day, render.Hour, render.Minute, render.Second,
		render.Age, render.Province, render.City, render.District, render.Address,
		render.Gender)
	m.Register(opts...)
	return m
}

func (m *Mocker) Register(fn ...InvokeFn) {
	m.parser.invoker.Add(fn...)
}

func (m *Mocker) Mock() error {
	v := m.v
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Ptr {
		return fmt.Errorf("expected kind: pointer,but found: %v", vType.Kind().String())
	}
	vType.Elem()
	if vType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected kind: struct,but found: %v", vType.Elem().Kind().String())
	}
	vValue := reflect.ValueOf(v)
	newValue, err := m.typeOf(vType, "")
	if err != nil {
		return err
	}
	vValue.Elem().Set(newValue)
	return nil
}

func (m *Mocker) typeOf(vType reflect.Type, expression string) (reflect.Value, error) {
	vKind := vType.Kind()
	switch vKind {
	case reflect.Ptr, reflect.Interface:
		t := vType.Elem()
		return m.typeOf(t, expression)
	case reflect.Struct:
		return m.accessField(vType)
	default:
		value := reflect.New(vType).Elem()
		err := m.accessValue(value, expression)
		if err != nil {
			return emptyValue, nil
		}
		return value, nil
	}
}

func (m *Mocker) accessField(t reflect.Type) (reflect.Value, error) {
	va := reflect.New(t)
	elem := va.Elem()
	numField := elem.NumField()
	for i := 0; i < numField; i++ {
		fv := elem.Field(i)
		ft := t.Field(i)
		if !fv.CanSet() {
			continue
		}
		tag := ft.Tag.Get(tagKey)
		err := m.accessValue(fv, tag)
		if err != nil {
			return emptyValue, err
		}
	}
	return va.Elem(), nil
}

func (m *Mocker) accessSlice(sliceValue reflect.Value, count int, itemExpression string) error {
	typeOfSlice := sliceValue.Type()
	itemType := typeOfSlice.Elem()
	newSliceValue := reflect.MakeSlice(typeOfSlice, 0, count)
	for i := 0; i < count; i++ {
		vt := reflect.New(itemType).Elem()
		err := m.accessValue(vt, itemExpression)
		if err != nil {
			return err
		}
		newSliceValue = reflect.Append(newSliceValue, vt)
	}
	sliceValue.Set(newSliceValue)
	return nil
}

func (m *Mocker) accessArray(arrayValue reflect.Value, itemExpression string) error {
	typeOfArray := arrayValue.Type()
	itemType := typeOfArray.Elem()
	vt := reflect.New(itemType).Elem()
	for i := 0; i < arrayValue.Len(); i++ {
		err := m.accessValue(vt, itemExpression)
		if err != nil {
			return err
		}
		indexValue := arrayValue.Index(i)
		indexValue.Set(vt)
	}
	return nil
}

func (m *Mocker) accessMap(mapValue reflect.Value, count int, keyExpression, valueExpression string) error {
	typeOfMap := mapValue.Type()
	newMap := reflect.MakeMap(typeOfMap)
	vt := typeOfMap.Elem()
	kt := typeOfMap.Key()
	kv := reflect.New(kt).Elem()
	vv := reflect.New(vt).Elem()
	for i := 0; i < count; i++ {
		err := m.accessValue(kv, keyExpression)
		if err != nil {
			return err
		}
		err = m.accessValue(vv, valueExpression)
		if err != nil {
			return err
		}
		newMap.SetMapIndex(kv, vv)
	}
	mapValue.Set(newMap)
	return nil
}

func (m *Mocker) accessStruct(structValue reflect.Value) {
	typeOfStruct := structValue.Type()
	v, err := m.accessField(typeOfStruct)
	if err != nil {
		logx.Error(err)
		return
	}
	structValue.Set(v)
}

func (m *Mocker) accessPtr(ptrValue reflect.Value, expression string) error {
	vType := ptrValue.Type()
	newPtrValue := reflect.New(vType.Elem())
	newValue, err := m.typeOf(vType, expression)
	if err != nil {
		return err
	}
	newPtrValue.Elem().Set(newValue)
	ptrValue.Set(newPtrValue)
	return nil
}
func (m *Mocker) accessInterface(interfaceValue reflect.Value) {
	vType := interfaceValue.Type()
	newInterfaceValue := reflect.New(vType)
	interfaceValue.Set(newInterfaceValue)
}

func (m *Mocker) accessValue(v reflect.Value, expression string) error {
	value := new(Value)
	value.Value = v
	switch v.Kind() {
	case reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64:
		value.Type = TypeInt
	case reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		value.Type = TypeUInt
	case reflect.Float32, reflect.Float64:
		value.Type = TypeFloat
	case reflect.Bool:
		value.Type = TypeBool
	case reflect.String:
		value.Type = TypeString
	case reflect.Slice:
		value.Type = TypeSlice
		value.SliceLooper = func(count uint64, itemExpression string) error {
			return m.accessSlice(v, int(count), itemExpression)
		}
	case reflect.Array:
		value.Type = TypeSlice
		value.SliceLooper = func(count uint64, itemExpression string) error {
			return m.accessArray(v, itemExpression)
		}
	case reflect.Map:
		value.Type = TypeMap
		value.MapLooper = func(count uint64, keyExpression, valueExpression string) error {
			return m.accessMap(v, int(count), keyExpression, valueExpression)
		}
	case reflect.Struct:
		m.accessStruct(v)
	case reflect.Ptr:
		err := m.accessPtr(v, expression)
		if err != nil {
			return err
		}
	case reflect.Interface:
		value.Type = TypeInterface
		m.accessInterface(v)
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind().String())
	}
	if value.Type != UnKnown {
		if len(expression) == 0 {
			return fmt.Errorf("mock expression not found")
		}
		m.parser.Parse(expression, value)
	}
	return nil
}
