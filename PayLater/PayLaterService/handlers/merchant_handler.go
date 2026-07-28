package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"paylaterservice/generated"
	"paylaterservice/services"
)

// Create Merchant
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

// Get All Merchants
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

// Get Merchant By ID
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

// Update Merchant
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

// Update Commission
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

// Delete Merchant
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