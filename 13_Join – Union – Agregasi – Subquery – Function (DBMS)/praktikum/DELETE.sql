-- delete data product id = 1
DELETE FROM product_descriptions WHERE product_description_id = 1;
DELETE FROM transaction_details WHERE product_id = 1;
DELETE FROM products WHERE product_id = 1;

-- delete product type id = 1
ALTER TABLE product_descriptions
DROP FOREIGN KEY fk_product_description_id;
ALTER TABLE transaction_details
DROP FOREIGN KEY fk_product_id;
DELETE FROM products WHERE product_type_id = 1;