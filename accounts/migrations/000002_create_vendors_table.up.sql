CREATE TABLE vendors (
    id uuid PRIMARY KEY  default gen_random_uuid(),
    user_id uuid UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    vendor_name TEXT,
    vendor_license TEXT,
    is_approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
