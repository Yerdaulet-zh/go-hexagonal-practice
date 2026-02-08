-- Trigger to auto-fill history
CREATE OR REPLACE FUNCTION log_profile_changes()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'UPDATE') THEN
        INSERT INTO user_profile_history (user_id, first_name, last_name, country_code, country_source, avatar_url, changed_at, operation)
        VALUES (OLD.user_id, OLD.first_name, OLD.last_name, OLD.country_code, OLD.country_source, OLD.avatar_url, NOW(), 'UPDATE');
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_log_profile_changes
AFTER UPDATE ON user_profiles
FOR EACH ROW
WHEN (OLD.* IS DISTINCT FROM NEW.*) -- Only log if data actually changed
EXECUTE PROCEDURE log_profile_changes();
