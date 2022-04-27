package ftp

import (
	"io"
	"reflect"
	"testing"
)

func TestNewFtpClient(t *testing.T) {
	tests := []struct {
		name string
		want FtpClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFtpClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFtpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ResponseHandler(t *testing.T) {
	testHandlersCalls := make(map[string]int)
	increaseCall := func(b []byte) {
		testHandlersCalls[string(b)]++
	}

	tests := []struct {
		name             string
		responseHandlers map[string]func([]byte)
		command          string
		response         []byte
		expectCalls      int
		wantErr          bool
	}{
		{
			name: "when command's verb handler exists, it should call this handler",
			responseHandlers: map[string]func([]byte){
				"existentCommand": increaseCall,
			},
			command:     "existentCommand",
			response:    []byte("existentCommandResponse"),
			expectCalls: 1,
			wantErr:     false,
		},
		{
			name:             "when command's verb handler does not exist, it should not call the handler and return an error",
			responseHandlers: map[string]func([]byte){},
			command:          "nonExistentCommand",
			response:         []byte("nonExistentCommandResponse"),
			expectCalls:      0,
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				ResponseHandlers: tt.responseHandlers,
			}
			err := c.ResponseHandler(tt.command, tt.response)

			calls := testHandlersCalls[string(tt.response)]
			if calls != tt.expectCalls {
				t.Errorf("%v handler has been called with response '%v' %v times, but expected was %v times", tt.command, string(tt.response), calls, tt.expectCalls)
			}

			hadErr := err != nil
			if hadErr != tt.wantErr {
				t.Errorf("client.ResponseHandler err = %v and wantErr is %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CommandHandler(t *testing.T) {
	type fields struct {
		ResponseHandlers map[string]func([]byte)
	}
	type args struct {
		command string
		reader  io.ReadWriteCloser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				ResponseHandlers: tt.fields.ResponseHandlers,
			}
			c.CommandHandler(tt.args.command, tt.args.reader)
		})
	}
}
