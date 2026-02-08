package verification

/*
CREATE TYPE verification_type AS ENUM ('email_activation', 'password_reset', 'mfa_challenge', 'device_verification');

CREATE TABLE verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type verification_type NOT NULL,
    token_hash VARCHAR(255) NOT NULL, -- Hashed short code or link token
    metadata JSONB, -- Store context (e.g., "triggered_by_ml_risk_engine")

    is_used BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE user_mfa_secrets (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    secret_key VARCHAR(255) NOT NULL, -- Encrypted TOTP secret
    backup_codes VARCHAR[] NOT NULL, -- Array of hashed backup codes
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
*/
