DO $$ BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'verification_type') THEN 
        CREATE TYPE verification_type AS ENUM ('email_activation', 'password_reset', 'mfa_challenge', 'device_verification');
    END IF;
END $$;
