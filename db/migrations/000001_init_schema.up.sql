CREATE TYPE "account_roles" AS ENUM (
  'admin',
  'user'
);

CREATE TYPE "feedback_status" AS ENUM (
  'processing',
  'resolved',
  'not_resolved'
);

CREATE TABLE "accounts" (
  "id" BIGINT PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "avatar_uri" VARCHAR(255),
  "role" account_roles NOT NULL DEFAULT 'user',
  "modified_at" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "author" VARCHAR(255) NOT NULL,
  "description" VARCHAR(10000) NOT NULL,
  "modified_at" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(10000),
  "modified_at" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "book_category" (
  "book_id" INT NOT NULL,
  "category_id" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "bookmarks" (
  "id" BIGINT PRIMARY KEY,
  "book_id" INT NOT NULL,
  "account_id" BIGINT NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "rates" (
  "id" BIGINT PRIMARY KEY,
  "book_id" INT NOT NULL,
  "account_id" BIGINT NOT NULL,
  "modified_at" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "comments" (
  "id" BIGINT PRIMARY KEY,
  "content" VARCHAR(1000) NOT NULL,
  "book_id" INT NOT NULL,
  "created_by" BIGINT NOT NULL,
  "modified_at" TIMESTAMP,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "feedbacks" (
  "id" BIGINT PRIMARY KEY,
  "content" VARCHAR(1000) NOT NULL,
  "created_by" BIGINT NOT NULL,
  "message" VARCHAR(1000),
  "status" feedback_status NOT NULL,
  "modified_at" timestamp,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "book_category" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

ALTER TABLE "feedbacks" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "accounts" ("email");

CREATE INDEX ON "books" ("author");

CREATE INDEX ON "books" ("modified_at");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "book_category" ("book_id");

CREATE INDEX ON "book_category" ("category_id");

CREATE UNIQUE INDEX ON "book_category" ("book_id", "category_id");

CREATE INDEX ON "bookmarks" ("book_id");

CREATE INDEX ON "bookmarks" ("account_id");

CREATE UNIQUE INDEX ON "bookmarks" ("book_id", "account_id");

CREATE INDEX ON "comments" ("book_id");

CREATE INDEX ON "comments" ("created_by");
