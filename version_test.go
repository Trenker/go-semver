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

type TestVersionStub struct {
	Test     string
	Expected *Version
}

func TestNewVersion(t *testing.T) {
	tests := []*TestVersionStub{
		&TestVersionStub{"v1.2.3", &Version{"v1.2.3", 1, 2, 3, 0, 3, 0, 0, false, false, ""}},
		&TestVersionStub{"v1.2.3.4", &Version{"v1.2.3.4", 1, 2, 3, 4, 4, 0, 0, false, false, ""}},
		&TestVersionStub{"v1.2.3.4-alpha1", &Version{"v1.2.3.4", 1, 2, 3, 4, 4, 1, 1, false, false, ""}},
		&TestVersionStub{"v1.2.3.4_RC2", &Version{"v1.2.3.4", 1, 2, 3, 4, 4, 4, 2, false, false, ""}},
	}

	for _, testVersion := range tests {
		actual, err := NewVersion(testVersion.Test)
		exp := testVersion.Expected

		if err != nil {
			t.Errorf("Cannot parse %s: %s", testVersion, err)
		}

		if exp.major != actual.major {
			t.Errorf("Expected major %v but got %v", exp.major, exp.major)
		}

		if exp.minor != actual.minor {
			t.Errorf("Expected minor %v but got %v", exp.minor, exp.minor)
		}

		if exp.bugfix != actual.bugfix {
			t.Errorf("Expected bugfix %v but got %v", exp.bugfix, exp.bugfix)
		}
		if exp.sub != actual.sub {
			t.Errorf("Expected sub %v but got %v", exp.sub, exp.sub)
		}

		if exp.suffixTyp != actual.suffixTyp {
			t.Errorf("Expected suffix type %v but got %v", exp.suffixTyp, exp.suffixTyp)
		}

		if exp.suffixInc != actual.suffixInc {
			t.Errorf("Expected suffix number %v but got %v", exp.suffixInc, exp.suffixInc)
		}

		if exp.n != actual.n {
			t.Errorf("Expected number count %v but got %v", exp.n, exp.n)
		}
	}
}
