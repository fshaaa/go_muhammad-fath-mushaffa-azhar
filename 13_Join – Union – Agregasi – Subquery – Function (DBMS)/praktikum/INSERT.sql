-- 1. insert 5 operator di tabel operators
INSERT INTO operators(operator_id, name) VALUES 
  (1, "Telkomsel"),
  (2, "XL"),
  (3, "Indosat"),
  (4, "Axis"),
  (5, "Smartfren");

-- 2. insert 3 type product
INSERT INTO product_types(product_type_id, name) VALUES
  (1, "Pulsa"),
  (2, "Paket Internet"),
  (3, "Paket Telepon");

-- 3. insert 2 product with product type id = 1 dan operator id = 3
INSERT INTO products(product_id, product_type_id, operator_id, code, name, status, price) VALUES
  (1, 1, 3, "PLS10KIND", "Pulsa 10000 Indosat", 1, 12000),
  (2, 1, 3, "PLS25KIND", "Pulsa 25000 Indosat", 1, 27000);

-- 4. insert 3 product with product type id = 2 dan operator id = 1
INSERT INTO products(product_id, product_type_id, operator_id, code, name, status, price) VALUES
  (3, 2, 1, "PI3GBTLK", "Paket Internet 3 GB Telkomsel", 1, 30000),
  (4, 2, 1, "PI5GBTLK", "Paket Internet 5 GB Telkomsel", 1, 45000),
  (5, 2, 1, "PI10GBTLK", "Paket Internet 10 GB Telkomsel", 1, 80000);

-- 5. insert 3 product product type id = 3 dan operator id = 4
INSERT INTO products(product_id, product_type_id, operator_id, code, name, status, price) VALUES
  (6, 3, 4, "PT30MNTAX", "Paket Telepon 30 Menit Axis", 1, 10000),
  (7, 3, 4, "PT60MNTAX", "Paket Telepon 60 Menit Axis", 1, 15000),
  (8, 3, 4, "PT90MNTAX", "Paket Telepon 90 Menit Axis ", 1, 20000);

-- 6. insert description di setiap product
INSERT INTO product_descriptions(product_description_id, description) VALUES 
  (1, "Mengisi Pulsa Indosat senilai Rp10.000,- dan menambah masa aktif selama 30 hari"),
  (2, "Mengisi Pulsa Indosat senilai Rp25.000,- dan menambah masa aktif selama 30 hari"),
  (3, "Menambahkan Paket Internet Telkomsel sebesar 3 GB selama 30 hari"),
  (4, "Menambahkan Paket Internet Telkomsel sebesar 5 GB selama 30 hari"),
  (5, "Menambahkan Paket Internet Telkomsel sebesar 10 GB selama 30 hari"),
  (6, "Menambahkan Paket Telepon selama 30 Menit selama 30 hari"),
  (7, "Menambahkan Paket Telepon selama 60 Menit selama 30 hari"),
  (8, "Menambahkan Paket Telepon selama 90 Menit selama 30 hari");
  
-- 7. Insert 3 Payment Methods
INSERT INTO payment_methods(payment_method_id, name, status) VALUES
  (1, "QRIS", 1),
  (2, "Gopay", 1),
  (3, "Dana", 1);

-- 8. insert 5 user
INSERT INTO users(user_id, name, status, dob, gender) VALUES 
  (1, "azhar", 1, "2002/04/15", "M"),
  (2, "akmal", 1, "2002/09/17", "M"),
  (3, "linda", 1, "2003/05/13", "F"),
  (4, "fasya", 1, "2001/09/02", "F"),
  (5, "sultan", 1, "2001/02/01", "M");

-- 9. insert 3 transaksi masing-masing users
INSERT INTO transactions(transaction_id, user_id, payment_method_id, status) VALUES 
  (1, 1, 2, 1),
  (2, 5, 2, 1),
  (3, 3, 1, 1),
  (4, 4, 2, 1),
  (5, 2, 1, 1),
  (6, 4, 1, 1),
  (7, 5, 2, 1),
  (8, 1, 3, 1),
  (9, 3, 2, 1),
  (10, 3, 2, 1),
  (11, 2, 1, 1),
  (12, 1, 2, 1),
  (13, 4, 2, 1),
  (14, 2, 1, 1),
  (15, 5, 3, 1);

-- 10. insert 3 product di masing-masing transaksi
INSERT INTO transaction_details(transaction_id, product_id, status) VALUES
  (1, 2, 1),
  (1, 1, 1),
  (1, 1, 1),
  (2, 3, 1),
  (2, 3, 1),
  (2, 5, 1),
  (3, 6, 1),
  (3, 6, 1),
  (3, 7, 1),
  (4, 8, 1),
  (4, 8, 1),
  (4, 6, 1),
  (5, 1, 1),
  (5, 1, 1),
  (5, 2, 1),
  (6, 2, 1),
  (6, 2, 1),
  (6, 2, 1),
  (7, 5, 1),
  (7, 5, 1),
  (7, 3, 1),
  (8, 4, 1),
  (8, 4, 1),
  (8, 4, 1),
  (9, 6, 1),
  (9, 8, 1),
  (9, 7, 1),
  (10, 2, 1),
  (10, 2, 1),
  (10, 1, 1),
  (11, 3, 1),
  (11, 3, 1),
  (11, 4, 1),
  (12, 2, 1),
  (12, 1, 1),
  (12, 1, 1),
  (13, 8, 1),
  (13, 7, 1),
  (13, 6, 1),
  (14, 1, 1),
  (14, 1, 1),
  (14, 1, 1),
  (15, 5, 1),
  (15, 3, 1),
  (15, 3, 1);
