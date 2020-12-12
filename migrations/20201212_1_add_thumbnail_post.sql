alter table posts add column thumbnail varchar not null default 'https://northernlife-content.net/lunch.jpg';
comment on column posts.thumbnail is 'サムネ画像のリンク';

insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20201212_1_add_thumbnail_post.sql'
);
