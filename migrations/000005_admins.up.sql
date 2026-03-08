CREATE TABLE admins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL UNIQUE
        REFERENCES users(id) ON DELETE CASCADE,

    clinic_id UUID NOT NULL
        REFERENCES clinics(id) ON DELETE CASCADE,

    name VARCHAR,
    phone VARCHAR,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_admins_clinic_id ON admins(clinic_id);
