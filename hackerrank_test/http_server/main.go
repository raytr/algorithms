package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Lake struct {
	Name string
	Area int32
}

type Action struct {
	Type    string
	Payload string
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	decode := json.NewDecoder(req.Body)
	l := Lake{}
	decode.Decode(&l)

	bs, _ := json.Marshal(l)
	action := Action{
		Type: "post",
		Payload: string(bs),
	}

	fmt.Print(action)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "0" {
		fmt.Print("404")
	}
	fmt.Print(id)
}

func getHandler(w http.ResponseWriter, req *http.Request) {
}

func main() {