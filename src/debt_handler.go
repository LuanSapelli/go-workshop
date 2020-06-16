package src

import (
	"github.com/gin-gonic/gin"
)

//GetDebts (todas as debts existentes)
func GetDebts(c *gin.Context) {

	var debts []debt
	allDebts := selectAll(&debts, c)

	c.JSON(200, allDebts)
}

//GetDebt (debt pelo ID)
func GetDebt(c *gin.Context) {
	ID := c.Param("id")

	debt, _ := selectDebtID(ID, c)

	c.JSON(200, debt)
}

//PostDebt (criar debt)
func PostDebt(c *gin.Context) {

	var newDebt debt

	if err := c.ShouldBindJSON(&newDebt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbConnect()
	db.Create(&newDebt)

	c.JSON(201, newDebt)
}

//PutDebt (update debt)
func PutDebt(c *gin.Context) {

	ID := c.Param("id")

	var updateDebt debt

	if err := c.ShouldBindJSON(&updateDebt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	debt, db := selectDebtID(ID, c)

	debt.CompanyName = updateDebt.CompanyName
	debt.Value = updateDebt.Value
	debt.Date = updateDebt.Date
	debt.Status = updateDebt.Status

	db.Save(&debt)

	c.JSON(200, debt)
}

//DeleteDebt remove one user
func DeleteDebt(c *gin.Context) {

	ID := c.Param("id")

	debt, db := selectDebtID(ID, c)

	if debt.ID != " " {
		db.Delete(&debt)
	}

	c.JSON(204, nil)
}
