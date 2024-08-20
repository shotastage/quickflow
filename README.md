# QuickFlow
QuickFlow is a high-performance and flexible content management system (CMS) and API platform developed in Go. It leverages the Echo web framework and GORM, using a PostgreSQL database for its backend. A key feature of this application is that it allows users to freely define their own data schemas and easily build customized APIs. While offering functionalities similar to MicroCMS and Hasura, it serves as an efficient web service development platform that takes advantage of Go's speed and simplicity.

## Important Notice
This project is closed-source and is subject to an End-User License Agreement (EULA). Please make sure to review the EULA before using this software.

## Supported Operating Systems
QuickFlow is designed to run on the following operating systems:
- macOS
- Ubuntu

Please note that Windows and Windows Server are currently not supported.

## Prerequisites
- Go (version 1.16 or later)
- PostgreSQL
- migrate (Go's SQL migration tool)

## Setup
1. Clone this repository (access rights required).
2. Create a `.env` file and set the following environment variable:
   ```
   DATABASE_URL=postgres://username:password@localhost/database_name
   ```
3. Install migrate:
   ```
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```
4. Set up the database and run migrations:
   ```
   migrate -database ${DATABASE_URL} -path db/migrations up
   ```
5. Build the application:
   ```
   go build -o quickflow
   ```

## Running the Application
Execute the built binary:
```
./quickflow
```

## Configuration
Application configurations are managed in YAML files located in the `config` directory. Modify the settings for each environment as needed.

## Testing
To run unit and integration tests:
```
go test ./...
```

## License
This project is closed-source and is provided under the terms of the accompanying EULA. Please refer to the EULA file for details.

## Support
If you encounter any issues or have questions, please contact your internal support representative.
