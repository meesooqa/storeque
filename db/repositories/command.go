package repositories

import (
	"context"
	"database/sql"

	"github.com/meesooqa/storeque/common/domain"
	"github.com/meesooqa/storeque/db/entities"
)

type CommandRepository struct {
	db      *sql.DB
	adapter *entities.CommandAdapter
}

func NewCommandRepository(db *sql.DB) *CommandRepository {
	return &CommandRepository{
		db:      db,
		adapter: entities.NewCommandAdapter(),
	}
}

func (this *CommandRepository) FindByRoleID(ctx context.Context, roleID int64) ([]*domain.Command, error) {
	query := `
		SELECT c.id, c.code, r.id AS role_id, r.code AS role_code
		FROM commands c
		INNER JOIN role_commands rc ON c.id = rc.command_id
		INNER JOIN roles r ON r.id = rc.role_id
		WHERE rc.role_id = $1
	`

	rows, err := this.db.QueryContext(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commands []*domain.Command
	for rows.Next() {
		item := &entities.Command{}
		err = rows.Scan(&item.ID, &item.Code, &item.RoleID, &item.Role)
		if err != nil {
			return nil, err
		}
		commands = append(commands, this.adapter.ToDomain(item))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
