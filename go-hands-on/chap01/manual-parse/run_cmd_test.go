package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		{
			c:      config{printUsage: true},
			output: usageString,
		},
		{
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("Your name please? Press the Enter key when done. \n", 1),
			err:    errors.New("You didn't enter your name"),
		},
		{
			c:      config{numTimes: 5},
			input:  "Bill Bryson",
			output: "Your name please? Press the Enter key when done. \n" + strings.Repeat("Nice to meet you Bill Bryson\n", 5),
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		rd := strings.NewReader(tc.input)
		err := runCmd(rd, byteBuf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("Expected nil error, got: %v \n", err)
		}
		if err == nil && tc.err != nil {
			t.Fatalf("Expected error to be: %v, got: %v \n", tc.err, err)
		}
		gotMsg := byteBuf.String()
		if gotMsg != tc.output {
			t.Fatalf("Expected output to be: %v, got: %v \n", tc.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
