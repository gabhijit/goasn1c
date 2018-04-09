// Copyright (c) 2018 Abhijit Gadgil <gabhijit@iitbombay.org>. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"Asn1/goasn1c/parser"
	"fmt"
	"flag"
	"os"
	"io/ioutil"

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

	flag.Parse()

	// read a file
	for _, f := range(flag.Args()) {
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("Unable to read file \"%s\"\n", f)
			os.Exit(-1)
		}
		s := string(d)
		p := parser.NewParser(f)
		p.Parse(f, s)
	}

	os.Exit(0)
}
