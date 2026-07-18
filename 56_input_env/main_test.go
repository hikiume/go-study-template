package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	oldPort := os.Getenv("APP_PORT")
	oldURL := os.Getenv("API_URL")

	t.Cleanup(func() {
		os.Setenv("APP_PORT", oldPort)
		os.Setenv("API_URL", oldURL)
	})

	tests := []struct {
		name     string
		setEnv   func()
		wantPort int
		wantErr  bool
	}{
		{
			name: "正常系: 全て正しく設定されている場合",
			setEnv: func() {
				os.Setenv("APP_PORT", "9000")
				os.Setenv("API_URL", "https://test.com")
			},
			wantPort: 9000,
			wantErr:  false,
		},
		{
			name: "正常系: ポートが未指定ならデフォルト値が(8080)になる",
			setEnv: func() {
				os.Unsetenv("APP_PORT")
				os.Setenv("API_URL", "https://test.com")
			},
			wantPort: 8080,
			wantErr:  false,
		},
		{
			name: "異常系: ポートが数値ではない",
			setEnv: func() {
				os.Setenv("APP_PORT", "abc")
				os.Setenv("API_URL", "https://test.com")
			},
			wantErr: true,
		},
		{
			name: "異常系: 必須のURLが空っぽ",
			setEnv: func() {
				os.Setenv("APP_PORT", "8080")
				os.Unsetenv("API_URL")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv()

			cfg, err := LoadConfig()

			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && cfg.Port != tt.wantPort {
				t.Errorf("cfg.Port = %d, want %d", cfg.Port, tt.wantPort)
			}
		})
	}
}
