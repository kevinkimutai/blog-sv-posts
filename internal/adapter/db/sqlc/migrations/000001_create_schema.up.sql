CREATE TABLE movies (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "director" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);