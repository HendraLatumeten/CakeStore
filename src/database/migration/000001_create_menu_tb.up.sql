create table menucake_tb(
    id int primary key AUTO_INCREMENT,
    title varchar(50) NOT NULL ,
    description text NOT NULL,
    rating int(1) NOT NULL,
    image varchar(200) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)   
