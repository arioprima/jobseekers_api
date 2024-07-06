create table if not exists pivot_certification(
    id varchar(36) primary key,
    certification_id varchar(36),
    user_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key(certification_id) references certification(id),
    foreign key(user_id) references users(id)
);