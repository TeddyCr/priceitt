CREATE TABLE IF NOT EXISTS DATABASE_MIGRATION_LOGS (
    installed_rank INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    version VARCHAR(256),
    query VARCHAR(10000),
    checksum VARCHAR(256),
    installed_on DATETIME DEFAULT (CURRENT_TIMESTAMP),
    execution_time INTEGER
);