export PGUSER=postgres
export PGPASSWORD=postgres
export PGHOST=localhost
export PGPORT=5432
export PGDATABASE=postgres
export PGAPPNAME=goofer
export PGSCHEMA=public
go generate rester/*
go run rester/*
