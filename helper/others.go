package helper

import (
	"bytes"
	"encoding/json"
)

type User struct {
	UserId string				`json:"userId"`
	FirstName string			`json:"firstName"`
	LastName string				`json:"lastName"`
	Address string				`json:"address"`
}

type ResponseUser struct {
	Code int                    `json:"code"`
	Error string			    `json:"error"`
	Data User					`json:"data"`		
}

func InputTranscode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}