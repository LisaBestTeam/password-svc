-- +migrate Up

create table password
(
    id                bigint,
    hash_of_file      text,
    sender_address    text,
    receiver_address  text,
    encrypts_password text,
    type_of_file      text
);

-- +migrate Down

drop table password;