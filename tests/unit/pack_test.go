package pack_test

import (
	"os"
	"testing"

	"github.com/cvele/reptask/internal/db"
	"github.com/cvele/reptask/internal/pack"
)

func TestMain(m *testing.M) {
	db.InitDB("/tmp/packs_test.db")
	defer db.CloseDB()

	code := m.Run()

	os.Exit(code)
}

func seedDatabase(t *testing.T) {
	t.Helper()
	_, err := db.DB.Exec("DELETE FROM packs")
	if err != nil {
		t.Fatalf("Failed to clear packs table: %v", err)
	}
	db.SeedDefaultPackSizes()
}
func TestCalculateOptimalPacks(t *testing.T) {
	seedDatabase(t)

	tests := []struct {
		name     string
		order    int
		expected map[int]int
		wantErr  bool
	}{
		{
			name:  "Exact match with multiple packs",
			order: 12001,
			expected: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			wantErr: false,
		},
		{
			name:  "Single pack exact match",
			order: 5000,
			expected: map[int]int{
				5000: 1,
			},
			wantErr: false,
		},
		{
			name:  "Minimal order uses smallest pack",
			order: 1,
			expected: map[int]int{
				250: 1,
			},
			wantErr: false,
		},
		{
			name:  "Order slightly larger than pack size (251)",
			order: 251,
			expected: map[int]int{
				500: 1,
			},
			wantErr: false,
		},
		{
			name:  "Order slightly larger than pack size (501)",
			order: 501,
			expected: map[int]int{
				500: 1,
				250: 1,
			},
			wantErr: false,
		},
		{
			name:  "Combination of packs (1750)",
			order: 1750,
			expected: map[int]int{
				1000: 1,
				500:  1,
				250:  1,
			},
			wantErr: false,
		},
		{
			name:     "Large order",
			order:    123456789,
			expected: nil, // Validate just that it doesn't error out for large inputs.
			wantErr:  false,
		},
		{
			name:     "Invalid order quantity",
			order:    -1,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Zero order quantity",
			order:    0,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := pack.CalculateOptimalPacks(tt.order)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Expected error: %v, got: %v", tt.wantErr, err)
			}

			if !tt.wantErr && tt.expected != nil && !equalMaps(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

// helper function to compare two maps
func equalMaps(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
