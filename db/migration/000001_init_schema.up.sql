CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "price" float NOT NULL,
  "stock" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "book_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "quantity" integer NOT NULL,
  "total_price" float NOT NULL,
  "status"   boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
