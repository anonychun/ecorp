-- +goose Up
-- +goose StatementBegin
CREATE TABLE admins (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email_address TEXT UNIQUE NOT NULL,
    password_digest TEXT NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE admin_sessions (
    id UUID PRIMARY KEY,
    admin_id UUID REFERENCES admins(id),
    token TEXT UNIQUE NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE admin_sessions;

DROP TABLE admins;
-- +goose StatementEnd
