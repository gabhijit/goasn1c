

package main

import (
	"fmt"
	"Asn1/goasn1c/parser"

)

func main() {

	fmt.Println("Hello world!")

	fmt.Printf("%x\n", parser.ModuleFlagUnkInstr)
	fmt.Printf("%x\n", parser.ModuleFlagTagInstr)
	fmt.Printf("%x\n", parser.ModuleFlagXerInstr)
	fmt.Printf("%x\n", parser.ModuleFlagImplicitTags)
	fmt.Printf("%x\n", parser.ModuleFlagExplicitTags)
	fmt.Printf("%x\n", parser.ModuleFlagExplicitTags)
	fmt.Printf("%x\n", parser.ModuleFlagAutomaticTags)
	fmt.Printf("%x\n", parser.Asn1ExprTypeBoolean)
	fmt.Printf("%x\n", parser.Asn1ExprTypeMax)
	fmt.Printf("%x\n", parser.Asn1ExprTypeInvalid)

}


