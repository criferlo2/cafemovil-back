CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    phone VARCHAR(30) NOT NULL UNIQUE,
    guest BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    sort_order INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(180) NOT NULL,
    brand VARCHAR(120) NOT NULL,
    category_id VARCHAR(64) NOT NULL,
    presentation VARCHAR(120) NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    featured BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO categories (id, name, sort_order) VALUES
    ('coffee', 'Café', 1),
    ('accessories', 'Accesorios', 2),
    ('machines', 'Máquinas para hacer café', 3),
    ('manuals', 'Manuales', 4),
    ('books', 'Libros', 5)
ON DUPLICATE KEY UPDATE name = VALUES(name), sort_order = VALUES(sort_order);
