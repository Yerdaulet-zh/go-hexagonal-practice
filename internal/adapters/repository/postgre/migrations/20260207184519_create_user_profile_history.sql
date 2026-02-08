-- Create "user_profile_history" table
CREATE TABLE "user_profile_history" (
  "history_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "first_name" character varying(100) NOT NULL,
  "last_name" character varying(100) NULL,
  "country_code" character varying(2) NULL,
  "country_source" character varying(50) NULL,
  "avatar_url" text NULL,
  "changed_at" timestamptz NOT NULL DEFAULT now(),
  "operation" character varying(10) NOT NULL,
  PRIMARY KEY ("history_id"),
  CONSTRAINT "fk_user_profile_history_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
