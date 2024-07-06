create table if not exists career(
    id varchar(36) primary key,
    position varchar(100),
    company varchar(100),
    start_date date,
    end_date date default null,
    description text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp default null
);