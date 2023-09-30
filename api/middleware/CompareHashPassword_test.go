package middleware

import "testing"

func TestShouldCompareHashPassword(t *testing.T) {
	var tests = []struct {
		password string
		hash     string
		expected bool
	}{
		{"password", "$2a$12$7g5qfLrUyFeBOQjrY6nfaeEmapwSVPMsjRiEuf8aXpz0EkytaQP6C", true},
		{"123456", "$2a$12$Ep56IJ9YghGF1KAJ/bbt6.rBwy.WfyTQ.xR/oNwiMt5k6mztdiCBu", true},
		{"foobar", "$2a$12$U6xedOPvqG2JVxmvOVao8OjuD6Yglg.z7.PtSsMVuj48BzuhoelHy", true},
		{"badpassword", "$2a$12$U6xedOPvqG2JVxmvOVao8OjuD6Yglg.z7.PtSsMVuj48BzuhoelHy", false},
		{"zirgim", "$2a$12$U6xedOPvqG2JVxmvOVao8OjuD6Yglg.z7.PtSsMVuj48BzuhoelHy", false},
	}

	for _, test := range tests {
		test := test

		t.Run(test.password, func(t *testing.T) {
			t.Parallel()

			actual := CompareHashPassword(test.password, test.hash)

			if test.expected != actual {
				t.Errorf("The hash of %q is not %t.", test.password, test.expected)
			}
		})
	}
}
