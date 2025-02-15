CREATE TABLE "orders" (
  "order_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "total_price" DECIMAL(10,2) NOT NULL,
  "status" VARCHAR NOT NULL,
  "shipping_address_id" UUID NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "order_items" (
  "oi_id" UUID PRIMARY KEY NOT NULL,
  "order_id" UUID NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "pv_id" VARCHAR(12) NOT NULL,
  "quantity" INT NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "total" DECIMAL(10,2) NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);
ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");