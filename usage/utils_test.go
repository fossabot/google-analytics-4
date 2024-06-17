/*
Copyright 2023 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package usage

import "testing"

func TestToHumanSize(t *testing.T) {
	tests := map[string]struct {
		stringSize   string
		expectedSize string
		positiveTest bool
	}{
		"One Hundred Twenty Three thousand Four Hundred Fifty Six Tebibytes": {
			"123456 TiB",
			"121 PiB",
			true,
		},
		"One Gibibyte": {
			"1 GiB",
			"1.0 GiB",
			true,
		},
		"One Megabyte": {
			"1 MB",
			"977 KiB",
			true,
		},
		"One hundred four point five gigabyte": {
			"104.5 GB",
			"97 GiB",
			true,
		},
	}

	for testKey, testSuite := range tests {
		gotValue, err := toHumanSize(testSuite.stringSize)
		if (gotValue != testSuite.expectedSize || err != nil) && testSuite.positiveTest {
			t.Fatalf("Tests failed for %s, expected=%s, got=%s", testKey, testSuite.expectedSize, gotValue)
		}
	}
}
