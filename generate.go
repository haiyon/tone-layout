package generate

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target ./internal/data/ent ./internal/data/schema
