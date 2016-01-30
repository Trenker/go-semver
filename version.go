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

import "errors"

var suffix map[string]int = map[string]int{
	"a":      1,
	"alpha":  1,
	"b":      2,
	"beta":   2,
	"patch":  3,
	"pl":     3,
	"p":      3,
	"rc":     4,
	"":       5,
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

func parseVersion(s string, p []string) (*Version, error) {
	v := &Version{
		src:   s,
		major: str2int(p[1]),
		n:     1,
	}

	if p[2] != "" {
		v.minor = str2int(p[2][1:])
		v.n++
	}

	if p[3] != "" {
		v.bugfix = str2int(p[3][1:])
		v.n++
	}

	if p[4] != "" {
		v.sub = str2int(p[4][1:])
		v.n++
	}

	if p[5] != "" {
		if _, ok := suffix[p[5]]; !ok {
			return nil, errors.New("Unknown version suffix " + p[5])
		}
		v.suffixTyp = suffix[p[5]]
	}

	if p[6] != "" {
		v.suffixInc = str2int(p[6])
	}

	if p[7] != "" {
		v.dev = true
	}

	return v, nil
}
