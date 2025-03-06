package pack

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPacksHandler(c *fiber.Ctx) error {
	packs, err := GetAllPacks()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(packs)
}

func GetPackByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pack ID"})
	}

	pack, err := GetPackByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	if pack == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Pack not found"})
	}

	return c.JSON(pack)
}

func CreatePackHandler(c *fiber.Ctx) error {
	var req struct {
		Size int `json:"size"`
	}
	if err := c.BodyParser(&req); err != nil || req.Size <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input, size must be positive"})
	}
	if err := AddPack(req.Size); err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Pack size already exists"})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Pack size added"})
}

func UpdatePackHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pack ID"})
	}
	var req struct {
		Size int `json:"size"`
	}
	if err := c.BodyParser(&req); err != nil || req.Size <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input, size must be positive"})
	}
	if err := UpdatePack(id, req.Size); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update pack size"})
	}
	return c.JSON(fiber.Map{"message": "Pack size updated"})
}

func DeletePackHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pack ID"})
	}
	if err := DeletePack(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete pack"})
	}
	return c.JSON(fiber.Map{"message": "Pack deleted"})
}

// CalculatePacksHandler handles the request to calculate the optimal pack distribution.
// It validates the input, sorts the result, and returns a JSON response.
func CalculatePacksHandler(c *fiber.Ctx) error {
	orderStr := c.Query("order")
	order, err := strconv.Atoi(orderStr)
	if err != nil || order <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order quantity"})
	}

	packs, err := CalculateOptimalPacks(order)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to calculate packs"})
	}

	// Sort the packs by size in descending order
	sortedPacks := make([]struct {
		Size  int `json:"size"`
		Count int `json:"count"`
	}, 0, len(packs))

	for size, count := range packs {
		sortedPacks = append(sortedPacks, struct {
			Size  int `json:"size"`
			Count int `json:"count"`
		}{Size: size, Count: count})
	}

	sort.Slice(sortedPacks, func(i, j int) bool {
		return sortedPacks[i].Size > sortedPacks[j].Size
	})

	return c.JSON(sortedPacks)
}
