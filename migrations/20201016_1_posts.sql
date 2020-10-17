create table posts (
    created_at timestamptz,
    updated_at timestamptz,
    id uuid primary key,
    user_id varchar not null,
    constraint fk_user_id foreign key (user_id) references users (id),
    title text not null,
    body text,
    plain_body text,
    published boolean
);

alter table posts owner to postgres;

comment on table posts is 'ポスト';

comment on column posts.created_at       is '作成日時';
comment on column posts.updated_at       is '更新日時';
comment on column posts.id               is 'ポストのID';
comment on column posts.title            is 'ポストのタイトル';
comment on column posts.body             is 'ポストの本文(markdown形式)';
comment on column posts.plain_body       is 'ポストの本文(レンダリングされたmarkdownを平文化)';
comment on column posts.published        is 'ポストが公開されているか(trueなら公開, falseなら下書き状態)';


insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20201016_1_posts.sql'
);
