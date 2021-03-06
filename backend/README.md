# DB create tables

## Users

```sql
CREATE TYPE role_t AS ENUM('user', 'editor', 'admin');

CREATE TABLE users (
    id             integer GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    username       varchar(255),
    email          varchar(255) NOT NULL,
    password       varchar(255),
    fullName       varchar(255),
    role           role_t NOT NULL DEFAULT 'user',
    blocked        boolean NOT NULL DEFAULT FALSE,
    avatar         text
);
```

## Posts

```sql
CREATE TABLE posts (
    id             integer GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title          varchar(255) NOT NULL,
    summary        varchar(255) NOT NULL,
    body           text DEFAULT '',
    date_created   timestamp NOT NULL DEFAULT NOW(),
    date_updated   timestamp NOT NULL DEFAULT NOW(),
    author_id      integer,
    published      boolean NOT NULL DEFAULT FALSE,
    image          text DEFAULT '',
    color          varchar(50) DEFAULT '#26c6da',
    views          integer NOT NULL DEFAULT 0,
    CONSTRAINT fk_user
      FOREIGN KEY(author_id) REFERENCES users(id)
);
```

# Compile

```bash
go build -o backend.o
```

# Systemd config

```bash
[Unit]
Description=blog api

[Service]
User=kuder
Group=www-data
EnvironmentFile=/home/kuder/.blog_env
ExecStart=/home/kuder/myblog/backend/backend.o
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```
