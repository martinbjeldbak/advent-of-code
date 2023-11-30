package cmd

import "testing"

func TesttwentytwentytwoDay1(t *testing.T) {
	result := twentytwentytwoDay1([]string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	})

	if result != 24000 {
		t.Errorf("expecting 24000, got %v", result)
	}
}
