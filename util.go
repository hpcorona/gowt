package gowt

import "strconv"
import "os"
import "fmt"
import "strings"

func ToInt(v string) int {
	i, _ := strconv.Atoi(v)
	
	return i
}

func ToFloat(v string) float32 {
	f, _ := strconv.Atof32(v)
	
	return f
}

func ToBool(v string) bool {
	b := (v == "1" || v == "true" || v == "yes")
	
	return b
}

func ToChar(v string) uint8 {
	c, _ := strconv.Atoui(v)
	
	return uint8(c)
}

func ToByte(v string) int8 {
	b, _ := strconv.Atoi(v)
	
	return int8(b)
}

const baseGwtLong = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789$_"

func ToLong(v string) (int64, os.Error) {
	var t int64 = 0
	
	for i := 0; i < len(v); i++ {
		c := v[i:i+1]
		idx := strings.Index(baseGwtLong, c)
		if idx < 0 {
			return 0, os.NewError(fmt.Sprintf("Not long GWT, found: %s", v[i]))
		}
		
		t = t * 64
		t += int64(idx)
	}
	
	return t, nil
}
