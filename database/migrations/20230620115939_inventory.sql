-- +goose Up
CREATE TABLE IF NOT EXISTS inventories (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    user_id INT NOT NULL ,
    name CHAR(60) NOT NULL ,
    quantity INT NOT NULL ,
    sku_number VARCHAR(255) NOT NULL ,
    notes VARCHAR(255) NULL ,
    created_at TIMESTAMP NULL ,
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL , PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE inventories;
