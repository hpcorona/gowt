package gowt

import "strconv"

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
