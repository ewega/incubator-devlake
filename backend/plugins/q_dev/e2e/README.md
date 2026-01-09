# Q Developer Plugin E2E Tests

This directory contains end-to-end tests for the Q Developer plugin.

## Running the Tests

E2E tests require a MySQL database to run. You can set it up using the docker-compose file in the project root:

```bash
# From the project root directory
docker-compose -f docker-compose-dev.yml up mysql

# Set the E2E_DB_URL environment variable
export E2E_DB_URL="mysql://merico:merico@127.0.0.1:3306/lake_test?charset=utf8mb4&parseTime=True&loc=UTC"

# Create the test database
mysql -h 127.0.0.1 -P 3306 -u root -p -e "CREATE DATABASE IF NOT EXISTS lake_test; GRANT ALL ON lake_test.* TO 'merico'@'%';"

# Run the tests from the plugin directory
cd backend/plugins/q_dev
go test -v ./e2e/...
```

## Test Structure

The e2e tests follow the standard DevLake e2e pattern:

- `raw_tables/` - Contains raw test data (CSV files simulating S3 data)
- `snapshot_tables/` - Contains expected output data for verification
- `*_test.go` - Test files that verify the data transformation

## Test Coverage

- **TestQDevS3FileMeta** - Tests the S3 file metadata model
- **TestQDevUserData** - Tests the user data model and CSV parsing
