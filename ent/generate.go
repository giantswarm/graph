package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --storage gremlin --idtype int --template ../gen/template/dialect/gremlin/create.tmpl,../gen/template/dialect/gremlin/update.tmpl ./schema
