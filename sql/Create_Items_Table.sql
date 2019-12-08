CREATE TABLE `Items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `vendor_id` INT NOT NULL,
  `blue_essence` INT,
  `riot_points` INT,
  `solo` varchar(100),
  `flex` varchar(100),
  `price_dollars` INT,
  `price_cents` INT,
  `level` INT,
  `email` varchar(100),
  `password` varchar(100),
  `login` varchar(100),
  `login_password` varchar(100)
);

CREATE TABLE `Champions` (
  id INT AUTO_INCREMENT PRIMARY_KEY,
  name varchar(100),
  item_id INT,
  FOREIGN KEY(item_id) REFERENCES Items(id)
)