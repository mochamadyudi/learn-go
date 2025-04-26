drop table if exists schema_migrations;

drop table if exists  role_users;

drop table if exists  users;

drop table if exists roles;

drop table if exists permissions;

DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;
DROP EXTENSION IF EXISTS "dblink" CASCADE;