package main

import (
	"fmt"

	"github.com/panyaxbo/libs/maskx"
)

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func main() {
	m := maskx.Init([]string{"newField"})
	t, err := m.Json(j, "prd")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}
