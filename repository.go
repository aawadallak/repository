package repository

import "database/sql"

func Find[T any](conn sql.DB, entity any) (*T, error) {}
