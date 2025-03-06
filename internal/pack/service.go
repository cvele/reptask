package pack

import (
	"errors"
	"log"
	"math"
)

// CalculateOptimalPacks determines the most efficient way to fulfill a customer's order,
// balancing two goals:
// 1. Sending as few extra items as possible beyond the customer's request.
// 2. Using the fewest number of packs possible.
//
// This approach we always find the best solution, even if a greedy approach
// (selecting the largest packs first) might seem intuitive at first.
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

	// Identify the largest pack size available to handle worst-case overshoot scenarios.
	maxPack := packs[0].Size

	// dpLen defines the length of our DP array to account for the largest possible overshoot.
	dpLen := order + maxPack

	// Initialize dp array to store the minimum number of packs required for each possible quantity.
	dp := make([]int, dpLen)

	// packChoice remembers the pack size chosen for each total quantity.
	packChoice := make([]int, dpLen)

	// Initialize dp array with maximum integer value to indicate initially unreachable states
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 // Base case: 0 packs required to fulfill an order of 0 items

	for i := 1; i < dpLen; i++ {
		for _, pack := range packs {
			// Check if current pack can be used to reach a total of i items
			if i-pack.Size >= 0 && dp[i-pack.Size] != math.MaxInt32 {
				// Found a better (smaller) combination of packs for this total.
				if dp[i-pack.Size]+1 < dp[i] {
					dp[i] = dp[i-pack.Size] + 1
					packChoice[i] = pack.Size
				}
			}
		}
	}

	// Find the optimal solution: the smallest total >= order that we can fulfill.
	bestTotalItems := -1
	for totalItems := order; totalItems < dpLen; totalItems++ {
		if dp[totalItems] != math.MaxInt32 {
			bestTotalItems = totalItems
			break
		}
	}

	if bestTotalItems == -1 {
		return nil, errors.New("unable to fulfill order with available pack sizes")
	}

	// Reconstruct the exact distribution of packs based on the computed solution.
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
