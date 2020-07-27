create table users (
    created_at timestamptz,
    updated_at timestamptz,
    id varchar not null primary key,    
    email varchar not null unique,
    hashed_password varchar not null
);
alter table users owner to postgres;

comment on table users is 'ユーザ情報';

comment on column users.created_at         is '作成日時';
comment on column users.updated_at         is '更新日時';
comment on column users.id                 is 'ユーザID';
comment on column users.email              is 'メールアドレス';
comment on column users.hashed_password    is 'ハッシュ化されたパスワード';

insert into
    migration_history
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20200726_2_users.sql'
);
