package utils

import (
	"os"
	"testing"
)

func TestTryReadFile(t *testing.T) {
	testFileName := "file_test.go"
	var testFileContents string
	if data, err := os.ReadFile(testFileName); err == nil {
		testFileContents = string(data)
	} else {
		t.Errorf("TryReadFile() error reading test file: %w", err)
	}

	type args struct {
		filename string
		attempts int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Invalid number of attempts",
			args: args{
				filename: "",
				attempts: 0,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "File not exists",
			args: args{
				filename: "non-existent-file-name",
				attempts: 1,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Successful read",
			args: args{
				filename: testFileName,
				attempts: 1,
			},
			want:    testFileContents,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryReadFile(tt.args.filename, tt.args.attempts)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TryReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
