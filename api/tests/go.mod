module tests

go 1.18

replace github.com/AbdulkarimOgaji/kkmoney/db => ../../db

require (
	github.com/AbdulkarimOgaji/kkmoney/db v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
