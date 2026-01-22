package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

type Item struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	a.db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	// Create table if it doesn't exist
	_, err = a.db.Exec(`
		CREATE TABLE IF NOT EXISTS items (

			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code TEXT,
			name TEXT,
			qty INTEGER
		)
	`)
	if err != nil {
		panic(err)
	}

}

func (a *App) GetItems(search string) []Item {
	query := "SELECT id, code, name, qty FROM items where name LIKE ? OR code LIKE ?"
	rows, err := a.db.Query(query, "%"+search+"%", "%"+search+"%")
	if err != nil {
		return []Item{}
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		rows.Scan(&item.Id, &item.Code, &item.Name, &item.Qty)
		items = append(items, item)
	}

	fmt.Println(items)

	return items
}

// Greet returns a greeting for the given name
func (a *App) Login(name, password string) bool {
	return name == "mh" && password == "pass"
}

func (a *App) Print(msg string) {
	fmt.Println(msg)
}
