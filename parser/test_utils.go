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

package parser

import (
	"os"
	"strconv"
	"strings"
)

type byName []os.FileInfo

func (b byName) Len() int {
	return len(b)
}

func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byName) Less(i, j int) bool {

	f := strings.Split(b[i].Name(), "-")[0]

	l := strings.Split(b[j].Name(), "-")[0]

	x, err := strconv.ParseInt(f, 10, 0)
	if err != nil {
		return b[i].Name() < b[j].Name()
	}

	y, err := strconv.ParseInt(l, 10, 0)
	if err != nil {
		return b[i].Name() < b[j].Name()
	}

	return x < y
}

func getTestNum(name string) (int64, error) {

	n := strings.Split(name, "-")[0]

	return strconv.ParseInt(n, 10, 0)
}

type testResult map[int64]error
