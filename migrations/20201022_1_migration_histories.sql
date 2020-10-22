alter table migration_histories add constraint migration_histories_migration_name_unique unique (migration_name);
insert into migration_histories
(
    created_at, migration_name
)
values
(
    current_timestamp,
    '20201022_1_migration_histories.sql'
)