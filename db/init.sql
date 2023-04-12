-- Create users table and insert admin user
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, email, password) VALUES ('admin', 'admin@example.com', '$2a$10$JnJug/exipCaBCsoLbAcI.9PygrXxq7UryCSlKtkXgk3NBZhsaSEO');

-- Create potions table and insert sample data
CREATE TABLE IF NOT EXISTS potions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    properties VARCHAR(255) NOT NULL
);

INSERT INTO potions (name, description, price, properties) VALUES
('Elixir of Endless Endurance', 'A rejuvenating brew that imbues the drinker with seemingly limitless stamina.', '50.0', '{"Stamina", "Endurance"}'),
('Potion of the Wandering Mind', 'A concoction that enhances the imagination and creative thinking of the drinker.', '25.0', '{"Creativity", "Imagination"}'),
('Brew of the Fierce Flame', 'A fiery brew that imbues the drinker with the power of flame.', '75.0', '{"Fire", "Power"}'),
('Essence of the Silver Tongue', 'A smooth elixir that enhances the persuasive abilities of the drinker.', '40.0', '{"Persuasion", "Charisma"}'),
('Potion of the Calm Seas', 'A calming brew that soothes the nerves of the drinker.', '20.0', '{"Calm", "Relaxation"}'),
('Syrup of the Golden Glow', 'A sweet syrup that enhances the natural radiance of the drinker.', '30.0', '{"Radiance", "Beauty"}'),
('Tonic of the Unbreakable Will', 'A potent tonic that strengthens the willpower of the drinker.', '60.0', '{"Willpower", "Strength"}'),
('Philter of the Swift Feet', 'A swift potion that enhances the speed and agility of the drinker.', '80.0', '{"Speed", "Agility"}'),
('Balm of the Soaring Spirit', 'A restorative balm that uplifts the spirit of the drinker.', '35.0', '{"Uplift", "Spirit"}'),
('Elixir of the Shrouded Mind', 'A mysterious elixir that enhances the intuition and foresight of the drinker.', '45.0', '{"Intuition", "Foresight"}'),
('Sap of the Mighty Oak', 'A thick sap that enhances the strength and endurance of the drinker.', '70.0', '{"Strength", "Endurance"}');

-- Create reviews table and insert sample data
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    potion_id INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    comment TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (potion_id) REFERENCES potions(id)
);

INSERT INTO reviews (user_id, potion_id, rating, comment) VALUES
(1, 1, 4, 'Great product!'),
(1, 2, 5, 'Highly recommend this potion.'),
(1, 3, 3, 'Not bad, but could be better.'),
(1, 4, 4, 'I was skeptical at first, but this really works.'),
(1, 5, 2, 'Didn''t do much for me, unfortunately.'),
(1, 6, 5, 'Absolutely love this syrup!'),
(1, 7, 4, 'Definitely feel more focused and motivated after drinking this.'),
(1, 8, 5, 'WOW, what a difference!'),
(1, 9, 3, 'It was okay, but not as effective as I hoped.'),
(1, 10, 4, 'I''m really impressed with this elixir.');

-- Create cart_items table and insert sample data
CREATE TABLE IF NOT EXISTS cart_items (
id SERIAL PRIMARY KEY,
user_id INTEGER NOT NULL,
potion_id INTEGER NOT NULL,
quantity INTEGER NOT NULL,
FOREIGN KEY (user_id) REFERENCES users(id),
FOREIGN KEY (potion_id) REFERENCES potions(id)
);

INSERT INTO cart_items (user_id, potion_id, quantity) VALUES
(1, 1, 2),
(1, 2, 1),
(1, 3, 3),
(1, 4, 2),
(1, 5, 1),
(1, 6, 2),
(1, 7, 1),
(1, 8, 3),
(1, 9, 1),
(1, 10, 2);