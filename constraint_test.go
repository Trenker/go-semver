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

import "testing"

type TestConstraintStub struct {
	Version, Constraint string
	Expected            bool
}

func TestNewConstraint(t *testing.T) {
	d := []*TestConstraintStub{
		&TestConstraintStub{"1.2.3", "~1.2.0", true},
		&TestConstraintStub{"1.2.3", "^1.2.0", true},
		&TestConstraintStub{"1.2.3", "^1.1.0", true},
		&TestConstraintStub{"1.2.3", "~1.3.0", false},
		&TestConstraintStub{"1.2.3", "^1.3.0", false},
		&TestConstraintStub{"1.2.3", "^1.1.0", true},
		&TestConstraintStub{"1.2.3", ">1.1.0", true},
		&TestConstraintStub{"1.2.3", ">=1.1.0", true},
		&TestConstraintStub{"1.2.3", "<1.3.0", true},
		&TestConstraintStub{"1.2.3", "<=1.3.0", true},
		&TestConstraintStub{"1.2.3", "<1.2.0", false},
		&TestConstraintStub{"1.2.3", "<=1.2.0", false},
		&TestConstraintStub{"1.2.3", "<=1.2.0,>1.0", false},
		&TestConstraintStub{"1.2.3", "<=1.2.0|>1.0", true},
		&TestConstraintStub{"1.2.3", "^1.0.12", true},
		&TestConstraintStub{"1.2.3", "^1.0.12,>1.3", false},
		&TestConstraintStub{"1.2.3", "<2.0|>3.3", true},
		&TestConstraintStub{"1.2.3", "1.2.*", true},
		&TestConstraintStub{"1.2.3", "*", true},
		&TestConstraintStub{"1.2.3", "1.*", true},
		&TestConstraintStub{"1.2.3", "1.3.*", false},
		&TestConstraintStub{"1.2.3", "1.2.x", true},
		&TestConstraintStub{"1.2.3", "1.x", true},
		&TestConstraintStub{"1.2.3", "1.3.x", false},
		&TestConstraintStub{"dev-master", "dev-master", true},
		&TestConstraintStub{"dev-master", "!dev-master", false},
		&TestConstraintStub{"dev-master", "!dev-develop", true},
		&TestConstraintStub{"dev-master", "dev-develop", false},
		&TestConstraintStub{"dev-master", "dev-develop|dev-master", true},
	}

	for _, s := range d {
		c, err := NewConstraint(s.Constraint)

		if err != nil {
			t.Errorf("Cannot parse constraint %s: %s, %s -> %s\n", s.Constraint, err, s.Version, s.Constraint)
		}

		v, err := NewVersion(s.Version)

		if err != nil {
			t.Errorf("Cannot parse version %s: %s\n", s.Version, err)
		}

		actual := c.Match(v)

		if actual != s.Expected {
			t.Errorf("Expected %v, but got %v when checking if %s is in %s\n", s.Expected, actual, s.Version, s.Constraint)
		}
	}
}
