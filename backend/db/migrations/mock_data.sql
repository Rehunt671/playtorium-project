-- Insert mock data
INSERT INTO users (name, username, password, points) VALUES
('John Doe', 'johndoe', 'password123', 100),
('Jane Smith', 'janesmith', 'password123', 200);

INSERT INTO item_categories (name) VALUES
('clothing'),
('accessories'),
('electronics');

INSERT INTO items (name, price, item_category_id, stock, image_name, description) VALUES
('T-Shirt', 350, 1, 100, 'tshirt.png', 'A comfortable cotton T-Shirt'),
('Hat', 250, 2, 50, 'hat.png', 'A stylish hat for sunny days'),
('Hoodie', 700, 1, 30, 'hoodie.png', 'A warm and cozy hoodie'),
('Watch', 850, 3, 20, 'watch.png', 'A sleek and modern watch'),
('Bag', 640, 2, 40, 'bag.png', 'A durable and spacious bag');


INSERT INTO carts (user_id) VALUES
(1),
(2);

INSERT INTO cart_items (cart_id, item_id, quantity) VALUES
(1, 1, 2),
(1, 2, 1),
(2, 3, 1),
(2, 4, 1);

INSERT INTO discount_categories (name) VALUES
('Coupon'),
('On Top'),
('Seasonal');

INSERT INTO campaigns (name) VALUES
('FIXED_AMOUNT'),
('PERCENTAGE'),
('PERCENTAGE_CATEGORY'),
('POINT'),
('SEASONAL');

INSERT INTO campaign_on_discount_categories (campaign_id, discount_category_id) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 3);

INSERT INTO fixed_amount_discounts (amount) VALUES
(50);

INSERT INTO point_discounts (point) VALUES
(20);

INSERT INTO percentage_discounts (percentage) VALUES
(10);

INSERT INTO percentage_category_discounts (item_category_id, percentage) VALUES
(1, 10);

INSERT INTO seasonal_discounts (every_xthb, discount_ythb) VALUES
(300, 40);

INSERT INTO discounts (fixed_amount_id, discount_category_id) VALUES
(1, 1);

INSERT INTO discounts (percentage_category_id, discount_category_id) VALUES
(1, 2);

INSERT INTO discounts (seasonal_id, discount_category_id) VALUES
(1, 3);