package numToWords

import "testing"

func RangeableMap[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))

    for k := range m {
        keys = append(keys, k)
    }

    return keys
}

func TestLargeNumbers(t *testing.T) {
	expectations := map[uint64]string{
		999:                     "nine hundred ninety-nine",
		999_999:                 "nine hundred ninety-nine thousand nine hundred ninety-nine",
		999_999_999:             "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		999_999_999_999:         "nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	}

	for number, wanted := range expectations {
		got := Convert(number)
		if got != wanted {
			t.Errorf("Wanted `%s`, but got `%s` (number: %d)", wanted, got, number)
		}
	}
}
