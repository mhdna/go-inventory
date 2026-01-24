package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	ctx context.Context
	db  *sql.DB
}

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

	_, err = a.db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code TEXT,
			name TEXT,
			description TEXT,
			quantity INTEGER,
			created_at INTEGER DEFAULT (strftime('%s', 'now'))
		)
	`)
	if err != nil {
		panic(err)
	}

}

func (a *App) Login(name, password string) bool {
	return name == "mh" && password == "pass"
}

func (a *App) Print(msg string) {
	fmt.Println(msg)
}
