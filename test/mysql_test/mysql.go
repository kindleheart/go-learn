package main

import (
	"database/sql"
	"fmt"
)

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			fmt.Printf("Rollback, err:%v\n", err)
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("UPDATE products SET views = views + 1")
	if err != nil {
		fmt.Println(err)
		return
	}

	id, _ := result.LastInsertId()
	affected, _ := result.RowsAffected()
	fmt.Println(id, affected)

	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}

func main() {
	// @NOTE: the real connection is not required for tests
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}
}
