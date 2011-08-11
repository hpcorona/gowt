package gowt

import "os"
import "fmt"

func parse_7_0(reg *Registry, ver, strtable, payload []string) (*RpcRequest, os.Error) {
	req := &RpcRequest {
		VerMajor: ToInt(ver[0]),
		VerMinor: ToInt(ver[1]),
		URL: strtable[ToInt(payload[0])],
		StrongName: strtable[ToInt(payload[1])],
		Service: strtable[ToInt(payload[2])],
		Method: strtable[ToInt(payload[3])],
	}
	
	np := ToInt(payload[4])
	req.Params = make([]interface{}, np)
	
	fmt.Printf("NP: %d\n", np)
	fmt.Printf("Str: %v\n", strtable)
	fmt.Printf("Pay: %v\n", payload)
	idxv := 5 + np
	for i := 0; i < np; i++ {
		idxt := 5 + i
		fmt.Printf("Parameter %d: %d is %d\n", i, idxv, idxt)

		partype := strtable[ToInt(payload[idxt])]
		parval, adv, err := GwtParse(reg, strtable, payload, partype, idxv)
		if err != nil {
			return nil, err
		}
		idxv += adv
		fmt.Printf("Parameter %d: %v is %s\n", i, parval, partype)
		
		req.Params[i] = parval
	}
	
	fmt.Printf("BUILT: %v\n", req)
	return req, os.NewError("Not implemented yet, or you've made all ok :)")
}
