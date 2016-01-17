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
	}

	for _, s := range d {
		c, err := NewConstraint(s.Constraint)

		if err != nil {
			t.Errorf("Cannot parse constraint %s: %s, %s -> %s\n", s.Constraint, err, s.Version, s.Constraint)
		}

		actual := c.MatchString(s.Version)

		if actual != s.Expected {
			t.Errorf("Expected %v, but got %v when checking %s is in %s\n", s.Expected, actual, s.Version, s.Constraint)
		}
	}
}
