create table migration_history (
    created_at timestamptz,
    migration_name varchar
);

alter table migration_history owner to postgres;

comment on table migration_history is 'DBのMigration履歴用テーブル';
comment on column  migration_history.created_at is '作成日時';
insert into
    migration_history
(
    created_at,
    migration_name
)
values (
     current_timestamp,
     '20200726_1_migration_history.sql'
);
