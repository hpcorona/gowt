package gowt

import "reflect"
import "os"
import "time"

type Unmarshaller func(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error)

var a Unmarshaller = GwtInt

type Equivalence struct {
	JavaType			string
	GoType				reflect.Type
	Parser				Unmarshaller
}

type Service struct {
	Name					string
	Methods				map[string]reflect.Method
}

type Registry struct {
	Data			map[string]*Equivalence
	Parsers		map[reflect.Type]*Equivalence
	Services	map[string]*Service
}

func NewRegistry() (*Registry) {
	reg := &Registry {
		Data: make(map[string]*Equivalence),
		Parsers: make(map[reflect.Type]*Equivalence),
		Services: make(map[string]*Service),
	}
	
	reg.AddEquivalence("I", reflect.TypeOf(int(0)), GwtInt)
	reg.AddEquivalence("J", reflect.TypeOf(int64(0)), GwtLong)
	reg.AddEquivalence("F", reflect.TypeOf(float32(0)), GwtFloat)
	reg.AddEquivalence("D", reflect.TypeOf(float64(0)), GwtFloat)
	reg.AddEquivalence("Z", reflect.TypeOf(bool(false)), GwtBool)
	reg.AddEquivalence("B", reflect.TypeOf(int8(0)), GwtByte)
	reg.AddEquivalence("C", reflect.TypeOf(uint8(0)), GwtChar)
	reg.AddEquivalence("java.lang.String", reflect.TypeOf(""), GwtString)
	reg.AddEquivalence("java.util.Date", reflect.TypeOf(&time.Time{}), GwtDate)
	reg.AddEquivalence("java.util.ArrayList", reflect.TypeOf(make([]interface{},0)), GwtArray)
	
	return reg
}

func (reg *Registry) AddEquivalence(javaType string, goType reflect.Type, parser Unmarshaller) {
	eq := &Equivalence {
		JavaType: javaType,
		GoType: goType,
		Parser: parser,
	}
	
	reg.Data[javaType] = eq
	reg.Parsers[goType] = eq
}
