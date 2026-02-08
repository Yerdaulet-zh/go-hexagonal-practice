-- Create "oauth_identities" table
CREATE TABLE "oauth_identities" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "provider" character varying(50) NOT NULL,
  "provider_id" character varying(255) NOT NULL,
  "profile_data" jsonb NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_oauth_identities_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "idx_user_provider" to table: "oauth_identities"
CREATE UNIQUE INDEX "idx_user_provider" ON "oauth_identities" ("user_id", "provider");
-- Create index "idx_user_provider_id" to table: "oauth_identities"
CREATE UNIQUE INDEX "idx_user_provider_id" ON "oauth_identities" ("provider", "provider_id");
