-- gabungkan data user id = 1 dan 2
SELECT * FROM users WHERE user_id = 1 OR user_id = 2;

-- tampilkan jumlah harga transaksi user id 1
SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- tampilkan total transaksi dengan product type 2
SELECT COUNT(t.transaction_id) FROM transactions t 
INNER JOIN transaction_details d ON t.transaction_id = d.transaction_id
INNER JOIN products p ON p.product_id = d.product_id
INNER JOIN product_types a ON a.product_type_id = p.product_type_id
WHERE a.product_type_id = 2;

-- tampilkan semua field table product dan field name table product type
-- yang saling berhubungan
SELECT p.*, pt.name FROM products p, product_types pt
WHERE p.product_type_id = pt.product_type_id; 

-- tampilkan semua field transaction dan field nama table product dan 
-- field nama user
SELECT t.*, p.name, u.name FROM transactions t, products p, users u, transaction_details td 
WHERE t.user_id = u.user_id AND t.transaction_id = td.transaction_id
AND td.product_id = p.product_id;

-- buat function setelah data transaksi dihapus maka transaction detail 
-- juga terhapus
DELIMITER $$
CREATE TRIGGER delete_transaction  
  BEFORE DELETE ON transactions
  FOR EACH ROW
BEGIN
  DELETE FROM transaction_details td
  WHERE td.transaction_id = OLD.transaction_id;
END; $$
DELIMITER ;

-- function setelah data transaksi detail dihapus maka total qty akan terupdate
-- dengan sisa qty
DELIMITER $$
CREATE TRIGGER delete_transaction_details  
  AFTER DELETE ON transaction_details
  FOR EACH ROW
BEGIN
  UPDATE transactions
  SET total_qty = (SELECT SUM(td.qty) FROM transaction_details td WHERE td.transaction_id = OLD.transaction_id);
END; $$
DELIMITER ;

-- tampilkan data products yang tidak pernah ada hubungan label transaction_detail
SELECT * FROM products
WHERE product_id NOT IN (SELECT DISTINCT product_id FROM transaction_details);