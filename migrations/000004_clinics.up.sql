CREATE TABLE clinics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR NOT NULL,
    address TEXT,
    phone VARCHAR,
    email VARCHAR,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_clinics_is_active ON clinics(is_active);
