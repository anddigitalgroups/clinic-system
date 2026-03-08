CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    role TEXT NOT NULL
        CHECK (role IN ('super_admin', 'admin', 'doctor', 'patient')),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
