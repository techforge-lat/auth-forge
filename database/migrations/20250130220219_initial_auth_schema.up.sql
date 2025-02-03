-----------------------------
-- 1. Enable Extensions --
-----------------------------
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

---------------------------------
-- 2. Create Enum Types --
---------------------------------
CREATE TYPE user_origin AS ENUM ('STAFF', 'CUSTOMER', 'SYSTEM');
CREATE TYPE mfa_method_type AS ENUM ('TOTP', 'SMS', 'EMAIL', 'WEBAUTHN');

---------------------------------
-- 3. Core Tables --
---------------------------------
CREATE TABLE apps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(255) NOT NULL UNIQUE, -- Unique across all apps
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    app_id UUID NOT NULL REFERENCES apps(id),
    code VARCHAR(255) NOT NULL UNIQUE, -- Unique across all tenants
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

---------------------------------
-- 4. Users & Auth Tables --
---------------------------------
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origin user_origin NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE app_users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    app_id UUID NOT NULL REFERENCES apps(id),
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE(app_id, user_id)
);

CREATE TABLE email_credentials (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255),
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE social_providers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(255) NOT NULL UNIQUE, -- Unique across all social providers
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE social_credentials (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    social_provider_id UUID NOT NULL REFERENCES social_providers(id),
    provider_user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE(user_id, social_provider_id)
);

---------------------------------
-- 5. Permission System --
---------------------------------
CREATE TABLE resources (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    app_id UUID NOT NULL REFERENCES apps(id),
    code VARCHAR(100) NOT NULL UNIQUE, -- Unique across all resources
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE resource_actions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    resource_id UUID NOT NULL REFERENCES resources(id),
    code VARCHAR(100) NOT NULL UNIQUE, -- Unique across all resource actions
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    app_id UUID NOT NULL REFERENCES apps(id),
    code VARCHAR(100) NOT NULL UNIQUE, -- Unique across all roles
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE role_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL REFERENCES roles(id),
    resource_action_id UUID NOT NULL REFERENCES resource_actions(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE(role_id, resource_action_id)
);

CREATE TABLE user_roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    app_user_id UUID NOT NULL REFERENCES app_users(id),
    role_id UUID NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE(app_user_id, role_id)
);

---------------------------------
-- 6. Security Tables --
---------------------------------
CREATE TABLE mfa_methods (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    method mfa_method_type NOT NULL,
    secret TEXT,
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    app_id UUID NOT NULL REFERENCES apps(id),
    device_fingerprint VARCHAR(64),
    refresh_token_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

---------------------------------
-- 7. Timestamp Triggers --
---------------------------------
-- CREATE OR REPLACE FUNCTION update_modified_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = NOW();
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- -- Create triggers for all tables
-- DO $$
-- DECLARE
--     tbl text;
-- BEGIN
--     FOR tbl IN
--         SELECT table_name
--         FROM information_schema.tables
--         WHERE table_schema = 'public'
--         AND table_type = 'BASE TABLE'
--     LOOP
--         EXECUTE format('CREATE TRIGGER update_%s_modtime
--             BEFORE UPDATE ON %I
--             FOR EACH ROW EXECUTE FUNCTION update_modified_column()',
--             tbl, tbl);
--     END LOOP;
-- END$$;

---------------------------------
-- 8. Indexes --
---------------------------------
CREATE INDEX idx_apps_code ON apps(code);
CREATE INDEX idx_tenants_code ON tenants(code);
CREATE INDEX idx_resources_code ON resources(code);
CREATE INDEX idx_roles_code ON roles(code);
CREATE INDEX idx_sessions_user_app ON sessions(user_id, app_id);
