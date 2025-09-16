package migrations

import "embed"

//go:embed *.sql *.go
var MigrationsFs embed.FS
