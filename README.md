# Simple CMS Admin V2

Welcome to the Simple CMS Admin Service V2. An open-source Content Management System based on the echo framework. As admin, we can use the features provided by this service in the form of management of articles and categories. By using this service we can insert, update, delete and get details of each item (article and category). This service has implemented clean architecture principles, a practical software architecture solution from Robert C. Martin (known as Uncle Bob).

## What's New?

The following are some of the updates in this latest version of the content management system:

- Standardize the response format.
- Refactor algorithm including handler, usecase, and repository layer.
- Add custom middlewares and validator.
- Update database schema.
- Minor bug fixes.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- Elasticsearch
- Redis

### Installation

- Clone the git repository:

```
$ git clone https://github.com/backendmagang/project-1.git
$ cd simple-cms-admin-v2
```

- Install Dependencies

```
$ go mod tidy
```

- Install Elasticsearch and Redis

```
$ docker pull docker.elastic.co/elasticsearch/elasticsearch:8.8.1
$ docker pull redis
```

- Adjust Configuration

```
$ cp example.config.json config.json
```

Then update config file as needed. I think you are quite easy for a smart person like you:)

### Running

```
$ go run cmd/main.go
```

or simply

```
$ make run
```

## API Reference

#### Get List of Articles

```http
  GET /admin/v2/article
```

| Parameter  | Type     | Description          |
| :--------- | :------- | :------------------- |
| `limit`    | `int`    |                      |
| `offset`   | `int`    |                      |
| `sort_by`  | `string` | can be every columns |
| `order_by` | `string` | "asc", "desc"        |

#### Get Article Details

```http
  GET /admin/v2/article/{id}
```

| Parameter | Type     | Description                             |
| :-------- | :------- | :-------------------------------------- |
| `id`      | `string` | **[Required]** UUID of article to fetch |

#### Insert Article

```http
  POST /admin/v2/article
```

| Parameter     | Type     | Description                                                  |
| :------------ | :------- | :----------------------------------------------------------- |
| `id`          | `string` | UUID will be generated automatically.                        |
| `title`       | `string` | **[Required]** Title of the article.                         |
| `slug`        | `string` | **[Required]** URL-friendly slug of the article.             |
| `htmlcontent` | `string` | **[Required]** HTML content of the article.                  |
| `category_id` | `[]int`  | **[Required]** IDs of the categories the article belongs to. |
| `metadata`    | `string` | **[Required]** Metadata of the article.                      |
| `created_at`  | `string` | Creation timestamp of the article.                           |
| `updated_at`  | `string` | Last update timestamp of the article.                        |

#### Update Article

```http
  PATCH /admin/v2/article/{id}
```

| Parameter     | Type     | Description                                   |
| :------------ | :------- | :-------------------------------------------- |
| `id`          | `string` | **[Required]** UUID of the article to update. |
| `title`       | `string` | Title of the article to update.               |
| `slug`        | `string` | URL-friendly slug of the article to update.   |
| `htmlcontent` | `string` | HTML content of the article to update.        |
| `category_id` | `[]int`  | IDs of the categories the article belongs to. |
| `metadata`    | `string` | Metadata of the article to update.            |
| `created_at`  | `string` | Creation timestamp of the article.            |
| `updated_at`  | `string` | Last update timestamp of the article.         |

### Delete Article

```http
  DELETE /admin/v2/article/{id}
```

| Parameter | Type     | Description                              |
| :-------- | :------- | :--------------------------------------- |
| `id`      | `string` | **[Required]** UUID of article to delete |

Everything you need, such as the postman collection and database migration script, can be found in the folder in `simple-cms-admin-v2/tools`.

### Unit Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/backendmagang/project-1/blob/master/LICENSE) file for details.
