package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Item struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	CreatedAt   int    `json:"created_at"`
}

func (a *App) GetItems(search string) []Item {
	query := "SELECT id, code, name, description, quantity, created_at FROM items where name LIKE ? OR code LIKE ? ORDER BY created_at DESC"
	rows, err := a.db.Query(query, "%"+search+"%", "%"+search+"%")
	if err != nil {
		return []Item{}
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		rows.Scan(&item.Id, &item.Code, &item.Name, &item.Description, &item.Quantity, &item.CreatedAt)
		items = append(items, item)
	}

	fmt.Println(items)

	return items
}

func (a *App) CreateItem(code, name, description string, quantity int) (*Item, error) {
	item := &Item{
		Code:        code,
		Name:        name,
		Description: description,
		Quantity:    quantity,
	}

	exists, err := a.checkItemExists(code)
	if err != nil {
		return nil, fmt.Errorf("failed to check code existenc: %w", err)
	}
	if exists {
		return nil, errors.New("code already exists")
	}

	result, err := a.db.Exec(`
    INSERT INTO items (code, name, description, quantity)
    VALUES (?, ?, ?, ?)
  `, code, name, description, quantity)

	if err != nil {
		return nil, fmt.Errorf("failed to create item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last inserted id: %w", err)
	}
	item.Id = int(id)
	return item, nil
}


func (a *App) checkItemExists(code string) (bool, error) {
	var exists bool
	err := a.db.QueryRow(`SELECT EXISTS(SELECT code FROM items WHERE code = ?) from items`, code).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("failed to check code existence: %w", err)
	}
	return exists, nil
}
