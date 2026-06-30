package main

import "testing"

func TestCheckAge(t *testing.T) {
	tests := []struct {
		name      string
		age       int
		expect    string
		expectErr bool
	}{
		{name: "正常系:18歳ぴったり", age: 18, expect: "アクセス許可", expectErr: false},
		{name: "正常系:成人", age: 25, expect: "アクセス許可", expectErr: false},
		{name: "異常系:18歳未満", age: 15, expect: "", expectErr: true},
		{name: "異常系:負の数", age: -5, expect: "", expectErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkAge(tt.age)

			if (err != nil) != tt.expectErr {
				t.Errorf("checkAge() error = %v, expectErr %v", err, tt.expectErr)

				return
			}

			if got != tt.expect {
				t.Errorf("checkAge() got = %v, expect %v", got, tt.expect)
			}
		})
	}
}
