create table if not exists city(
    id varchar(36) primary key,
    name varchar(100),
    province_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key(province_id) references province(id)
);