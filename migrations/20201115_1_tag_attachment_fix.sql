alter table tags_posts_attachment drop column updated_at;

insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20201115_1_tag_attachment_fix.sql'
);
