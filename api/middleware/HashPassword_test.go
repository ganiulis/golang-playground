package middleware

import "testing"

func TestShouldHashPassword(t *testing.T) {
	var tests = []struct {
		password string
	}{
		{"password"},
		{"123456"},
		{"foobar"},
	}

	for _, test := range tests {
		test := test

		t.Run(test.password, func(t *testing.T) {
			t.Parallel()

			actual, err := HashPassword(test.password)

			if nil != err {
				t.Errorf("Something went wrong.")
			}

			if test.password == actual {
				t.Error()
			}
		})
	}
}
