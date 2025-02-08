ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS order_items_order_id_fkey;

DROP TRIGGER IF EXISTS generate_order_id_trigger ON orders;
DROP TRIGGER IF EXISTS generate_order_item_id_trigger ON order_items;

DROP FUNCTION IF EXISTS generate_order_id();
DROP FUNCTION IF EXISTS generate_order_item_id();

DROP SEQUENCE IF EXISTS order_id_seq;
DROP SEQUENCE IF EXISTS order_item_id_seq;


DROP TABLE IF EXISTS "orders";
DROP TABLE IF EXISTS "order_items";
