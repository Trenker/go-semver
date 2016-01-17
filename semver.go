// Package semver exports functions and types for validating and comparing semver versions
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

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var versionParts *regexp.Regexp = regexp.MustCompile(
	`(?i)^v?(\d{1,5})(\.\d+)?(\.\d+)?(\.\d+)?` +
		`[._-]?(?:(stable|beta|b|rc|alpha|a|patch|pl|p)((?:[.-]*)\d+)?)?` +
		`([.-]?dev)?$`)

var constParts *regexp.Regexp = regexp.MustCompile(`(?P<or>[,\|]+)(?P<rev>\!)?(?P<con>[\^~<>=]+)(?P<ver>[v\.\-a-z0-9]+)`)

var space *regexp.Regexp = regexp.MustCompile(`\s+`)

func NewVersion(s string) (*Version, error) {
	if len(s) > 4 && s[0:4] == "dev-" {
		v := &Version{
			src:     s,
			rolling: true,
			branch:  space.ReplaceAllString(strings.ToLower(s[4:]), ``),
		}
		return v, nil
	}

	res := versionParts.FindAllStringSubmatch(space.ReplaceAllString(s, ``), -1)

	if len(res) == 1 {
		matches := res[0]

		if len(matches) > 0 && matches[1] != "" {
			v := &Version{
				src:   s,
				major: str2int(matches[1]),
				n:     1,
			}

			if matches[2] != "" {
				v.minor = str2int(matches[2][1:])
				v.n++
			}

			if matches[3] != "" {
				v.bugfix = str2int(matches[3][1:])
				v.n++
			}

			if matches[4] != "" {
				v.sub = str2int(matches[4][1:])
				v.n++
			}

			if matches[5] != "" {
				if _, ok := suffix[matches[5]]; !ok {
					return nil, errors.New("Unknown version suffix " + matches[5])
				}
				v.suffixTyp = suffix[matches[5]]
			}

			if matches[6] != "" {
				v.suffixInc = str2int(matches[6])
			}

			if matches[7] != "" {
				v.dev = true
			}

			return v, nil
		}
	}
	return nil, errors.New("Cannot parse version " + s)
}

func NewConstraint(s string) (Constraints, error) {
	pt := constParts.FindAllStringSubmatch(`,`+strings.ToLower(space.ReplaceAllString(s, ``)), -1)
	cs := make([]*Constraint, len(pt))

	for i, s := range pt {
		c, err := parseConst(s[1], s[2], s[3], s[4])

		if err != nil {
			return nil, err
		}

		cs[i] = c
	}

	return Constraints(cs), nil
}

func str2int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
