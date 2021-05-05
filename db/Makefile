SQLITE=sqlite3
SCHEMAFILE=schema.sql
DBFILE=sensors.db
SEEDFILE=seed.sql

seed:
	cat $(SEEDFILE) | $(SQLITE) $(DBFILE)

db:
	rm -f $(DBFILE)
	cat $(SCHEMAFILE) | $(SQLITE) $(DBFILE)

schema:
	$(SQLITE) $(DBFILE) ".schema --indent" | grep -v sqlite_sequence  > $(SCHEMAFILE)

dump:
	$(SQLITE) $(DBFILE) ".dump" | grep "^INSERT INTO" | grep -v sqlite_sequence  > $(SEEDFILE)
