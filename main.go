package main

import (
	"github.com/TheDylong/goPackages/myTestPackage/helpers"
)

func main() {
	println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	println(myVar.TypeName)

}
