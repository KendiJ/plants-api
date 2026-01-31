SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS plants;
DROP TABLE IF EXISTS rooms;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE rooms (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE plants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    room_id INT, 
    water_freq INT COMMENT 'Days between watering',
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE SET NULL
);

-- INSERT ROOMS 
INSERT INTO rooms (id, name, image_url) VALUES 
(1, 'Living Room', 'https://images.unsplash.com/photo-1583847268964-b28dc8f51f92?auto=format&fit=crop&w=800&q=80'),
(2, 'Bedroom', 'https://images.unsplash.com/photo-1615874959474-d609969a20ed?q=80&w=1480&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(3, 'Kitchen', 'https://images.unsplash.com/photo-1721742146260-dd19d80e1d52?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(4, 'Bathroom', 'https://plus.unsplash.com/premium_photo-1661884424253-08db7c7758ce?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(5, 'Dining Room', 'https://plus.unsplash.com/premium_photo-1674815329400-f17578e2b03b?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(6, 'Family Room', 'https://plus.unsplash.com/premium_photo-1736194028567-54d9c9f3cc56?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(7, 'Home Office', 'https://images.unsplash.com/photo-1641178796760-a35e54c27aed?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(8, 'Foyer/Entryway', 'https://images.unsplash.com/photo-1560185127-59e4420e2c93?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
(9, 'Laundry Room', 'https://images.unsplash.com/photo-1655041448985-f6666cba2d6c?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D');

-- INSERT PLANTS 
INSERT INTO plants (name, room_id, water_freq, image_url) VALUES 
('Monstera Deliciosa', 1, 7, 'https://images.unsplash.com/photo-1614594975525-e45190c55d0b?auto=format&fit=crop&w=600&q=80'),
('Fiddle Leaf Fig', 1, 10, 'https://images.unsplash.com/photo-1622673037899-d7914ece9fe8?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 1, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 1, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 1, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Bird of Paradise', 1, 4, 'https://plus.unsplash.com/premium_photo-1677653224336-9bad06bd859f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Rubber Plant', 1, 8, 'https://images.unsplash.com/photo-1616122236015-b37dc314d875?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Kentia Palm', 1, 3, 'https://images.unsplash.com/photo-1697028896377-07ac7b746f01?q=80&w=1064&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 1, 6, 'https://images.unsplash.com/photo-1611527664689-d430dd2a6774?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 2, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 2, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Pothos', 2, 4, 'https://images.unsplash.com/photo-1595524147656-eb5d0a63e9a9?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 2, 6, 'https://images.unsplash.com/photo-1611527664689-d430dd2a6774?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Philodendron', 2, 4, 'https://images.unsplash.com/photo-1611211233623-1b1e2162633f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Calathea', 2, 5, 'https://images.unsplash.com/photo-1616276595426-81d5a3cfc772?q=80&w=1027&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Basil', 3, 3, 'https://images.unsplash.com/photo-1619805640532-21cce5fe542b?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Rosemary', 3, 4, 'https://images.unsplash.com/photo-1515586000433-45406d8e6662?q=80&w=1065&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Mint', 3, 3, 'https://images.unsplash.com/photo-1588908933351-eeb8cd4c4521?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 3, 2, 'https://images.unsplash.com/photo-1646295671963-55e360bab51f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('English Ivy', 3, 4, 'https://images.unsplash.com/photo-1542347369-6fb75718485c?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Boston Fern', 3, 1, 'https://images.unsplash.com/photo-1515586000433-45406d8e6662?q=80&w=1065&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Aloe Vera', 3, 14, 'https://images.unsplash.com/photo-1671166739837-b175ef95cb48?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 4, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Philodendron', 4, 4, 'https://images.unsplash.com/photo-1611211233623-1b1e2162633f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 4, 6, 'https://images.unsplash.com/photo-1611527664689-d430dd2a6774?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 4, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 4, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Boston Fern', 4, 1, 'https://images.unsplash.com/photo-1515586000433-45406d8e6662?q=80&w=1065&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Birds Nest Fern', 4, 3, 'https://images.unsplash.com/photo-1636901942318-972ea62b4d5d?q=80&w=1064&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Orchids', 4, 3, 'https://images.unsplash.com/photo-1618080578815-335456280012?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Monstera Deliciosa', 5, 7, 'https://images.unsplash.com/photo-1614594975525-e45190c55d0b?auto=format&fit=crop&w=600&q=80'),
('Bird of Paradise', 5, 4, 'https://plus.unsplash.com/premium_photo-1677653224336-9bad06bd859f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Rubber Plant', 5, 8, 'https://images.unsplash.com/photo-1616122236015-b37dc314d875?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Kentia Palm', 5, 3, 'https://images.unsplash.com/photo-1697028896377-07ac7b746f01?q=80&w=1064&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 5, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Palor Palm', 5, 4, 'https://images.unsplash.com/photo-1608270691744-3bd1b5a7833f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D' ),
('Monstera Deliciosa', 6, 7, 'https://images.unsplash.com/photo-1614594975525-e45190c55d0b?auto=format&fit=crop&w=600&q=80'),
('Fiddle Leaf Fig', 6, 10, 'https://images.unsplash.com/photo-1622673037899-d7914ece9fe8?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 6, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 6, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 6, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Bird of Paradise', 6, 4, 'https://plus.unsplash.com/premium_photo-1677653224336-9bad06bd859f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Rubber Plant', 6, 8, 'https://images.unsplash.com/photo-1616122236015-b37dc314d875?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Kentia Palm', 6, 3, 'https://images.unsplash.com/photo-1697028896377-07ac7b746f01?q=80&w=1064&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 6, 6, 'https://images.unsplash.com/photo-1611527664689-d430dd2a6774?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 7, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 7, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Pothos', 7, 4, 'https://images.unsplash.com/photo-1595524147656-eb5d0a63e9a9?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 7, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Philodendron', 7, 4, 'https://images.unsplash.com/photo-1611211233623-1b1e2162633f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 7, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Cactus Plant', 7, 1, 'https://images.unsplash.com/photo-1586633070827-b1de85383601?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 7, 6, 'https://images.unsplash.com/photo-1611527664689-d430dd2a6774?q=80&w=2073&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Snake Plant', 8, 14, 'https://images.unsplash.com/photo-1616961065849-edf307a08bcb?q=80&w=1035&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 8, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Kentia Palm', 8, 3, 'https://images.unsplash.com/photo-1697028896377-07ac7b746f01?q=80&w=1064&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Monstera Deliciosa', 8, 7, 'https://images.unsplash.com/photo-1614594975525-e45190c55d0b?auto=format&fit=crop&w=600&q=80'),
('Fiddle Leaf Fig', 8, 10, 'https://images.unsplash.com/photo-1622673037899-d7914ece9fe8?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Boston Fern', 9, 1, 'https://images.unsplash.com/photo-1515586000433-45406d8e6662?q=80&w=1065&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Spider Plant', 9, 2, 'https://images.unsplash.com/photo-1646295671963-55e360bab51f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Peace Lily', 9, 5, 'https://images.unsplash.com/photo-1593691509543-c55fb32d8de5?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Pothos', 9, 4, 'https://images.unsplash.com/photo-1595524147656-eb5d0a63e9a9?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('Philodendron', 9, 4, 'https://images.unsplash.com/photo-1611211233623-1b1e2162633f?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'),
('ZZ Plant', 9, 3, 'https://images.unsplash.com/photo-1606256419855-d72ce8675863?q=80&w=987&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D');