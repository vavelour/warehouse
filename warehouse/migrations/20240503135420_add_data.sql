-- +goose Up
-- +goose StatementBegin
INSERT INTO warehouse (name, accessibility)
VALUES
    ('Warehouse A', true),
    ('Warehouse B', true),
    ('Warehouse C', false);

INSERT INTO items (name, size, unique_code)
VALUES
    ('Item 1', 5, 123),
    ('Item 2', 8, 456),
    ('Item 3', 3, 789),
    ('Item 4', 2, 234),
    ('Item 5', 7, 567),
    ('Item 6', 4, 890),
    ('Item 7', 6, 321),
    ('Item 8', 1, 654),
    ('Item 9', 9, 987),
    ('Item 10', 3, 543),
    ('Item 11', 5, 876);

INSERT INTO item_warehouse_quantity (item_id, warehouse_id, quantity)
VALUES
    (1, 1, 100),  -- Item 1, Warehouse A, 100 items
    (2, 1, 50),   -- Item 2, Warehouse A, 50 items
    (3, 1, 200),  -- Item 3, Warehouse A, 200 items
    (4, 2, 75),   -- Item 4, Warehouse B, 75 items
    (5, 2, 120),  -- Item 5, Warehouse B, 120 items
    (6, 2, 30),   -- Item 6, Warehouse B, 30 items
    (7, 3, 80),   -- Item 7, Warehouse C, 80 items
    (8, 3, 150),  -- Item 8, Warehouse C, 150 items
    (9, 3, 90),   -- Item 9, Warehouse C, 90 items
    (10, 1, 60),  -- Item 10, Warehouse A, 60 items
    (11, 2, 110); -- Item 11, Warehouse B, 110 items
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM reserved_items;

DELETE FROM item_warehouse_quantity;

DELETE FROM items;

DELETE FROM warehouse;
-- +goose StatementEnd
