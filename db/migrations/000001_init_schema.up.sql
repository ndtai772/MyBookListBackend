CREATE TABLE "accounts" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "encoded_hash" VARCHAR(255) NOT NULL,
  "avatar_uri" VARCHAR(255) NOT NULL DEFAULT (`/asserts/images/deafault_avatar.jpeg`),
  "is_admin" boolean NOT NULL DEFAULT false,
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "author" VARCHAR(255) NOT NULL,
  "description" VARCHAR(10000) NOT NULL,
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(10000),
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "book_category" (
  "book_id" INT NOT NULL,
  "category_id" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "bookmarks" (
  "id" INT PRIMARY KEY,
  "book_id" INT NOT NULL,
  "created_by" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "rates" (
  "id" SERIAL PRIMARY KEY,
  "book_id" INT NOT NULL,
  "created_by" INT NOT NULL,
  "rate_value" INT NOT NULL,
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "comments" (
  "id" SERIAL PRIMARY KEY,
  "content" VARCHAR(1000) NOT NULL,
  "book_id" INT NOT NULL,
  "created_by" INT NOT NULL,
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "feedbacks" (
  "id" SERIAL PRIMARY KEY,
  "content" VARCHAR(1000) NOT NULL,
  "created_by" INT NOT NULL,
  "is_viewed" boolean NOT NULL DEFAULT false,
  "is_processing" boolean NOT NULL DEFAULT false,
  "is_resolved" boolean NOT NULL DEFAULT false,
  "message" VARCHAR(1000),
  "modified_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "book_category" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

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

CREATE INDEX ON "bookmarks" ("created_by");

CREATE UNIQUE INDEX ON "bookmarks" ("book_id", "created_by");

CREATE INDEX ON "comments" ("book_id");

CREATE INDEX ON "comments" ("created_by");

CREATE INDEX ON "feedbacks" ("created_by");
