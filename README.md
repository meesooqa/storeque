### Migration

```shell
go run ./db/cmd/migration/main.go
```

### Run

```shell
go run ./tg/cmd/bot/main.go
```

### Add new command

1. Create a record in `commands` table.
2. Create records in `role_commands` table for roles.
3. Add handler file in `tg/handlers/commands/` directory.
4. Register handler in `commands.GetAll()` method in file `tg/handlers/commands/base.go`.
5. Add translation in `tg/locales/*.json` files.
