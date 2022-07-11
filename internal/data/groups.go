package data

import (
	"context"
	"time"
)

type Group struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GET/groups
func (g *Group) GetAllGroups() ([]*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select * from groups
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var groups []*Group
	for rows.Next() {
		var group Group
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.CreatedAt,
			&group.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}
	return groups, nil
}

func (g *Group) GetGroupById(id int) (*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, name, created_at, updated_at from groups where id = $1`
	var group Group

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&group.ID,
		&group.Name,
		&group.CreatedAt,
		&group.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// POST/groups/create
func (g *Group) CreateGroup(group Group) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newId int // TODO: create serialize new ID
	query := `
		insert into groups (name, created_at, updated_at)
		values ($1, $2, $3) returning id
	`
	err := db.QueryRowContext(
		ctx,
		query,
		group.Name,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil

}
