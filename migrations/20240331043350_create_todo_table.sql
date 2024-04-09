-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE SCHEMA IF NOT EXISTS go_api;

CREATE TABLE IF NOT EXISTS "go_api"."users"
(
  "id" uuid not null,
  "name" varchar(255) not null,
  "email" varchar(255) not null,
  "created_at" timestamptz not null,
  "updated_at" timestamptz not null,
  primary key("id")
);

CREATE TABLE IF NOT EXISTS "go_api"."todos"
(
  "id" uuid not null,
  "todo_id" uuid not null unique,
  "title" varchar(255) not null,
  "description" text ,
  "is_deletable" boolean not null,
  "created_at" timestamptz not null,
  "updated_at" timestamptz not null,
  primary key("id", "todo_id"),
  FOREIGN KEY ("id") REFERENCES "go_api"."users"("id")
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "go_api"."tags"
(
  "id" uuid not null,
  "tag_id" BIGSERIAL unique,
  "name" varchar(32) not null,
  "created_at" timestamptz not null,
  "updated_at" timestamptz not null,
  primary key("id", "tag_id"),
  FOREIGN KEY ("id") REFERENCES "go_api"."users"("id")
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "go_api"."todo_tag"
(
  "todo_id" uuid NOT NULL,
  "tag_id" BIGINT NOT NULL,
  primary key("todo_id", "tag_id"),
  FOREIGN KEY ("todo_id") REFERENCES "go_api"."todos"("todo_id")
    ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY ("tag_id") REFERENCES "go_api"."tags"("tag_id")
    ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP SCHEMA IF EXISTS go_api CASCADE;
-- +goose StatementEnd
