# recoshelf-api

REST API for ReCoShelf, a app to track and manage your physical music collection.

# Tech Stack

- Golang
- Fiber
- MySQL
- OpenAPI

# Features

Get user's releases
- `GET /me/releases`
- return all releases for a authenticated user

Add a release to user collection
- `POST /me/releases`
- try to get the release by `source` and `sourceReleaseID` first
- if the release has already exist, add the relationship between user and release
- if the release doesn't exist, create the release, then add the relationship between user and release

Delete a release from user collection
- `DELETE /me/releases/{releaseID}`
- delete the relationship between user and release

Batch delete releases from user collection
- `POST /me/releases/batch-delete`
- batch delete releases from user collection

Currently in the mock-auth mode (header-based)
- use header `X-User-Id` to set the user

# API docs

- OpenAPI spec: [docs/openapi.yaml](docs/openapi.yaml)

# DB schema

- Schema file: [db/schema.sql](db/schema.sql)

## users

| Column | Type | Notes |
| --- | --- | --- |
| id | INT unsigned | Primary key, auto-increment. |
| account | VARCHAR(255) | Required. |
| created_at | DATETIME | Default `CURRENT_TIMESTAMP`. |
| updated_at | DATETIME | Default `CURRENT_TIMESTAMP`, auto-updates on row change. |

Constraints:
- Primary key: `id`

## releases

- Store the required fields of release data from external source.
- Store external identifiers as (source, source_release_id) to support multiple providers in the future.

| Column | Type | Notes |
| --- | --- | --- |
| id | INT unsigned | Primary key, auto-increment. |
| source | VARCHAR(32) | Required. The source of release data. e.g. discogs |
| source_release_id | BIGINT | Required. The release id from source. |
| title | VARCHAR(200) | Required. Release title. |
| country | VARCHAR(5) | Required. Release country. |
| genres | JSON | Required. Genres of release. |
| release_year | SMALLINT unsigned | Required. Release year. |
| tracklist | JSON | Required. Track list of release. |
| images | VARCHAR(512) | Optional. Release cover images. |
| barcode | VARCHAR(20) | Required. Release barcode. |
| fetched_at | DATETIME | Required. The time that fetch release data from source. |
| created_at | DATETIME | Default `CURRENT_TIMESTAMP`. |
| updated_at | DATETIME | Default `CURRENT_TIMESTAMP`, auto-updates on row change. |

Constraints:
- Primary key: `id`
- Unique: (`source`, `source_release_id`)

## releases_users

| Column | Type | Notes |
| --- | --- | --- |
| id | INT unsigned | Primary key, auto-increment. |
| user_id | INT unsigned | Required. |
| release_id | INT unsigned | Required. |
| created_at | DATETIME | Default `CURRENT_TIMESTAMP`. |
| updated_at | DATETIME | Default `CURRENT_TIMESTAMP`, auto-updates on row change. |

Constraints:
- Primary key: `id`
- Unique: (`user_id`, `release_id`)

# Installation

Prerequisites:
- Go (recent version)
- MySQL

1) Create database (adjust name if you want):
```sh
mysql -u root -p -e "CREATE DATABASE recoshelf CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

2) Run schema:
```sh
mysql -u root -p recoshelf < db/schema.sql
```

3) Create `.env`:
```sh
cp .env.local .env
```

```env
MYSQL_DSN=YOUR_USER:YOUR_PASS@tcp(127.0.0.1:3306)/recoshelf?parseTime=true
```

4) Add demo data:
```sh
mysql -u root -p recoshelf < db/seed.sql
```

5) Run the API:
```sh
go run main.go
```

6) Test the API:
- Open `tests/api/me.http` with a REST client (VS Code REST Client, IntelliJ HTTP client) and run the requests.
