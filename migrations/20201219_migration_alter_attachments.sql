alter table tags_posts_attachment rename to tags_posts_attachments;

insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20201219_migration_alter_attachments.sql'
);
