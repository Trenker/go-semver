package semver

//
// Copyright 2016 Georg Großberger
//
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
//

var suffix map[string]int = map[string]int{
	"a":      1,
	"alpha":  1,
	"b":      2,
	"beta":   2,
	"patch":  3,
	"pl":     3,
	"p":      3,
	"rc":     4,
	"stable": 5,
}

type Version struct {
	src       string
	major     int
	minor     int
	bugfix    int
	sub       int
	n         int
	suffixTyp int
	suffixInc int
	dev       bool
	rolling   bool
	branch    string
}

// Function String implements the Stringer interface
func (v *Version) String() string {
	return v.src
}
