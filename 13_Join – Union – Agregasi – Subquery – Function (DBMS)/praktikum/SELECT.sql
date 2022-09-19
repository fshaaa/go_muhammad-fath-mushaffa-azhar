-- menampilkan user laki-laki
SELECT name FROM users WHERE gender = 'M';

-- tampilkan product_id = 3
SELECT name FROM products WHERE product_id = 3;

-- tampilkan pelanggan yang created at kurang dalam seminggu 
-- dan mempunyai nama yang terdapat 'a'
SELECT name FROM users WHERE created_at > (DATE_SUB(CURDATE(), INTERVAL 7 DAY)) AND name LIKE "%a%"