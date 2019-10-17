DROP INDEX IF EXISTS refresh_token_uindex;
DROP INDEX IF EXISTS access_token_uindex;
DROP INDEX IF EXISTS sessions_id_uindex;
DROP TABLE IF EXISTS sessions;

DROP INDEX IF EXISTS actions_roles_action_title_role_id_uindex;
DROP TABLE IF EXISTS actions_roles;

DROP INDEX IF EXISTS users_roles_user_id_role_id_uindex;
DROP TABLE IF EXISTS users_roles;

DROP INDEX IF EXISTS actions_title_uindex;
DROP INDEX IF EXISTS actions_id_uindex;
DROP TABLE IF EXISTS actions;

DROP INDEX IF EXISTS roles_title_uindex;
DROP INDEX IF EXISTS roles_id_uindex;
DROP TABLE IF EXISTS roles;

DROP INDEX IF EXISTS users_login_uindex;
DROP INDEX IF EXISTS users_id_uindex;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";
