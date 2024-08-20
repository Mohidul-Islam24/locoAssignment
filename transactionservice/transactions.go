package transactionservice

import (
	"database/sql"
	"locoassignment/db"
	"locoassignment/dbqueries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var req struct {
		Amount   float64 `json:"amount"`
		Type     string  `json:"type"`
		ParentID *int    `json:"parent_id"` // pointer to handle optional value
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err = db.Driver.Exec(dbqueries.AddTransaction, id, req.Amount, req.Type, req.ParentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetTransaction(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var amount float64
	var type_ string
	var parentID sql.NullInt64

	err = db.Driver.QueryRow(dbqueries.GetTransaction, id).Scan(&amount, &type_, &parentID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"amount":    amount,
		"type":      type_,
		"parent_id": parentID.Int64,
	})
}

func GetTransactionsIDsByType(c *gin.Context) {
	type_ := c.Param("type")

	rows, err := db.Driver.Query(dbqueries.GetTransactionsIDsByType, type_)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse result"})
			return
		}
		ids = append(ids, id)
	}

	c.JSON(http.StatusOK, ids)
}

func GetTransactionsSumByID(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var totalSum float64
	err = db.Driver.QueryRow(dbqueries.GetTransactionsSumByID, id).Scan(&totalSum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate sum"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sum": totalSum})
}
