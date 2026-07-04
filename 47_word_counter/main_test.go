package main

import (
	"maps"
	"strings"
	"testing"
)

func TestCountWord(t *testing.T) {
	tests := []struct {
		name   string
		word   []string
		expect map[string]int
	}{
		{name: "正常系", word: strings.Fields("go python go rust go python java"), expect: map[string]int{"go": 3, "python": 2, "rust": 1, "java": 1}},
		{name: "空のリスト", word: []string{}, expect: map[string]int{}},
		{name: "引数がnilの場合", word: nil, expect: map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWord(tt.word)

			if !maps.Equal(got, tt.expect) {
				t.Errorf("%s の結果が一致しません\n期待値: %v\n実際の値: %v", tt.name, tt.expect, got)
			}
		})
	}
}
