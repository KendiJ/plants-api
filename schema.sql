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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE SET NULL
);

-- Seed Data (Optional: only adds if tables are empty)
INSERT IGNORE INTO rooms (id, name) VALUES 
(1, 'Living Room'),
(2, 'Bedroom'),
(3, 'Kitchen');

INSERT IGNORE INTO plants (name, room_id, water_freq) VALUES 
('Monstera', 1, 7),
('Snake Plant', 2, 14),
('Basil', 3, 3);
