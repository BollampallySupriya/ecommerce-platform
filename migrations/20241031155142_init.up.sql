-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    customer_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    price FLOAT NOT NULL,
    line_items INTEGER[] NOT NULL,
    delivery_address TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
