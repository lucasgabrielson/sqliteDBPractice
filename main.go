package main

import (
    "database/sql"
    "fmt"
    "log"
	"net/http"
	"github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
	"github.com/blockloop/scan"
)

type Payment struct {
	ID int `db:"id"`
    Name string
    Email string
	AmountCents int `db:"amount_cents"`
	Datetime int 
	Status string 
}

func openDB(path string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        log.Printf("failed to open sqlite db")
        return nil, err
    }
    if err = db.Ping(); err != nil {
        log.Printf("failed to keep db connection")
        return nil, err
    }
    return db, nil
}

func getPayments(db *sql.DB, date string, dbStatus string) ([]Payment) {
	var payments []Payment
	var query string
	if date != "" && dbStatus != "" {
		query = fmt.Sprintf("SELECT id, name, email, amount_cents, datetime, status FROM payments WHERE datetime = %s AND status = '%s'", date, dbStatus)
	} else if date != "" {
		query = fmt.Sprintf("SELECT id, name, email, amount_cents, datetime, status  FROM payments WHERE datetime = %s", date)
	} else if dbStatus != "" {
		query = fmt.Sprintf("SELECT id, name, email, amount_cents, datetime, status  FROM payments WHERE status = '%s'", dbStatus)
	} else {
		query = fmt.Sprintf("SELECT id, name, email, amount_cents, datetime, status  FROM payments")
	}

    rows, err := db.Query(query)
    if err != nil {
        fmt.Printf("error obtaining rows from db: %v", err)
    }
    defer rows.Close()

	err = scan.Rows(&payments, rows)
	if err != nil {
		fmt.Printf("error scanning rows: %v", err)
	}
	return payments
}

func main() {
	r := gin.Default()

	r.GET("/payments", func(c *gin.Context){
		db, _ := openDB("./take_home.db")
   		defer db.Close()
		datetime := c.Query("date")
		status := c.Query("status")
    	payments := getPayments(db, datetime, status)
		c.JSON(http.StatusOK, payments)
	})

	r.Run()
}
