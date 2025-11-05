CREATE TABLE users (
    id uuid PRIMARY KEY  default gen_random_uuid(),
    full_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(15),
    role SMALLINT,
    is_admin BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT FALSE,
    password_hash TEXT NOT NULL,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_profiles (
    id  uuid PRIMARY KEY default gen_random_uuid(),
    user_id uuid UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    profile_picture VARCHAR(255),
    cover_photo VARCHAR(255),
    address_line1 TEXT,
    address_line2 TEXT,
    country VARCHAR(100),
    state VARCHAR(100),
    city VARCHAR(100),
    pin_code VARCHAR(20),
    longitude VARCHAR(50),
    latitude VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
