createtable:
	sqlite3 test.db < script_create_table.sql
mockdata:
	sqlite3 test.db < script_insert_mock.sql
demo: createtable mockdata
	go run main.go
