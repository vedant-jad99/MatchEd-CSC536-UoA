#!/bin/bash

DB_NAME="match_db"
BACKUP_FILE="backup_$(date +%Y%m%d_%H%M%S).sql"

# SQLite
if [ -f "$DB_NAME.sqlite" ]; then
  if [ "$1" = "backup" ]; then
    cp "$DB_NAME.sqlite" "$DB_NAME.backup.sqlite"
    echo "✅ SQLite backup created: $DB_NAME.backup.sqlite"
  elif [ "$1" = "restore" ]; then
    cp "$DB_NAME.backup.sqlite" "$DB_NAME.sqlite"
    echo "✅ SQLite database restored from backup."
  else
    echo "Usage: ./db_testing.sh [backup|restore]"
  fi
  exit 0
fi

# PostgreSQL/MySQL
case "$1" in
  "backup")
    # PostgreSQL
    if command -v pg_dump &> /dev/null; then
      pg_dump -U match_user -d "$DB_NAME" -f "$BACKUP_FILE"
    # MySQL
    elif command -v mysqldump &> /dev/null; then
      mysqldump -u match_user -p "$DB_NAME" > "$BACKUP_FILE"
    fi
    echo "✅ Backup created: $BACKUP_FILE"
    ;;
  "restore")
    # PostgreSQL
    if command -v psql &> /dev/null; then
      psql -U match_user -d "$DB_NAME" -f "$BACKUP_FILE"
    # MySQL
    elif command -v mysql &> /dev/null; then
      mysql -u match_user -p "$DB_NAME" < "$BACKUP_FILE"
    fi
    echo "✅ Database restored from $BACKUP_FILE"
    ;;
  *)
    echo "Usage: ./db_testing.sh [backup|restore]"
    ;;
esac