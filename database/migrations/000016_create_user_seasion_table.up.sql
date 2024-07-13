create table user_sessions(
    id SERIAL PRIMARY KEY,
    user_id varchar(36) not null,
    token varchar(255) not null,
    last_login timestamp,
    expired_at timestamp not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users(id)
);