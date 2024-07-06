create table if not exists otp_code(
    id varchar(36) primary key,
    user_id varchar(36),
    code varchar(100),
    expired_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key (user_id) references users(id)
);