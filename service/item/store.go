package item

import (
	"database/sql"

	"github.com/Ion-Stefan/go-kickstart-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateItem(item types.Item) error {
	_, err := s.db.Exec(
		"INSERT INTO items (name, description, imageURL) VALUES (?, ?, ?)",
		item.Name,
		item.Description,
		item.ImageURL,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetItems() (*[]types.Item, error) {
	rows, err := s.db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}

	items := make([]types.Item, 0)
	for rows.Next() {
		item, err := scanRowsIntoItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	return &items, nil
}

func scanRowsIntoItem(rows *sql.Rows) (*types.Item, error) {
	item := new(types.Item)

	err := rows.Scan(
		&item.CreatedAt,
		&item.Name,
		&item.Description,
		&item.ImageURL,
		&item.ID,
	)
	if err != nil {
		return nil, err
	}

	return item, nil
}
