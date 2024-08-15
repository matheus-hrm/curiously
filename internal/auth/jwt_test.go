package auth

import (
	"testing"
)

func TestComparePasswords(t *testing.T) {
	type args struct {
		hashed string
		plain  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Add at least one test case here
		{
			name: "Test Case 1",
			args: args{
				hashed: "hashed value",
				plain:  []byte("plain value"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePasswords(tt.args.hashed, tt.args.plain); got != tt.want {
				t.Errorf("ComparePasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// Add at least one test case here
		{
			name: "Test Case 1",
			args: args{
				password: "password",
			},
			want:    "hashed password",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}