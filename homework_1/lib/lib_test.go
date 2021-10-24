package lib

import "testing"

func TestChangeLetterCase(t *testing.T) {
	var (
		testCase   = []string{"alO", "laLAlaA", "aAaaaAAaaA"}
		testResult = []string{"ALo", "LAlaLAa", "AaAAAaaAAa"}
	)

	for i := range testCase {
		if result := ChangeLetterCase(testCase[i]); result != testResult[i] {
			t.Errorf("Result was incorrect, got: %s, want: %s", result, testResult[i])
		}
	}
}
