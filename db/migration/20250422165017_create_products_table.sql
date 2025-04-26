-- +goose Up

CREATE TABLE categories (
 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
 name TEXT NOT NULL UNIQUE -- ex: Telefon, Laptop, Tableta, Casti
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id UUID NOT NULL REFERENCES categories(id),    
    device_type TEXT NOT NULL, -- Telefon, Laptop, Tableta, Casti
    model TEXT NOT NULL,
    color TEXT,
    storage TEXT, -- gen 128GB, 512GB
    battery_health TEXT, -- procent sau text ("Bună", "Foarte Bună")
    processor TEXT, -- doar pentru MacBook
    ram TEXT,       -- doar pentru MacBook
    description TEXT, -- text liber despre device

    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down

DROP TABLE products;
DROP TABLE categories;
