package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"merchantservice/generated"
	"merchantservice/services"
)

// CreateMerchant creates a new merchant.
// POST /merchants
func CreateMerchant(c *gin.Context) {
	var merchant generated.CreateMerchantParams

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := services.CreateMerchant(merchant)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Merchant Created Successfully",
	})
}

// GetMerchants returns all merchants (no passwords).
// GET /merchants
func GetMerchants(c *gin.Context) {
	merchants, err := services.GetMerchants()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, merchants)
}

// GetMerchantByID returns merchant details (no password).
// GET /merchants/:id
func GetMerchantByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	merchant, err := services.GetMerchantByID(int32(id))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, merchant)
}

// UpdateMerchant updates merchant details.
// PUT /merchants/:id
func UpdateMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	var merchant generated.UpdateMerchantParams

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	merchant.ID = int32(id)

	err = services.UpdateMerchant(merchant)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Merchant Updated Successfully",
	})
}

// UpdateCommission updates only the commission percentage.
// PATCH /merchants/:id/commission
func UpdateCommission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	var merchant generated.UpdateCommissionParams

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	merchant.ID = int32(id)

	err = services.UpdateCommission(merchant)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Commission Updated Successfully",
	})
}

// DeleteMerchant deletes a merchant.
// DELETE /merchants/:id
func DeleteMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	err = services.DeleteMerchant(int32(id))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Merchant Deleted Successfully",
	})
}

// GetMerchantCommission returns only commission (internal only).
// GET /internal/merchants/:id/commission
func GetMerchantCommission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Merchant ID",
		})
		return
	}

	commission, err := services.GetMerchantCommission(int32(id))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, commission)
}
