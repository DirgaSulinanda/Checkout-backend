-- Create product table
CREATE TABLE IF NOT EXISTS product (
  sku varchar (250) NOT NULL,
  name varchar (250) NOT NULL,
  price decimal (20,2),
  quantity int,
  PRIMARY KEY (sku)
);

-- Fill product table
INSERT INTO product (sku, name, price, quantity) VALUES ('120P90', 'Google Home', 49.99, 10);
INSERT INTO product (sku, name, price, quantity) VALUES ('43N23P', 'MacBook Pro', 5399.99, 5);
INSERT INTO product (sku, name, price, quantity) VALUES ('A304SD', 'Alexa Speaker', 109.50, 10);
INSERT INTO product (sku, name, price, quantity) VALUES ('234234', 'Raspberry Pi B', 30, 2);

-- Create promo table
CREATE TABLE IF NOT EXISTS promo (
  id serial,
  name varchar (250),
  description text,
  formula varchar (500),
  enabled boolean,
  PRIMARY KEY (id)
);

-- Fill promo table
INSERT INTO promo (name, description, formula, enabled) VALUES ('Free Rasp Pi', 'Each sale of a MacBook Pro comes with a free Raspberry Pi B', '1*{43N23P}=1*{234234}',true);
INSERT INTO promo (name, description, formula, enabled) VALUES ('Buy 2 get 1 free Google Home', 'Buy 3 Google Homes for the price of 2', '3*{120P90}=1*{120P90}', true);
INSERT INTO promo (name, description, formula, enabled) VALUES ('Alexa Spreaker 10% discount', 'Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa Speakers', '3*{A304SD}=0.1n*{A304SD}',true);

-- Create checkout_header table
CREATE TABLE IF NOT EXISTS checkout_header (
  id serial,
  date timestamp,
  cashier_name varchar (250),
  PRIMARY KEY (id)
);

-- Create checkout_detail table
CREATE TABLE IF NOT EXISTS checkout_detail (
  id serial,
  header_id int,
  sku varchar (250),
  quantity int,
  promo_ids varchar (250),
  discount_price decimal (20,2),
  PRIMARY KEY (id)
);