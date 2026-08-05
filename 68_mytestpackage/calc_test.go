package calc

import "testing"

// func TestAdd(t *testing.T) {
// 	got := Add(2, 3)
// 	want := 5
// 	if got != want {
// 		t.Errorf("Add(2,3) = %d,want %d", got, want)
// 	}
// }

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 0, 0},
		{"negative", -1, -2, -3},
		{"mixed", -1, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	// Errorf: テスト失敗を記録し、実行を続行
	got, err := Divide(10, 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got < 3.33 || got > 3.34 {
		t.Errorf("Divide(10,3) = %f,want ~3.33", got)
	}

	// Fatalf: テスト失敗を記録し、即時に停止
	_, err = Divide(10, 0)
	if err == nil {
		t.Fatalf("expected error for division by zero")
	}
}
