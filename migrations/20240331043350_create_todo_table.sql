-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE SCHEMA IF NOT EXISTS go_api;

CREATE TABLE IF NOT EXISTS "go_api"."users"
(
  "id" uuid NOT NULL,
  "firebase_uid" varchar(255) NOT NULL UNIQUE,
  "name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "go_api"."todos"
(
  "id" uuid NOT NULL,
  "todo_id" uuid NOT NULL unique,
  "title" varchar(255) NOT NULL,
  "description" text ,
  "is_deletable" boolean NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY("id", "todo_id"),
  FOREIGN KEY ("id") REFERENCES "go_api"."users"("id")
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "go_api"."tags"
(
  "id" uuid NOT NULL,
  "tag_id" BIGSERIAL unique,
  "name" varchar(32) NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY("id", "tag_id"),
  FOREIGN KEY ("id") REFERENCES "go_api"."users"("id")
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "go_api"."todo_tags"
(
  "todo_id" uuid NOT NULL,
  "tag_id" BIGINT NOT NULL,
  PRIMARY KEY("todo_id", "tag_id"),
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
