package auditlogs

/*
CREATE TABLE audit_logs (
    id BIGSERIAL PRIMARY KEY, -- Bigint for massive scale
    user_id UUID, -- Nullable because action might be by anonymous user (failed login)

    action VARCHAR(50) NOT NULL, -- e.g., 'LOGIN_SUCCESS', 'LOGIN_FAILED', 'PASSWORD_CHANGED', 'COUNTRY_UPDATED_ML'
    entity_type VARCHAR(50) NOT NULL, -- e.g., 'USER', 'SESSION', 'BILLING'
    entity_id UUID, -- The ID of the object being changed

    -- Context
    ip_address INET,
    user_agent TEXT,
    metadata JSONB DEFAULT '{}', -- Details: { "old_country": "US", "new_country": "CA", "confidence": 0.98 }

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Partitioning for scale (Enterprise Requirement)
-- In a real enterprise DB, you would partition audit logs by date (e.g., monthly)
-- CREATE TABLE audit_logs_2023_10 PARTITION OF audit_logs ...
*/
