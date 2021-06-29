package main

import (
	"scope/packageone"
)

var myVar = "This is myVar"

func main() {

	var blockVar = "This is blockVar"

	packageone.PrintMe(myVar)
	packageone.PrintMe(blockVar)
	packageone.PrintMe(packageone.PackageVar)
}
