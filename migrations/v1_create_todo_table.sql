create table todo (
    id varchar(100) not null,
    title varchar(max) not null,
    url varchar(max) not null,
    completed boolean not null,
    order int not null DEFAULT 0,
    text varchar(max) not null DEFAULT ""
);