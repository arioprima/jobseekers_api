create table if not exists career_history(
    id varchar(36) primary key,
    user_id varchar(36),
    career_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key(user_id) references users(id),
    foreign key(career_id) references career(id)
);