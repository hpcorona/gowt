package gowt

import "strings"
import "os"
import "strconv"
import "fmt"

type RpcRequest struct {
	VerMajor		int
	VerMinor		int
	URL					string
	StrongName	string
	Service			string
	Method			string
	Params			[]interface{}
}

func ParseRpcRequest(reg *Registry, request []byte) (*RpcRequest, os.Error) {
	req := string(request)
	params := strings.Split(req,"|",-1)
	
	if len(params) < 4 {
		return nil, os.NewError("Invalid GWT-RPC request")
	}
	
	ver := params[0:2]
	nump, err := strconv.Atoi(params[2])
	if err != nil {
		return nil, os.NewError("Invalid GWT-RPC request, cannot parse table size")
	}
	
	if len(params) <= 2+nump {
		return nil, os.NewError("Incomplete string table in GWT-RPC request")
	}
	
	strtable := params[2:3+nump]
	for i := 0; i < len(strtable); i++ {
		strtable[i] = strings.Replace(strtable[i], "\\!", "|", -1)
	}
	
	payload := params[3+nump:]
	
	if ver[0] == "7" && ver[1] == "0" {
		return parse_7_0(reg, ver, strtable, payload)
	}
	
	return nil, os.NewError(fmt.Sprintf("Versión not supported: %s.%s", ver[0], ver[1]))
}
