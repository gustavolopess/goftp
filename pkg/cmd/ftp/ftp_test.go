package ftp

import "testing"

type ftpMock struct {
	callsToGet int8
}

func (f *ftpMock) Get(filePath string) (string, error) {
	f.callsToGet += 1
	return "", nil
}

func Test_ftp_Get(t *testing.T) {
	type fields struct {
		dir string
	}
	type args struct {
		relativePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &ftp{
				dir: tt.fields.dir,
			}
			got, err := f.Get(tt.args.relativePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ftp.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ftp.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
