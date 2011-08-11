package gowt

import "fmt"
import "http"
import "appengine"
import "io/ioutil"

func HandleRequest(reg *Registry, w http.ResponseWriter, r *http.Request) {
	_ = appengine.NewContext(r)

	data, _ := ioutil.ReadAll(r.Body)
	total := len(data)
	
	fmt.Printf("Read Body: %d %s\n", total, string(data))
	_, err := ParseRpcRequest(reg, data)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Fprintf(w, string(data))
}
