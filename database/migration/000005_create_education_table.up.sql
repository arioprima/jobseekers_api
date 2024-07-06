create table if not exists education(
    id varchar(36) primary key,
    qualification varchar(100),
    institution varchar(100),
    graduated_at date,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
    deleted_at timestamp default null
);