create table if not exists biodata(
    id varchar(36) primary key,
    fisrtname varchar(100),
    lastname varchar(100),
    email varchar(100),
    phone varchar(100),
    birthdate date,
    birthplace varchar(100),
    province_id varchar(36),
    city_id varchar(36),
    district_id varchar(36),
    address varchar(100),
    education_id varchar(36),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp default null
);