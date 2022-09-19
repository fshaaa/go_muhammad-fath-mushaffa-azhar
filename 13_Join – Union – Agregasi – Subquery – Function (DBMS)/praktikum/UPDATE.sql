-- ubah nama product id = 1 dengan dummy product
UPDATE products SET name = "product dummy" WHERE product_id = 1;

-- ubah qty = 3 pada transaction detail dengan product id = 1
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;