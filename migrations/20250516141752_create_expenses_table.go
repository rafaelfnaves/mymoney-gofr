package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const create_table = `CREATE TABLE IF NOT EXISTS expenses (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	value NUMERIC(10, 2) NOT NULL,
	payment_type VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

func create_expenses_table() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(create_table)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
