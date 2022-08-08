package main

import (
	"fmt"
)

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func main() {
	m := Init([]string{"newField"})
	t, err := m.Json(j)
	if err != nil {
		panic(err)
	}
	fmt.Println(*t)
}
