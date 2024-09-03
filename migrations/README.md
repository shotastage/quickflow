# QuickFlow System Migrations

This directory is used to manage database migration files for our Go-based web system.

## Overview

Migrations are a system for tracking and applying changes to the database schema. The files in this directory record the evolution of the database structure and help maintain consistency across the entire team.

## Directory Structure

```
migrations/
├── YYYYMMDDHHMMSS_create_users_table.go
├── YYYYMMDDHHMMSS_add_email_to_users.go
└── ...
```

Each migration file is created with a timestamp and a name indicating the change.

## Creating Migration Files

To create a new migration, use the following command:

```
go run cmd/migrate/main.go create <migration_name>
```

## Running Migrations

To apply migrations, use the following command:

```
go run cmd/migrate/main.go up
```

To rollback migrations:

```
go run cmd/migrate/main.go down
```

## Important Notes

- Once created, migration files should not be modified.
- If new changes are needed, create a new migration file.
- Always backup your database before running migrations.

This README.md file serves as a guide for the development team to understand and properly manage the migration process.
