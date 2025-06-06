### Migration

```shell
go run ./db/cmd/migration/main.go
```

### Run

```shell
go run ./tg/cmd/bot/main.go
```

### Add new command

1. Update `roleservice.Service.roleCommands`.
2. Add handler file in `tg/handlers/commands/` directory.
3. Register handler in `commands.GetAll()` method in file `tg/handlers/commands/base.go`.
4. Add translation in `tg/locales/*.json` files.
