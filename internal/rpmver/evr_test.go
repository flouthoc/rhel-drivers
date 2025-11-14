package rpmver

import (
	"testing"
)

func TestCompareEVR(t *testing.T) {
	tests := []struct {
		name                       string
		epoch1, version1, release1 string
		epoch2, version2, release2 string
		expected                   int
	}{
		{"empty epochs", "", "1.0", "1", "", "1.0", "1", 0},
		{"epoch1 greater", "1", "1.0", "1", "", "1.0", "1", 1},
		{"epoch2 greater", "", "1.0", "1", "1", "1.0", "1", -1},
		{"empty epoch treated as zero", "", "1.0", "1", "0", "1.0", "1", 0},

		{"version1 greater", "", "2.0", "1", "", "1.0", "1", 1},
		{"version2 greater", "", "1.0", "1", "", "2.0", "1", -1},
		{"equivalent versions", "", "1.0", "1", "", "1.0", "1", 0},

		{"release1 greater", "", "1.0", "2", "", "1.0", "1", 1},
		{"release2 greater", "", "1.0", "1", "", "1.0", "2", -1},
		{"equivalent releases", "", "1.0", "1", "", "1.0", "1", 0},

		{"epoch and version comparison", "1", "2.0", "1", "0", "1.0", "1", 1},
		{"differing epochs with same version and release", "2", "1.0", "3", "1", "1.0", "3", 1},
		{"differing all fields", "0", "2.0", "5", "1", "1.0", "3", -1},
		{"complex comparison", "1", "1.2.3", "4", "1", "2.0", "0", -1},

		{"empty version1", "", "", "", "0", "1.0", "1", -1},
		{"empty version2", "0", "1.0", "1", "", "", "", 1},
		{"empty release1", "0", "1.0", "", "0", "1.0", "1", -1},
		{"empty release2", "0", "1.0", "1", "0", "1.0", "", 1},

		{"all fields empty", "", "", "", "", "", "", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CompareEVR(tc.epoch1, tc.version1, tc.release1, tc.epoch2, tc.version2, tc.release2)
			if result != tc.expected {
				t.Errorf("CompareEVR(%q, %q, %q, %q, %q, %q) = %d; want %d",
					tc.epoch1, tc.version1, tc.release1, tc.epoch2, tc.version2, tc.release2, result, tc.expected)
			}
		})
	}
}
