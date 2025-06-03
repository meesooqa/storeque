package repositories

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Failed to open in-memory database")

	_, err = db.Exec(`
		CREATE TABLE commands (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			code TEXT NOT NULL UNIQUE
		);
		
		CREATE TABLE roles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			code TEXT NOT NULL UNIQUE
		);
		
		CREATE TABLE role_commands (
			role_id INTEGER REFERENCES roles(id),
			command_id INTEGER REFERENCES commands(id),
			PRIMARY KEY (role_id, command_id)
		);
	`)
	require.NoError(t, err, "Failed to create tables")

	return db
}

func seedTestData(t *testing.T, db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO roles (id, code) VALUES 
		(1, 'admin'),
		(2, 'moderator'),
		(3, 'user')
	`)
	require.NoError(t, err, "Failed to seed roles")

	_, err = db.Exec(`
		INSERT INTO commands (id, code) VALUES 
		(101, 'create_user'),
		(102, 'delete_user'),
		(103, 'edit_user'),
		(104, 'view_user'),
		(105, 'ban_user')
	`)
	require.NoError(t, err, "Failed to seed commands")

	_, err = db.Exec(`
		INSERT INTO role_commands (role_id, command_id) VALUES 
		(1, 101), (1, 102), (1, 103), (1, 104), (1, 105), -- admin: все права
		(2, 103), (2, 104), (2, 105),                    -- moderator: edit, view, ban
		(3, 104)                                         -- user: view
	`)
	require.NoError(t, err, "Failed to seed role_commands")
}

func TestCommandRepository_FindByRoleID_InMemory(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB(t)
	seedTestData(t, db)
	repo := NewCommandRepository(db)

	t.Run("admin role has all commands", func(t *testing.T) {
		commands, err := repo.FindByRoleID(ctx, 1)
		require.NoError(t, err)
		require.Len(t, commands, 5)

		expected := []string{"create_user", "delete_user", "edit_user", "view_user", "ban_user"}
		actual := make([]string, len(commands))
		for i, cmd := range commands {
			actual[i] = cmd.Code
			assert.Equal(t, int64(1), cmd.RoleID)
			assert.Equal(t, int64(1), cmd.Role.ID)
			assert.Equal(t, "admin", cmd.Role.Code)
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("moderator role has specific commands", func(t *testing.T) {
		commands, err := repo.FindByRoleID(ctx, 2)
		require.NoError(t, err)
		require.Len(t, commands, 3)

		expected := []string{"edit_user", "view_user", "ban_user"}
		actual := make([]string, len(commands))
		for i, cmd := range commands {
			actual[i] = cmd.Code
			assert.Equal(t, int64(2), cmd.RoleID)
			assert.Equal(t, "moderator", cmd.Role.Code)
		}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("user role has only view command", func(t *testing.T) {
		commands, err := repo.FindByRoleID(ctx, 3)
		require.NoError(t, err)
		require.Len(t, commands, 1)

		assert.Equal(t, "view_user", commands[0].Code)
		assert.Equal(t, int64(3), commands[0].RoleID)
		assert.Equal(t, "user", commands[0].Role.Code)
	})

	t.Run("unknown role returns empty", func(t *testing.T) {
		commands, err := repo.FindByRoleID(ctx, 99)
		require.NoError(t, err)
		assert.Empty(t, commands)
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		commands, err := repo.FindByRoleID(ctx, 1)
		assert.ErrorIs(t, err, context.Canceled)
		assert.Nil(t, commands)
	})
}

func TestCommandRepository_EdgeCases(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB(t)
	repo := NewCommandRepository(db)

	t.Run("empty database", func(t *testing.T) {
		commands, err := repo.FindByRoleID(ctx, 1)
		require.NoError(t, err)
		assert.Empty(t, commands)
	})

	t.Run("role without commands", func(t *testing.T) {
		_, err := db.Exec(`INSERT INTO roles (id, code) VALUES (10, 'guest')`)
		require.NoError(t, err)

		commands, err := repo.FindByRoleID(ctx, 10)
		require.NoError(t, err)
		assert.Empty(t, commands)
	})

	t.Run("duplicate commands not returned", func(t *testing.T) {
		_, err := db.Exec(`INSERT OR IGNORE INTO role_commands VALUES (1, 101)`)
		require.NoError(t, err)

		commands, err := repo.FindByRoleID(ctx, 1)
		require.NoError(t, err)

		// Проверяем что команды не дублируются
		codes := make(map[string]bool)
		for _, cmd := range commands {
			if codes[cmd.Code] {
				t.Errorf("Duplicate command found: %s", cmd.Code)
			}
			codes[cmd.Code] = true
		}
	})
}
