-- menampilkan user laki-laki
SELECT name FROM users WHERE gender = 'M';

-- tampilkan product_id = 3
SELECT name FROM products WHERE product_id = 3;

-- tampilkan pelanggan yang created at kurang dalam seminggu 
-- dan mempunyai nama yang terdapat 'a'
SELECT name FROM users WHERE created_at > (DATE_SUB(CURDATE(), INTERVAL 7 DAY)) AND name LIKE "%a%";

-- hitung jumlah user dengan status gender perempuan
SELECT COUNT(user_id) FROM users WHERE gender = 'F';

-- tampilkan data pelanggan dengan nama urut abjad
SELECT * FROM users ORDER BY name ASC;

-- tampilkan 5 data product
SELECT * FROM products LIMIT 5;