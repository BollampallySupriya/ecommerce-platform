# ecommerce-platform
A Basic Ecommerce Platform to check products and add them to cart for placing order.


migration for tables .up file

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS order (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "customer" INT NOT NULL,
    "name" varchar NOT NULL,
    "price" FLOAT NOT NULL,
    "lineItems" []INT NOT NULL,
    "deliveryAddress" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


.down file

DROP TABLE IF EXISTS order;