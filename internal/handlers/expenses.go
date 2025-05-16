package handlers

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafaelfnaves/mymoney-gofr/internal/models"
	"gofr.dev/pkg/gofr"
)

func AddExpense(ctx *gofr.Context) (any, error) {
	var expense models.Expense
	err := ctx.Bind(&expense)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	createdAt := time.Now()
	_, err = ctx.SQL.ExecContext(ctx,
		"INSERT INTO expenses (id, title, description, value, payment_type, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		id, expense.Title, expense.Description, expense.Value, expense.PaymentType, createdAt,
	)
	if err != nil {
		return nil, err
	}

	expense.ID = id
	expense.CreatedAt = createdAt

	return expense, nil
}

func ListExpenses(ctx *gofr.Context) (any, error) {
	var expenses []models.Expense

	rows, err := ctx.SQL.QueryContext(
		ctx,
		"SELECT id, title, description, value, payment_type, created_at FROM expenses",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.Title, &expense.Description, &expense.Value, &expense.PaymentType, &expense.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetExpense(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	var expense models.Expense

	err := ctx.SQL.QueryRowContext(
		ctx,
		"SELECT id, title, description, value, payment_type, created_at FROM expenses WHERE id = $1",
		id,
	).Scan(&expense.ID, &expense.Title, &expense.Description, &expense.Value, &expense.PaymentType, &expense.CreatedAt)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func DeleteExpense(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")

	_, err := ctx.SQL.ExecContext(
		ctx,
		"DELETE FROM expenses WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
