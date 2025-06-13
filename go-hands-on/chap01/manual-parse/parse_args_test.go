package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	config
	args []string
	err  error
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args: []string{"-h"},
			err:  nil,
			config: config{
				printUsage: true,
				numTimes:   0,
			},
		},
		{
			args: []string{"10"},
			err:  nil,
			config: config{
				printUsage: false,
				numTimes:   10,
			},
		},
		{
			args: []string{"abc"},
			err:  errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{
				printUsage: false,
				numTimes:   0,
			},
		},
		{
			args: []string{"1", "foo"},
			err:  errors.New("Invalid number of arguments"),
			config: config{
				printUsage: false,
				numTimes:   0,
			},
		},
	}
	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v \n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got: %v \n", err)
		}
		if c.numTimes != tc.config.numTimes {
			t.Fatalf("Expected numTimes to be: %v, got: %v \n", tc.config.numTimes, c.numTimes)
		}
	}
}
