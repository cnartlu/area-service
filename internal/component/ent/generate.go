package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/lock,sql/modifier,sql/execquery --template ./../../../pkg/ent/tmpl/ ./schema
