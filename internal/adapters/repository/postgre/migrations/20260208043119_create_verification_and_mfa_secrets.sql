-- Create "users_mfa_secrets" table
CREATE TABLE "users_mfa_secrets" (
  "user_id" uuid NOT NULL,
  "secret_key" character varying(255) NOT NULL,
  "backup_codes" text[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("user_id"),
  CONSTRAINT "fk_users_mfa_secrets_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "verifications" table
CREATE TABLE "verifications" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "type" "verification_type" NOT NULL,
  "token_hash" character varying(255) NOT NULL,
  "metadata" jsonb NULL,
  "is_used" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_verifications_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_verifications_user_id" to table: "verifications"
CREATE INDEX "idx_verifications_user_id" ON "verifications" ("user_id");
