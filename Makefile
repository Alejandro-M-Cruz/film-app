fresh:
	if [ -f ./db/database.sqlite3 ]; then rm ./db/database.sqlite3; fi
	sqlite3 db/database.sqlite3 < db/migrations/2025_02_24_220600_init.sql
	sqlite3 db/database.sqlite3 < db/seeders/db_seeder.sql
