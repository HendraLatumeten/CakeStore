create table menucake_tb(
    id int primary key AUTO_INCREMENT,
    title varchar(200) NOT NULL ,
    description text NOT NULL,
    rating int NOT NULL,
    image varchar(200) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)   