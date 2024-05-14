-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION check_reserved_items_insert()
RETURNS TRIGGER AS $$
DECLARE
    warehouse_accessible BOOLEAN;
    item_quantity INTEGER;
BEGIN

    SELECT w.accessibility
    INTO warehouse_accessible
    FROM warehouse w
    JOIN item_warehouse_quantity iwq ON iwq.warehouse_id = w.id
    WHERE iwq.id = NEW.item_warehouse_id;

    IF NOT warehouse_accessible THEN
        RAISE EXCEPTION 'Warehouse with id % is not accessible', (SELECT warehouse_id FROM item_warehouse_quantity WHERE id = NEW.item_warehouse_id);
    END IF;

    SELECT quantity
    INTO item_quantity
    FROM item_warehouse_quantity
    WHERE id = NEW.item_warehouse_id;

    IF item_quantity IS NULL OR item_quantity < NEW.quantity THEN
        RAISE EXCEPTION 'Insufficient quantity of item with item_warehouse_id % in warehouse', NEW.item_warehouse_id;
    END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_reserved_items
    BEFORE INSERT ON reserved_items
    FOR EACH ROW
    EXECUTE FUNCTION check_reserved_items_insert();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS before_insert_reserved_items ON reserved_items;
DROP FUNCTION IF EXISTS check_reserved_items_insert();
-- +goose StatementEnd
