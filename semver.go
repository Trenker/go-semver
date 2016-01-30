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

// vParts is a regex to determine all parts of a version string
var vParts *regexp.Regexp = regexp.MustCompile(`(?i)` +
	`^v?(\-?\d{1,5})(\.\-?\d+)?(\.\-?\d+)?(\.\-?\d+)?` +
	`[._-]?(?:(stable|beta|b|rc|alpha|a|patch|pl|p)((?:[.-]*)\d+)?)?` +
	`([.-]?dev)?$`)

// branch is a regex to determine if a version number is a branch reference, like used in rolling / development versions
var branch = *regexp.MustCompile(`(?i)^\!?dev\-`)

// wildcard is a regex to check for the wildcard characters * and x
var wildcard *regexp.Regexp = regexp.MustCompile(`[x\*]+`)

// cParts is a regex to determine the parts of a constraint
// and split multiple constraints
var cParts *regexp.Regexp = regexp.MustCompile(`(?i)(?P<or>[,\|]+)(?P<rev>\!)?(?P<con>[\^~<>=]+)?(?P<ver>[\.\-0-9\*a-z]+)`)

// space is a regex to determine white-space characters
var space *regexp.Regexp = regexp.MustCompile(`\s+`)

// NewVersion constructs a new version type based on the given string
// The argument s must be semver version string, possibly containing a wildcard character, as used in constraints
// or "floating" versions
func NewVersion(s string) (*Version, error) {
	if len(s) > 4 && branch.MatchString(s) {
		b := branch.ReplaceAllString(s, "")
		b = space.ReplaceAllString(b, "")
		b = strings.ToLower(b)

		v := &Version{
			src:     s,
			rolling: true,
			branch:  b,
		}
		return v, nil
	}

	o := s
	s = wildcard.ReplaceAllString(s, "-1")
	res := vParts.FindAllStringSubmatch(space.ReplaceAllString(s, ``), -1)

	if len(res) == 1 {
		matches := res[0]

		if len(matches) > 0 && matches[1] != "" {
			return parseVersion(s, matches)
		}
	}
	return nil, errors.New("Malformed version " + o)
}

// NewConstraint returns a Constraints type which is a list of Constraint values
// The argument s is a string value containing one or more valid constraints
func NewConstraint(s string) (Constraints, error) {
	pt := cParts.FindAllStringSubmatch(`,`+strings.ToLower(space.ReplaceAllString(s, ``)), -1)

	if len(pt) < 1 {
		return nil, errors.New("Cannot find constraints in " + s)
	}

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

// str2int is a convinience wrapper around strconv.Atoi that will always return an int
// and silently discard possible errors
func str2int(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return i
}
