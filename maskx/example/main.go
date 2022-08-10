package main

import (
	"fmt"

	"github.com/panyaxbo/libs/maskx"
)

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func main() {
	m := maskx.Init([]string{"newField"})
	t1, err := m.JsonMaskEncrypted(j, "dev")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t1)

	t2, err := m.JsonMaskEncrypted(j, "dev")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t2)

	t3, err := m.JsonMaskEncrypted(j, "prd")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t3)
}
