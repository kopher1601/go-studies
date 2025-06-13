package main

import (
	"os"
	"os/exec"
	"testing"
)

var (
	binName = "app"
)

func TestMain(m *testing.M) {
	// 빌드
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		os.Exit(1)
	}

	// 테스트 실행
	exitCode := m.Run()

	// 정리
	os.Remove(binName)
	os.Exit(exitCode)
}

func TestCLI(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
		exitCode int
	}{
		{
			name:     "도움말 표시",
			args:     []string{"-h"},
			expected: "Must specify a number greater than 0",
			exitCode: 1,
		},
		{
			name:     "잘못된 인수",
			args:     []string{"invalid"},
			expected: "strconv.Atoi: parsing",
			exitCode: 1,
		},
		{
			name:     "음수 입력",
			args:     []string{"-1"},
			expected: "Must specify a number greater than 0",
			exitCode: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("./"+binName, tc.args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					if exitErr.ExitCode() != tc.exitCode {
						t.Errorf("Expected exit code %d, got %d", tc.exitCode, exitErr.ExitCode())
					}
				} else {
					t.Fatalf("Failed to execute command: %v", err)
				}
			}

			if string(output) == "" && tc.expected != "" {
				t.Errorf("Expected output to contain %q, got empty string", tc.expected)
			} else if string(output) != "" && tc.expected != "" {
				if !contains(string(output), tc.expected) {
					t.Errorf("Expected output to contain %q, got %q", tc.expected, string(output))
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
