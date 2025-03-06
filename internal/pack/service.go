// NOTE FROM DEV: Initially considered using a greedy algorithm (always picking the largest available pack first),
// but realized this could lead to suboptimal results (excess items or more packs than necessary).
// Switched to dynamic programming to ensure the optimal solution in all cases.

package pack

import (
	"errors"
	"log"
	"math"
)

// CalculateOptimalPacks determines optimal pack distribution for the given order quantity.
// It uses dynamic programming to minimize the total number of items sent first,
// and secondly minimizes the total number of packs used.
func CalculateOptimalPacks(order int) (map[int]int, error) {
	if order <= 0 {
		return nil, errors.New("order must be greater than zero")
	}

	packs, err := GetAllPacks()
	if err != nil {
		log.Println("Failed to retrieve pack sizes:", err)
		return nil, errors.New("failed to retrieve pack sizes")
	}

	if len(packs) == 0 {
		return nil, errors.New("no pack sizes available")
	}

	// Determine the largest available pack size
	maxPack := packs[0].Size

	// Dynamic programming array length set to cover all possible minimal overshoots
	dpLen := order + maxPack

	// dp[i] stores the minimum number of packs required to fulfill an order of exactly i items
	dp := make([]int, dpLen)

	// packChoice[i] stores the pack size chosen to reach a total of exactly i items optimally
	packChoice := make([]int, dpLen)

	// Initialize dp array with maximum integer value to indicate initially unreachable states
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 // Base case: 0 packs required to fulfill an order of 0 items

	// Populate the dp array
	for i := 1; i < dpLen; i++ {
		for _, pack := range packs {
			// Check if current pack can be used to reach a total of i items
			if i-pack.Size >= 0 && dp[i-pack.Size] != math.MaxInt32 {
				// Update dp[i] if a better (fewer packs) solution is found
				if dp[i-pack.Size]+1 < dp[i] {
					dp[i] = dp[i-pack.Size] + 1
					packChoice[i] = pack.Size
				}
			}
		}
	}

	// Find the smallest number of items greater than or equal to the order that can be fulfilled
	bestTotalItems := -1
	for totalItems := order; totalItems < dpLen; totalItems++ {
		if dp[totalItems] != math.MaxInt32 {
			bestTotalItems = totalItems
			break
		}
	}

	// Return an error if no valid solution exists
	if bestTotalItems == -1 {
		return nil, errors.New("unable to fulfill order with available pack sizes")
	}

	// Reconstruct the optimal pack distribution from the dp solution
	result := make(map[int]int)
	remaining := bestTotalItems
	for remaining > 0 {
		packSize := packChoice[remaining]
		if packSize == 0 {
			log.Printf("DP table error: packChoice[%d] = 0\n", remaining)
			return nil, errors.New("internal error in pack calculation")
		}
		result[packSize]++
		remaining -= packSize
	}

	log.Printf("Order %d: Pack distribution %v (Total items sent: %d)\n", order, result, bestTotalItems)

	return result, nil
}
