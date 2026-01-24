package main

import (
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
