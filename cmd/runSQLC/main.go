package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joaolima7/sqlc-goexpert/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "New Category Test Query",
		Description: sql.NullString{String: "This is a test category created using sqlc", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for i, category := range categories {
		fmt.Println(i, category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:   "e4f1d2f6-f02d-435a-a5d0-6c8489d235e5",
		Name: "Updated Category Name",
		Description: sql.NullString{String: "This is an updated description for the category",
			Valid: true,
		},
	})

	if err != nil {
		panic(err)
	}

	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for i, category := range categories {
		fmt.Println(i, category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, "e4f1d2f6-f02d-435a-a5d0-6c8489d235e5")
	if err != nil {
		panic(err)
	}

}
