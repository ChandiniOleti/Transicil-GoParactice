package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"paylaterservice/services"
)


// GET /reports/merchant-fee/:id
func MerchantFeeReport(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid merchant id",
		})
		return
	}


	result, err := services.GetMerchantFeeReport(int32(id))

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, result)
}



// GET /reports/users-with-due
func UsersWithDueReport(c *gin.Context) {

	result, err := services.GetUsersWithDueReport()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, result)
}



// GET /reports/user-due/:id
func UserDueReport(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid user id",
		})
		return
	}


	result, err := services.GetUserDueReport(int32(id))

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, result)
}



// GET /reports/credit-limit-users
func CreditLimitUsersReport(c *gin.Context) {

	result, err := services.GetCreditLimitUsersReport()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, result)
}



// GET /reports/total-dues
func TotalDuesReport(c *gin.Context) {

	result, err := services.GetTotalDuesReport()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, result)
}