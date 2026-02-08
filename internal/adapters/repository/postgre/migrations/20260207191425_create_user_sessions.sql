-- Create "user_sessions" table
CREATE TABLE "user_sessions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "refresh_token_hash" character varying(255) NOT NULL,
  "ip_address" inet NOT NULL,
  "user_agent" text NULL,
  "device" character varying(255) NULL,
  "geo_location" jsonb NULL,
  "is_revoked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "last_active_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_user_sessions_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_sessions_token" to table: "user_sessions"
CREATE INDEX "idx_sessions_token" ON "user_sessions" ("refresh_token_hash");
-- Create index "idx_sessions_user_active" to table: "user_sessions"
CREATE INDEX "idx_sessions_user_active" ON "user_sessions" ("user_id", "is_revoked") WHERE (is_revoked = false);
-- Create index "idx_user_sessions_expires_at" to table: "user_sessions"
CREATE INDEX "idx_user_sessions_expires_at" ON "user_sessions" ("expires_at");
