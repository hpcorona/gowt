package gowt

import "strings"
import "os"
import "fmt"
import "strconv"
import "reflect"
import "time"

func GwtParse(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	idxp := strings.Index(partype, "/")
	if idxp >= 0 {
		partype = partype[0:idxp]
	}
	
	isarray := false
	if partype[0] == '[' {
		isarray = true
		partype = partype[1:]
	}
	
	equiv := reg.Data[partype]
	if equiv == nil {
		return nil, 1, os.NewError(fmt.Sprintf("Invalid type: %s\n", partype))
	}
	
	if isarray == false {
		return equiv.Parser(reg, strtable, payload, partype, idxv)
	}
	
	idxvp := idxv + 2
	total := ToInt(payload[idxv + 1])
	
	arr := reflect.MakeSlice(reflect.TypeOf([]interface{} {}), 0, total)
	
	for i := 0; i < total; i++ {
		v, adv, err := equiv.Parser(reg, strtable, payload, partype, idxvp)
		if err != nil {
			return nil, idxvp - idxv, err
		}
		
		arr = reflect.Append(arr, reflect.ValueOf(v))
		
		idxvp += adv
	}
	
	return arr, idxvp - idxv, nil
}

func GwtInt(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := strconv.Atoi(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return v, 1, nil
}

func GwtLong(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := ToLong(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return v, 1, nil
}

func GwtFloat(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := strconv.Atof32(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return v, 1, nil
}

func GwtDouble(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := strconv.Atof64(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return v, 1, nil
}

func GwtChar(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := strconv.Atoui(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return uint8(v), 1, nil
}

func GwtByte(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	v, err := strconv.Atoi(payload[idxv])
	if err != nil {
		return nil, 1, err
	}
	
	return int8(v), 1, nil
}

func GwtBool(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	return ToBool(payload[idxv]), 1, nil
}

func GwtString(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	return strtable[ToInt(payload[idxv])], 1, nil
}

func GwtDate(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	vb64 := payload[idxv + 1]
	datab, err := ToLong(vb64)
	if err != nil {
		return nil, 2, err
	}
	
	//milli := datab % 1000 (maybe when GAE supports Nanoseconds)
	datab = datab / 1000
	t := time.SecondsToLocalTime(datab)
	
	return t, 2, nil
}

func GwtArray(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	idxvp := idxv + 1
	total := ToInt(payload[idxvp])
	
	idxvp += 1
	elems := make(map[int]interface{})
	array := make([]interface{}, total)
	for i := 0; i < total; i++ {
		v := ToInt(payload[idxvp])
		if v < 0 {
			array[i] = elems[v]
			idxvp += 1
		} else {
			a, adv, err := GwtParse(reg, strtable, payload, strtable[v], idxvp)
			if err != nil {
				return nil, 1, err
			}
			
			array[i] = a
			idxvp += adv
			elems[(i + 1) * -2] = a
		}
	}
	
	return array, idxvp - idxv, nil
}

func GwtObject(reg *Registry, strtable, payload []string, partype string, idxv int) (interface{}, int, os.Error) {
	equiv := reg.Data[partype]
	if equiv == nil {
		return nil, 1, os.NewError(fmt.Sprintf("Invalid type: %s", partype))
	}
	
	if equiv.GoType.Kind() != reflect.Struct {
		return nil, 1, os.NewError(fmt.Sprintf("The type: %s is not a structure, it's a %s", equiv.JavaType, equiv.GoType.Kind().String()))
	}
	
	idxvp := idxv + 1
	obj := reflect.New(equiv.GoType).Elem()
	total := obj.NumField()
	for i := 0; i < total; i++ {
		fvalue := obj.Field(i)
		
		ftype := fvalue.Type()
		equivf := reg.Parsers[ftype]
		parser := equivf.Parser
		if parser == nil {
			return nil, 1, os.NewError(fmt.Sprintf("Cannot find parser for type %s", ftype.Kind().String()))
		}
		
		if ToInt(payload[idxvp]) == 0 {
			return nil, 1, nil
		}
		
		pvalue, adv, err := parser(reg, strtable, payload, equivf.JavaType, idxvp)
		if err != nil {
			return nil, 1, err
		}
		
		idxvp += adv
		fvalue.Set(reflect.ValueOf(pvalue))
	}
	
	adv := idxvp - idxv
	return obj.Interface(), adv, nil
}
