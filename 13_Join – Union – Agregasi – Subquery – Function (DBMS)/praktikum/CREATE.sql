-- create table users
CREATE TABLE users(
  user_id int PRIMARY KEY,
  name varchar(255),
  status SMALLINT,
  dob DATE,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--create table payment_methods
CREATE TABLE payment_methods(
  payment_method_id int PRIMARY KEY,
  name varchar(255),
  status SMALLINT, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- create table product_types
CREATE TABLE product_types (
  product_type_id int PRIMARY KEY,
  name varchar(255), 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- create table product_types
CREATE TABLE operators (
  operator_id int PRIMARY KEY,
  name varchar(255), 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- create table product
CREATE TABLE products(
  product_id int PRIMARY KEY,
  product_type_id int,
  operator_id int,
  code varchar(255), 
  name varchar(255),
  status SMALLINT, 
  price int, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_product_type_id FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id),
  CONSTRAINT fk_operator_id FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);

-- create table product_descriptions 
CREATE TABLE product_descriptions(
  product_description_id int PRIMARY KEY,
  description TEXT, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
  CONSTRAINT fk_product_description_id FOREIGN KEY (product_description_id) REFERENCES products(product_id)
);

-- create table transactions
CREATE TABLE transactions(
  transaction_id INT PRIMARY KEY,
  user_id INT,
  payment_method_id INT,
  status SMALLINT, 
  total_qty INT,
  total_price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
  CONSTRAINT fk_payment_method_id FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id)
);

-- create transaction_details
CREATE TABLE transaction_details(
  transaction_id INT,
  product_id INT,
  status VARCHAR(10),
  qty int, 
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products(product_id),
  CONSTRAINT fk_transaction_id FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
);

DELIMITER $$
CREATE TRIGGER transaction_detail_trigger
  BEFORE INSERT ON transaction_details
  FOR EACH ROW 
BEGIN
  IF NEW.qty IS NULL THEN
    SET NEW.qty := 1;
  END IF;
  IF NEW.price IS NULL THEN
    SET NEW.price = (SELECT p.price FROM products p WHERE p.product_id = NEW.product_id);
  END IF;
  UPDATE transactions t
  SET t.total_qty = (SELECT COUNT(d.transaction_id) FROM transaction_details d WHERE d.transaction_id = t.transaction_id),
  t.total_price = (SELECT SUM(d.price) FROM transaction_details d WHERE d.transaction_id = t.transaction_id);
END $$
DELIMITER ; 