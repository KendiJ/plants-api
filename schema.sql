CREATE TABLE IF NOT EXISTS rooms (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS plants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    room_id INT, 
    water_freq INT COMMENT 'Days between watering',
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE SET NULL
);

-- Seed Data (Optional: only adds if tables are empty)
INSERT IGNORE INTO rooms (id, name) VALUES 
(1, 'Living Room'),
(2, 'Bedroom'),
(3, 'Kitchen'),
(4, 'Bathroom'),
(5, 'Dining Room'),
(6, 'Family Room'),
(7, 'Home Office'),
(8, 'Foyer/Entryway'),
(9, 'Laundry Room');

INSERT IGNORE INTO plants (name, room_id, water_freq, image_url) VALUES 
('Monstera', 1, 7, 'https://images.unsplash.com/photo-1614594975525-e45190c55d0b?auto=format&fit=crop&w=400&q=80'),
('Snake Plant', 2, 14, 'https://unsplash.com/photos/green-snake-plant-in-brown-woven-pot-fWPyYuOq08g'),
('Basil', 3, 3, 'https://unsplash.com/photos/green-vegeatables-rICRgergpIc');


