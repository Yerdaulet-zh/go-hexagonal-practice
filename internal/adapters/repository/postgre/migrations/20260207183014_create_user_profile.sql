-- Create "user_profiles" table
CREATE TABLE "user_profiles" (
  "user_id" uuid NOT NULL,
  "first_name" character varying(100) NOT NULL,
  "last_name" character varying(100) NULL,
  "country_code" character varying(2) NULL,
  "country_source" character varying(50) NULL,
  "avatar_url" text NULL,
  "preferences" jsonb NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("user_id"),
  CONSTRAINT "fk_user_profiles_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
