INSERT INTO users(name, email)
VALUES ("Budi", "budi@mail.com"),
("Rudi", "Rudi@mail.com"),
("Cindy", "cindy@mail.com"),
("Andi", "andi@mail.com");

INSERT INTO products(name, user_id, description, weight)
VALUES ("Mouse", 1, "Mouse Logitech", 100),
("Keyboard", 2, "Keyboard Logitech", 100),
("Mouse Pad", 1, "Mouse Pad Logitech", 10);

INSERT INTO address(user_id, address)
VALUES (1, "Jakarta"),
(2, "Malang"),
(3, "Surabaya"),
(4, "Blitar");

INSERT INTO receivers(product_id, user_id, address_id, notes)
VALUES (1,3,3,"terima kasih banyak"),
(2,4,4,"terima kasih");

SELECT * FROM users;
SELECT * FROM products;
SELECT * FROM address;
SELECT * FROM receivers;

-- tampilkan semua product beserta nama ownernya
SELECT pd.id, pd.name "product_name", us.name "owner", pd.description FROM products pd
INNER JOIN users us ON pd.user_id=us.id;

-- tampilkan semua product beserta nama owner dan nama receivernya
SELECT pd.id, pd.name "product_name", us.name "owner", rv.user_id, usrv.name "receiver", pd.description FROM products pd
INNER JOIN users us ON pd.user_id=us.id
INNER JOIN receivers rv ON rv.product_id=pd.id
INNER JOIN users usrv ON rv.user_id=usrv.id;

-- tampilkan semua product beserta nama owner dan nama receivernya dan alamat receivernya

