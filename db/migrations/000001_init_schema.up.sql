CREATE TABLE "accounts" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "hashed_password" VARCHAR(255) NOT NULL,
  "avatar_url" VARCHAR(255) NOT NULL DEFAULT ('avatars/deafault_avatar.jpeg'),
  "is_admin" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(1023) NOT NULL,
  "author" VARCHAR(511) NOT NULL,
  "description" VARCHAR(2047) NOT NULL,
  "year" SMALLINT NOT NULL,
  "language" VARCHAR(127) NOT NULL,
  "publisher" VARCHAR(511) NOT NULL,
  "pages" SMALLINT NOT NULL,
  "cover_url" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(2047) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "book_category" (
  "book_id" INT NOT NULL,
  "category_id" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "bookmarks" (
  "id" SERIAL PRIMARY KEY,
  "book_id" INT NOT NULL,
  "type" INT NOT NULL,
  "created_by" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "rates" (
  "id" SERIAL PRIMARY KEY,
  "book_id" INT NOT NULL,
  "created_by" INT NOT NULL,
  "rate_value" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "comments" (
  "id" SERIAL PRIMARY KEY,
  "content" VARCHAR(1000) NOT NULL,
  "book_id" INT NOT NULL,
  "created_by" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX ON "accounts" ("email");

CREATE INDEX ON "accounts" ("is_admin");

CREATE INDEX ON "books" ("title");

CREATE INDEX ON "books" ("author");

CREATE INDEX ON "books" ("language");

CREATE INDEX ON "books" ("publisher");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "book_category" ("book_id");

CREATE INDEX ON "book_category" ("category_id");

CREATE UNIQUE INDEX ON "book_category" ("book_id", "category_id");

CREATE INDEX ON "bookmarks" ("book_id");

CREATE UNIQUE INDEX ON "bookmarks" ("book_id", "created_by");

CREATE INDEX ON "comments" ("book_id");

CREATE INDEX ON "comments" ("created_by");

ALTER TABLE "book_category" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_category" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "rates" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("created_by") REFERENCES "accounts" ("id");
