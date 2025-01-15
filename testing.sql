BEGIN;

-- Add a product
INSERT INTO products (name, price, quantity) VALUES ('New Product', 50, 100);

-- Add to cart
INSERT INTO cart (product_id, quantity) VALUES (1, 2);

-- Place an order
INSERT INTO orders (user_id, status) VALUES (1, 'Pending') RETURNING id;

-- Move cart items to order items
INSERT INTO order_items (order_id, product_id, quantity)
SELECT currval('orders_id_seq'), product_id, quantity FROM cart;

-- Clear the cart
DELETE FROM cart;

COMMIT;
