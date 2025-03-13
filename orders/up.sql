Create table if not exists orders(
    id char(27) primary key,   
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    account_id char(27) NOT NULL,
    total_price MONEY NOT NULL
);



CREAT TABLE IF NOT EXISTS order_products(
    order_id char(27) reference orders (id) on delete cascade,
    product_id char(27)
    quantity INT NOT NULL,
    PRIMARY KEY ( product_id, order_id)
)