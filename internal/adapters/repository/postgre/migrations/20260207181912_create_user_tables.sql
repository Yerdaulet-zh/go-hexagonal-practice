-- Create "user" table
CREATE TABLE "user" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "email" character varying(100) NOT NULL,
  "user_status" "user_status" NOT NULL DEFAULT 'pending_verification',
  "is_mfa_enabled" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_user_email" UNIQUE ("email")
);
-- Create index "idx_user_deleted_at" to table: "user"
CREATE INDEX "idx_user_deleted_at" ON "user" ("deleted_at");
-- Create index "idx_user_email" to table: "user"
CREATE INDEX "idx_user_email" ON "user" ("email");
-- Create "user_credentials" table
CREATE TABLE "user_credentials" (
  "user_id" uuid NOT NULL,
  "password_hash" character varying(255) NOT NULL,
  "password_salt" character varying(255) NULL,
  "last_password_change_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("user_id"),
  CONSTRAINT "fk_user_credentials_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
