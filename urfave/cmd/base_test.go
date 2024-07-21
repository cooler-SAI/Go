package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRunApp(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "hello command",
			args:     []string{"app", "hello"},
			expected: "Hello urfave\n",
		},
		{
			name:     "goodbye command",
			args:     []string{"app", "goodbye"},
			expected: "Goodbye urfave\n",
		},
		{
			name:     "world command",
			args:     []string{"app", "world"},
			expected: "World\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			originalStdout := os.Stdout
			defer func() { os.Stdout = originalStdout }()

			os.Args = tt.args

			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}
			os.Stdout = w

			RunApp()

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(r); err != nil {
				t.Fatalf("failed to read from pipe: %v", err)
			}
			output := buf.String()

			assert.Equal(t, tt.expected, output)
		})
	}
}
