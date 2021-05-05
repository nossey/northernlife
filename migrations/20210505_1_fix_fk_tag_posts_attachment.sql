alter table tags_posts_attachments drop constraint fk_post_id;
alter table tags_posts_attachments add constraint fk_post_id foreign key (post_id) references posts(id) on delete cascade;
alter table tags_posts_attachments drop constraint fk_tag_id;
alter table tags_posts_attachments add constraint fk_tag_id foreign key (tag_id) references tags(id) on delete cascade;

insert into
	migration_histories 
(
	created_at,
	migration_name
)
values
(
	current_timestamp,
	'20210505_1_fix_fk_tag_posts_attachment.sql'
)