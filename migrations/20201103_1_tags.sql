create table tags
(
    created_at timestamptz,
    updated_at timestamptz,
    id uuid primary key,
    tag_name varchar unique,
    user_id varchar not null,
    constraint fk_user_id foreign key (user_id) references users (id)
);
alter table tags owner to postgres;
comment on table tags is 'タグ';
comment on column tags.created_at       is '作成日時';
comment on column tags.updated_at       is '更新日時';
comment on column tags.id               is 'ID';
comment on column tags.tag_name         is 'タグ名';
comment on column tags.user_id          is 'タグの作成者';

create table tags_posts_attachment
(
    created_at timestamptz,
    updated_at timestamptz,
    id uuid primary key,
    post_id uuid not null,
    constraint fk_post_id foreign key (post_id) references posts(id),
    tag_id uuid not null,
    constraint fk_tag_id foreign key (tag_id) references tags(id)
);
alter table tags_posts_attachment owner to postgres;
comment on column tags_posts_attachment.created_at       is '作成日時';
comment on column tags_posts_attachment.updated_at       is '更新日時';
comment on column tags_posts_attachment.id               is 'ID';
comment on column tags_posts_attachment.post_id          is 'ポストのID';
comment on column tags_posts_attachment.tag_id           is 'タグのID';

insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20201103_1_tags.sql'
);
