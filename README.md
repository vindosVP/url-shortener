# URL-Shortener

## ⚡️ Description

Service to create shortened urls

## ⚙️ Installation

### ENV variables

| Name              | Description                | Default value   | Expected value                  | Requiered |
|:------------------|:---------------------------|:----------------|:--------------------------------|:---------:|
| STORAGE_TYPE      | data storage type          | `postgres`      | `inmemory`/`postgres`           |    ✔️     |
| SERVER_TYPE       | server type                | `grpc`          | `api`/`grpc`                    |    ✔️     |
| API_PORT          | API server port            | `8080`          | api server port                 |    ✔️     |
| GRPC_PORT         | GRPC server port           | `8081`          | grpc server port                |    ✔️     |
| DB_HOST           | postgres host              | `postgres`      | postgres host                   |    ✔️     |
| DB_PORT           | postgres port              | `5432`          | postgres port                   |    ✔️     |
| DB_NAME           | database name              | `url-shortener` | postgres db name                |    ✔️     |
| DB_USER           | postgres username          | `admin`         | postgres db admin username      |    ✔️     |
| DB_PWD            | postgres password          | `admin`         | postgres db admin password      |    ✔️     |
| DB_SSL_MODE       | Disable or enable SSL mode | `disable`       | `enable`/`disable`              |    ✔️     |
| DB_TIMEZONE       | Database timezone          | `Europe/Moscow` | timezone                        |    ✔️     |
| POSTGRES_USER     | Postgres admin username    | `admin`         | admin username to init postgres |    ✔️     |
| POSTGRES_PASSWORD | Postgres admin password    | `admin`         | admin password to init postgres |    ✔️     |
| POSTGRES_DB       | Postgres initial database  | `url-shortener` | initial database name           |    ✔️     |

If you need to change the default values, you can do it in [.env](./.env)

1. Clone the GitHub repo

```Shell
git clone https://github.com/vindosVP/url-shortener.git
```

### 🐳 Start in docker

```Shell
cd url-shortener
```

```Shell
docker-compose up
```

### 🖥️ Start on local machine

1. Set up env variables (example in [.env](./.env))

2. Execute following lines

```Shell
cd url-shortener
```

```Shell
make build
```

```Shell
./url-shortener
```

### 📄 Testing

1. Execute following lines
```Shell
cd url-shortener
```
```Shell
make test 
```

### 📖 Docs

API documentation: [swagger](/docs/swagger/swagger.yaml)

GRPC documentation: [proto file](src/internal/controller/grpcController/url-shortener.proto)