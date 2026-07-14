package main

import "testing"

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		input   User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "正常系: 正しい入力値",
			input:   User{Email: "dev@example.com", Password: "password123"},
			wantErr: false,
		},
		{
			name:    "異常系: @がないメールアドレス",
			input:   User{Email: "example.com", Password: "password123"},
			wantErr: true,
			errMsg:  "無効なメールアドレスの形式です",
		},
		{
			name:    "異常系: @で始まるメールアドレス",
			input:   User{Email: "@example.com", Password: "password123"},
			wantErr: true,
			errMsg:  "無効なメールアドレスの形式です",
		},
		{
			name:    "異常系: @で終わるメールアドレス",
			input:   User{Email: "example@", Password: "password123"},
			wantErr: true,
			errMsg:  "無効なメールアドレスの形式です",
		},
		{
			name:    "異常系: パスワードが短すぎる",
			input:   User{Email: "dev@example.com", Password: "short"},
			wantErr: true,
			errMsg:  "パスワードは8文字以上で入力してください",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUser(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("ValidateUser() error message = %q, want %q", err.Error(), tt.errMsg)
			}
		})
	}
}
