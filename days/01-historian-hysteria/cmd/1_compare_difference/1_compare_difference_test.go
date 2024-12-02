package main

import (
	"testing"
)

func TestCompareLists(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		diff     uint64
		err      error
	}{
		{
			name:     "correct compare_lists output for example data",
			filePath: "../../test/testdata/example_input.txt",
			diff:     11,
			err:      nil,
		},
		{
			name:     "actual challenge solution",
			filePath: "../../assets/input_list.txt",
			diff:     1579939,
			err:      nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			diff, err := compare_lists(tt.filePath)
			if err != tt.err {
				t.Errorf("testing %s: expected error: %v, got: %v", tt.name, tt.err, err)
			}
			if diff != tt.diff {
				t.Errorf("testing %s: expected %d, got %d", tt.name, tt.diff, diff)
			}
		})
	}
}
