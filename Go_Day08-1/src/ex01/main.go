package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(st interface{}) {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)
	for i := 0; i < val.NumField(); i++ {
		tag := string(typ.Field(i).Tag)
		var tagString string
		if tag != "" {
			for _, key := range strings.Split(tag, " ") {
				tagParts := strings.Split(key, ":")
				tagString += fmt.Sprintf("%s=%s", tagParts[0], strings.Trim(tagParts[1], "\""))
			}
		}
		if tagString != "" {
			fmt.Printf("%s(%s):%v\n", typ.Field(i).Name, tagString, val.Field(i))
		} else {
			fmt.Printf("%s:%v\n", typ.Field(i).Name, val.Field(i))
		}
	}
}

func main() {
	flower1 := UnknownPlant{FlowerType: "aboba", LeafType: "square", Color: 255}
	flower2 := AnotherUnknownPlant{FlowerColor: 10, LeafType: "lanceolate", Height: 15}
	describePlant(flower1)
	describePlant(flower2)
}
