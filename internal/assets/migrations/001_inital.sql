-- +migrate Up

create table password
(
    id                INT,
    hash_of_file      TEXT,
    sender_address    TEXT,
    receiver_address  TEXT,
    encrypts_password TEXT,
    type_of_file      TEXT
);

-- +migrate Down

drop table password;