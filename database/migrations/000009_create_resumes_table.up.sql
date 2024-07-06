create table if not exists resumes(
    id varchar(36) primary key,
    resume varchar(100),
    user_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key(user_id) references users(id)
)