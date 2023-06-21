CREATE TABLE "users" (
    "id" varchar(255) PRIMARY KEY,
    "name" varchar not null,
    "birthday" date not null,
    "email" varchar not null unique,
    "password" varchar not null,
    "role" varchar not null,
    created_at timestamptz not null DEFAULT (now()),
    "updated_at" timestamptz not null DEFAULT (now())
);