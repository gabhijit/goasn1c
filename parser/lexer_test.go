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

// Testing for the Lexer - test ASN.1 files are taken from asn1c/testing @
// https://github.com/vlm/asn1c/tree/master/tests

package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEverything(t *testing.T) {
	args := os.Args[1:] // ignore the executable name

	var testsDir string

	if len(args) > 0 {
		testsDir = args[0]
	} else {
		testsDir = "tests"
	}

	d, _ := filepath.Abs(testsDir)

	f, err := os.Open(d)
	if err != nil {
		t.Error("Unable to Open tests directory : ", d)
		return
	}

	tests, err := f.Readdir(0)

	if err != nil {
		t.Error("Unable to read tests directory.")
	}

	for _, test := range tests {
		if test.IsDir() {
			continue
		}

		name := test.Name()
		name = filepath.Join(testsDir, name)
		name, _ = filepath.Abs(name)

		d, err := ioutil.ReadFile(name)
		if err != nil {
			t.Error("Unable to read file: ", name)
		}

		fmt.Println("Running parser for : ", name)
		s := string(d)
		p := NewParser(name)
		err = p.RunLexer(name, s, false)
		if err != nil {
			t.Error("Error during tokenization of file :", name)
		}
	}
}
