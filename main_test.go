package main

import (
	"bytes"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		args     []string
		fileGlob string
	}
	var tests = []struct {
		name       string
		args       args
		wantStdout string
		wantErr    bool
	}{
		{"no args", args{[]string{"fvv"}, ""}, "", true},
		{"missing template", args{[]string{"fvv", "t3"}, "*.tmpl"}, "", true},
		{"definition does not exist", args{[]string{"fvv", "t20"}, "test_files/*.tmpl"}, "", true},
		{"happy path", args{[]string{"fvv", "T3"}, "test_files/*.tmpl"}, "ONE\nStandard stuff\nTWO", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			err := run(tt.args.args, tt.args.fileGlob, stdout)
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
