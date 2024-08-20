package dbqueries

const (
	AddTransaction           = "INSERT INTO transactions (id, amount, type, parent_id) VALUES ($1, $2, $3, $4)"
	GetTransaction           = "SELECT amount, type, parent_id FROM transactions WHERE id = $1"
	GetTransactionsIDsByType = "SELECT id FROM transactions WHERE type = $1"
	GetTransactionsSumByID   = `WITH RECURSIVE transaction_tree AS (
        							SELECT id, amount, parent_id
        							FROM transactions
        							WHERE id = $1
        							UNION ALL
        							SELECT t.id, t.amount, t.parent_id
        							FROM transactions t
        							INNER JOIN transaction_tree tt ON t.parent_id = tt.id
    							)
    							SELECT SUM(amount) FROM transaction_tree;`
)
