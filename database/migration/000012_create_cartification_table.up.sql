create table if not exists certification(
    id varchar(36) primary key,
    name varchar(100) not null,
    organization varchar(100) not null,
    date_of_issue date not null,
    date_of_expiry date default null,
    description text,
    certificate_image varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    deleted_at timestamp default null
);