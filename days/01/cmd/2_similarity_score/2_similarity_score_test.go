package main

import (
	"testing"
)

func TestCompareLists(t *testing.T) {
	tests := []struct {
		name       string
		filePath   string
		similarity int
		err        error
	}{
		{
			name:       "correct compare_lists output for example data",
			filePath:   "../../test/testdata/example_input.txt",
			similarity: 31,
			err:        nil,
		},
		{
			name:       "actual challenge solution",
			filePath:   "../../assets/input_list.txt",
			similarity: 20351745,
			err:        nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			similarity, err := compare_lists(tt.filePath)
			if err != tt.err {
				t.Errorf("testing %s: expected error: %v, got: %v", tt.name, tt.err, err)
			}
			if similarity != tt.similarity {
				t.Errorf("testing %s: expected %d, got %d", tt.name, tt.similarity, similarity)
			}
		})
	}
}