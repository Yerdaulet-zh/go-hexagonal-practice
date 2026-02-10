ALTER TABLE user_profiles RENAME TO user_profile;

-- Rename the Foreign Key (GORM/Atlas will want this to match the new table name)
ALTER TABLE "user_profile" RENAME CONSTRAINT "fk_user_profiles_user" TO "fk_user_profile_user";

-- DROP TRIGGER IF EXISTS trigger_log_profile_changes ON user_profiles;

-- CREATE TRIGGER trigger_log_profile_changes
-- AFTER UPDATE ON user_profile
-- FOR EACH ROW
-- WHEN (OLD.* IS DISTINCT FROM NEW.*)
-- EXECUTE PROCEDURE log_profile_changes();
