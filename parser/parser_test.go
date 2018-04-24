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

// Testing for the Parser - test ASN.1 files are taken from asn1c/testing @
// https://github.com/vlm/asn1c/tree/master/tests

// +build parser all

package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

var testResults = testResult{
	0: nil, 1: nil, 2: ParserError, 3: nil, 4: nil, 5: nil, 6: nil, 7: nil, 8: nil, 9: nil,
	10: nil, 11: nil, 12: nil, 13: nil, 14: nil, 15: nil, 16: nil, 17: nil, 18: nil, 19: nil,
	20: nil, 21: nil, 22: nil, 23: nil, 24: nil, 25: nil, 26: nil, 27: nil, 28: nil, 29: nil,
	30: nil, 31: nil, 32: nil, 33: nil, 34: nil, 35: nil, 36: nil, 37: nil, 38: nil, 39: nil,
	40: nil, 41: nil, 42: nil, 43: nil, 44: nil, 45: nil, 46: nil, 47: nil, 48: nil, 49: nil,
	50: nil, 51: nil, 52: nil, 53: nil, 54: nil, 55: nil, 56: nil, 57: nil, 58: nil, 59: nil,
	60: nil, 61: nil, 62: nil, 63: nil, 64: nil, 65: nil, 66: nil, 67: nil, 68: nil, 69: nil,
	70: nil, 71: nil, 72: nil, 73: nil, 74: nil, 75: nil, 76: nil, 77: nil, 78: nil, 79: nil,
	80: nil, 81: nil, 82: nil, 83: nil, 84: nil, 85: nil, 86: nil, 87: nil, 88: nil, 89: nil,
	90: nil, 91: nil, 92: nil, 93: nil, 94: nil, 95: nil, 96: nil, 97: nil, 98: nil, 99: nil,
	100: nil, 101: nil, 102: nil, 103: nil, 104: nil, 105: nil, 106: nil, 107: nil, 108: nil, 109: nil,
	110: nil, 111: nil, 112: nil, 113: nil, 114: nil, 115: nil, 116: nil, 117: nil, 118: nil, 119: nil,
	120: nil, 121: nil, 122: nil, 123: nil, 124: nil, 125: nil, 126: nil, 127: nil, 128: nil, 129: nil,
	130: nil, 131: nil, 132: nil, 133: nil,
}

func TestParser(t *testing.T) {
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

	sort.Sort(byName(tests))

	if err != nil {
		t.Error("Unable to read tests directory.")
	}

	for _, test := range tests {
		if test.IsDir() {
			continue
		}

		name := test.Name()
		num, err := getTestNum(name)
		if err != nil {
			fmt.Println("Error getting test number for:", name)
		}

		name = filepath.Join(testsDir, name)
		name, _ = filepath.Abs(name)

		d, err := ioutil.ReadFile(name)
		if err != nil {
			t.Error("Unable to read file: ", name)
		}

		//fmt.Println("Running parser for : ", name)
		s := string(d)
		p := NewParser(name)

		err = p.Parse(name, s, false)

		if err != testResults[num] {
			fmt.Println(err)
			t.Error("Error during parsing file :", name)
		}
	}
}
