-- Create "permissions" table
CREATE TABLE "permissions" (
  "id" bigserial NOT NULL,
  "slug" character varying(50) NOT NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_permissions_slug" UNIQUE ("slug")
);
-- Create "role" table
CREATE TABLE "role" (
  "id" bigserial NOT NULL,
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_role_name" UNIQUE ("name")
);
-- Create "role_permissions" table
CREATE TABLE "role_permissions" (
  "role_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  PRIMARY KEY ("role_id", "permission_id"),
  CONSTRAINT "fk_role_permissions_permissions" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_role_permissions_role" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "user_roles" table
CREATE TABLE "user_roles" (
  "user_id" uuid NOT NULL,
  "role_id" bigint NOT NULL,
  "assigned_at" timestamptz NOT NULL DEFAULT now(),
  "assigned_by" uuid NOT NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
