package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

var (
	testFiles = "test_files"
)

func TestRun(t *testing.T) {
	type args struct {
		args  []string
		stdin string
	}
	var tests = []struct {
		name       string
		args       args
		wantStdout string
		wantErr    bool
	}{
		{"no args no stdin", args{[]string{"fvv"}, "empty"}, "", true},
		{"invalid template", args{[]string{"fvv"}, "bad.txt.tmpl"}, "", true},
		{"bad flag", args{[]string{"fvv", "-z"}, "empty"}, "", true},
		{"happy path", args{[]string{"fvv"}, "first.txt.tmpl"}, "ONE\nStandard stuff\nTWO\n", false},
		{"multi path", args{[]string{"fvv"}, "many.txt.tmpl"}, "## Middle One\ntop\nmiddle_one\n\n## Middle Two\ntop\nmiddle_two\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stdin, err := os.Open(fmt.Sprintf("%s/%s", testFiles, tt.args.stdin))
			if err != nil {
				t.Errorf("Failed opening test file %s: %v", tt.args.stdin, err)
			}

			err = run(tt.args.args, stdin, stdout)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("run() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
		})
	}
}
