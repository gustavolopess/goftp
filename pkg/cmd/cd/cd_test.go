package cd

import (
	"os"
	"testing"
)

func TestChangeDir(t *testing.T) {
	os.Mkdir("./test-dir", 0755)
	defer os.RemoveAll("./test-dir")

	type args struct {
		dirpath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "when directory exists, it should change the current working directory",
			args: args{dirpath: "./test-dir"},
			want: func() string {
				currentDir, _ := os.Getwd()
				return currentDir + "/test-dir"
			}(),
			wantErr: false,
		},
		{
			name:    "when directory does not exist, it should return an error",
			args:    args{dirpath: "/tmp/not-exist-foobarlhebsfoo.txt"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChangeDir(tt.args.dirpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChangeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
