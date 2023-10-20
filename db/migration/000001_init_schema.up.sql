CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" numeric(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" numeric(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "source_id" bigint NOT NULL,
  "dest_id" bigint NOT NULL,
  "amount" numeric(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("source_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("dest_id") REFERENCES "accounts" ("id");

CREATE INDEX "accounts_owner_idx" ON "accounts" ("owner");

CREATE INDEX "entries_account_id_idx" ON "entries" ("account_id");

CREATE INDEX "transfers_source_id_idx" ON "transfers" ("source_id");

CREATE INDEX "transfers_dest_id_idx" ON "transfers" ("dest_id");

CREATE INDEX "source_dest_idx" ON "transfers" ("source_id", "dest_id");
