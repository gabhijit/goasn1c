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
	"flag"
	"fmt"
	"github.com/gabhijit/goasn1c/parser"
	"io/ioutil"
	"os"
)

func main() {

	flag.Parse()

	// read a file
	for _, f := range flag.Args() {
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("Unable to read file \"%s\"\n", f)
			os.Exit(-1)
		}
		s := string(d)
		p := parser.NewParser(f)
		err = p.Parse(f, s, true)
		if err != nil {
			os.Exit(-1)
		}
	}

	os.Exit(0)
}
