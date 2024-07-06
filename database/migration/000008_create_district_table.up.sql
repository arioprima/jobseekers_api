create table if not exists district(
    id varchar(36) primary key,
    name varchar(100),
    city_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp default null,
    foreign key(city_id) references city(id)
);