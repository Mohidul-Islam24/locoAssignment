package main

import (
	"locoassignment/transactionservice"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.PUT("/transactionservice/transaction/:transaction_id", transactionservice.AddTransaction)
	r.GET("/transactionservice/transaction/:transaction_id", transactionservice.GetTransaction)
	r.GET("/transactionservice/types/:type", transactionservice.GetTransactionsIDsByType)
	r.GET("/transactionservice/sum/:transaction_id", transactionservice.GetTransactionsSumByID)

	r.Run(":8080")
}
