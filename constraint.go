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
)

type comparator func(v, c *Version) bool

var comperators = map[string]comparator{
	"^":  sameOrNewerInSameMajor,
	"~":  sameWithNewerLast,
	"=":  same,
	">=": sameOrHigher,
	"=>": sameOrHigher,
	"<=": sameOrLower,
	"=<": sameOrLower,
	"<":  lower,
	">":  higher,
}

type Constraint struct {
	f comparator
	v *Version
	r bool
	o bool
}

type Constraints []*Constraint

func (cs Constraints) Match(v *Version) bool {
	res := true
	for _, c := range cs {
		r := c.f(c.v, v)

		if (c.r) {
			r = !r
		}

		if c.o {
			res = res || r
		} else {
			res = res && r
		}
	}
	return res
}

func (c Constraints) MatchString(version string) bool {
	v, err := NewVersion(version)

	if err != nil {
		return false
	}

	return c.Match(v)
}

func sameOrNewerInSameMajor(v, c *Version) bool {
	return v.major == c.major && sameOrHigher(v, c)
}

func sameWithNewerLast(v, c *Version) bool {
	if v.n > 1 && v.major == c.major {
		switch v.n {
		case 2:
			return v.minor <= c.minor
		case 3:
			return v.minor == c.minor && v.bugfix <= c.bugfix
		case 4:
			return v.minor == c.minor && v.bugfix == c.bugfix && v.sub <= c.sub
		}
	}

	return false
}

func same(v, c *Version) bool {
	return v.src == c.src
}

func sameOrHigher(v, c *Version) bool {
	return same(v, c) || higher(v, c)
}

func sameOrLower(v, c *Version) bool {
	return same(v, c) || lower(v, c)
}

func lower(v, c *Version) bool {
	if v.major > c.major {
		return true
	}

	if v.major == c.major {
		if v.minor > c.minor {
			return true
		}

		if v.minor == c.minor {
			if v.bugfix > c.bugfix {
				return true
			}

			if v.bugfix == c.bugfix && v.sub > c.sub {
				return true
			}
		}
	}

	return false
}

func higher(v, c *Version) bool {
	if v.major < c.major {
		return true
	}

	if v.major == c.major {
		if v.minor < c.minor {
			return true
		}

		if v.minor == c.minor {
			if v.bugfix < c.bugfix {
				return true
			}

			if v.bugfix == c.bugfix && v.sub < c.sub {
				return true
			}
		}
	}

	return false
}

func parseConst(connect, negate, typ, version string) (*Constraint, error) {
	v, err := NewVersion(version)

	if err != nil {
		return nil, err
	}

	switch typ {
	case "":
		typ = "="
		break
	case "=>":
		typ = ">="
		break
	case "=<":
		typ = "<="
		break
	}

	if _, ok := comperators[typ]; !ok {
		return nil, errors.New("Unknown comperator " + typ)
	}

	c := &Constraint{
		v: v,
		r: negate == "!",
		o: connect == "|",
		f: comperators[typ],
	}

	return c, nil
}
