create extension if not exists pgcrypto;

create role budgeter_api noinherit login;

create schema auth;

grant usage on schema auth to budgeter_api;

create table auth.users (
  username text primary key,
  encrypted_password text
);

grant select on auth.users to budgeter_api;
grant insert on auth.users to budgeter_api;
grant update on auth.users to budgeter_api;
grant delete on auth.users to budgeter_api;

insert into auth.users (username, encrypted_password)
values ('admin', crypt('admin', gen_salt('bf')));

create view public.usernames as select
username from auth.users;

grant select on public.usernames to budgeter_api;

create table auth.sessions (
  token text primary key default encode(gen_random_bytes(32), 'base64'),
  username text not null,
  expire_at timestamptz default now() + '10 days'::interval
);

grant select on auth.sessions to budgeter_api;
grant insert on auth.sessions to budgeter_api;
grant update on auth.sessions to budgeter_api;
grant delete on auth.sessions to budgeter_api;

/*
  usage:
  select user_id from auth.verify_session('token goes here');
  if that returns 0 rows, session is invalid
  if that returns 1 row, session is valid, and user's id is returned user_id
*/
create function auth.verify_session(session_token text)
returns table(username text)
language sql
as $$
select username from auth.sessions where token = $1 and expire_at > (select now())
$$;

create procedure auth.delete_expired_sessions()
language sql
as $$
delete from auth.sessions where expire_at < (select now())
$$;

create type worker_type as enum ('floor', 'kitchen');

create table employees (
    name text primary key,
    type worker_type,
    wage numeric
);

grant select on public.employees to budgeter_api;
grant insert on public.employees to budgeter_api;
grant update on public.employees to budgeter_api;
grant delete on public.employees to budgeter_api;

create table day_inputs (
    id date primary key,
    floor_hours jsonb,
    kitchen_hours jsonb,
    food_costs jsonb
);

grant select on public.day_inputs to budgeter_api;
grant insert on public.day_inputs to budgeter_api;
grant update on public.day_inputs to budgeter_api;
grant delete on public.day_inputs to budgeter_api;

create table day_reports (
    id date primary key,
    food_cost_start numeric,
    floor_budget_start numeric,
    kitchen_budget_start numeric,
    food_cost_final numeric,
    floor_budget_final numeric,
    kitchen_budget_final numeric,
    current_wages jsonb
);

grant select on public.day_reports to budgeter_api;
grant insert on public.day_reports to budgeter_api;
grant update on public.day_reports to budgeter_api;
grant delete on public.day_reports to budgeter_api;

