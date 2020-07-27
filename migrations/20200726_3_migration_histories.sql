alter table migration_history rename to migration_histories;

insert into
    migration_histories
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20200726_3_migration_histories.sql'
);
