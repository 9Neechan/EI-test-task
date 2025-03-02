CREATE TABLE "services" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" DECIMAL(10,2) NOT NULL DEFAULT 0.00,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stats" (
  "user_id" bigint NOT NULL,
  "service_id" bigint NOT NULL,
  "count" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("user_id", "service_id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL,
  FOREIGN KEY ("service_id") REFERENCES "services" ("id") ON DELETE SET NULL
);

CREATE UNIQUE INDEX idx_stats_user_service ON "stats" ("user_id", "service_id");
CREATE INDEX idx_stats_user ON "stats" ("user_id");
CREATE INDEX idx_stats_service ON "stats" ("service_id");

