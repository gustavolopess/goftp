package get

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	os.WriteFile("/tmp/exist.txt", []byte("hello world"), 0644)
	defer os.Remove("/tmp/exist.txt")

	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		before  func()
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "when file exists, it should return its content",
			args:    args{filePath: "/tmp/exist.txt"},
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "when file does not exist, it should return an error",
			args:    args{filePath: "/tmp/not-exist-foobarlhebsfoo.txt"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
