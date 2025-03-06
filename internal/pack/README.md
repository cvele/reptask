`CalculateOptimalPacks` is the core function responsible for determining the most efficient way to fulfill an order with available pack sizes. The function prioritizes: 

1. **Sending the fewest possible extra items** (minimal overshoot). 
2. **Using the smallest number of packs** possible. 

It achieves this using dynamic programming, ensuring the optimal solution for all order quantities. 

```go
func CalculateOptimalPacks(order int) (map[int]int, error)
``` 

**Notes** Initially, a simple greedy method was considered—always selecting the largest available pack size first. However, this method can lead to inefficient outcomes (sending more items or packs than necessary). To guarantee optimal results in every scenario, dynamic programming was chosen instead. 

It calculates the fewest packs needed to reach **every possible order total** up to the maximum overshoot (order + largest pack size). 