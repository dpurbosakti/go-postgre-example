package service

import (
	"learn-echo/features/users/model/dto"
	"reflect"
	"testing"
)

func TestUserServiceImpl_Create(t *testing.T) {
	type args struct {
		input dto.UserCreateRequest
	}
	tests := []struct {
		name       string
		service    *UserServiceImpl
		args       args
		wantResult dto.UserResponse
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.service.Create(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UserServiceImpl.Create() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
