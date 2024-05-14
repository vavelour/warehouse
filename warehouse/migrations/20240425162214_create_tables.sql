-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouse (
    id SERIAL PRIMARY KEY,
    name TEXT,
    accessibility BOOLEAN
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT,
    size INTEGER CHECK (size >= 0),
    unique_code INTEGER UNIQUE CHECK (unique_code >= 0)
);

CREATE TABLE item_warehouse_quantity (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES items(id),
    warehouse_id INTEGER REFERENCES warehouse(id),
    quantity INTEGER CHECK (quantity >= 0)
);

CREATE TABLE reserved_items (
    id SERIAL PRIMARY KEY,
    item_warehouse_id INTEGER REFERENCES item_warehouse_quantity(id),
    quantity INTEGER CHECK (quantity >= 1)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reserved_items;

DROP TABLE IF EXISTS item_warehouse_quantity;

DROP TABLE IF EXISTS items;

DROP TABLE IF EXISTS warehouse;
-- +goose StatementEnd
