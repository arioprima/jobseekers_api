create table if not exists users(
    id varchar(36) primary key,
    biodata_id varchar(36),
    password varchar(255) not null,
    is_active boolean default true,
    is_verified boolean default false,
    profile_image varchar(255),
    role_id varchar(36),
    summary text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default null,
    foreign key (biodata_id) references biodata(id),
    foreign key (role_id) references user_roles(id)
)