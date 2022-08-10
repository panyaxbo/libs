package maskx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var j = []byte(`{"foo":1,"bar":2,"baz":[3,4],"phoneNo":123456789, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`)

func TestMaskEnncryptedDev(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.JsonMaskEncrypted(j, "dev")
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}

func TestMaskEnncryptedPrd(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.JsonMaskEncrypted(j, "prd")
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}
func TestMaskSymbolStar(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.JsonMaskSymbol(j, SymbolStar)
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}
func TestMaskSymbolDash(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.JsonMaskSymbol(j, SymbolDash)
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}
func TestMaskSymbolSharp(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.JsonMaskSymbol(j, SymbolSharp)
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}
func TestMask(t *testing.T) {
	m := Init([]string{"newField"})
	t1, err := m.Json(j)
	if err != nil {
		panic(err)
	}
	s1 := "{\"foo\":1,\"bar\":2,\"baz\":[3,4],\"phoneNo\":NODarD+pRu1XaV/o+GfnrqMRyytZ91mknA==, \"newField\":\"cbea7ImKlk/dWjB22kgR+sKD71I=\", \"userInfo\":{\"firstname\":\"TqCA7Gn3EKFdYu+NX5TZyMElibMZxqt0\", \"lastname\": \"V72D+WT+Ab0col3Ask/y/TwTigBCtI2D5Hpo\"}}"
	fmt.Println(*t1)

	assert.NotEmpty(t, t1, s1)
}
