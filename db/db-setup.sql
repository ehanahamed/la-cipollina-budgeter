create extension if not exists pgcrypto;

create role budgeter_api noinherit login;

create schema auth;

grant usage on schema auth to budgeter_api;

create table auth.users (
    id serial primary key,
    username text not null unique,
    encrypted_password text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

grant select on auth.users to budgeter_api;
grant insert on auth.users to budgeter_api;
grant update on auth.users to budgeter_api;
grant delete on auth.users to budgeter_api;
grant usage, select on auth.users_id_seq to budgeter_api;

insert into auth.users (username, encrypted_password)
values ('admin', crypt('admin', gen_salt('bf')));

create table auth.sessions (
    token text primary key default encode(gen_random_bytes(32), 'base64'),
    user_id int not null,
    expire_at timestamptz default now() + '10 days'::interval,
    foreign key(user_id) references auth.users(id) on delete cascade
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
select user_id from auth.sessions where token = $1 and expire_at > (select now())
$$;

create procedure auth.delete_expired_sessions()
language sql
as $$
delete from auth.sessions where expire_at < (select now())
$$;

create table employees (
    id serial primary key,
    name text not null,
    type text, /* 'FLOOR' or 'KITCHEN' */
    wage numeric(10, 2), /* 2 decimal places */
    special_pay jsonb, /* null for default, not special pay using wage column, jsonb for special per-hour or per-day pay
        special pay json example:
        {
            "mon": {
                "perDay": 123.40
            },
            "tue": {
                "perDay": 100.00
                "perHour": 15.00
            },
            "wed": {
                "perHour": 50.00
            },
            "thu": null
            "fri": {}
        }
        in the above example,
        on monday this employee is paid $123.40 if they work on that day
        on tuesday, they're paid $100 if they're working that day plus $15.00 per hour
        on wednesday, they're paid 50.00 per-hour
        this employee won't appear in the menu where hours/days are recorded on thursday or friday
    */
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

grant select on public.employees to budgeter_api;
grant insert on public.employees to budgeter_api;
grant update on public.employees to budgeter_api;
grant delete on public.employees to budgeter_api;
grant usage, select on employees_id_seq to budgeter_api;

create table valentinos (
    id serial primary key,
    name text not null,
    weekly_pay numeric(10, 2),
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

grant select on public.valentinos to budgeter_api;
grant insert on public.valentinos to budgeter_api;
grant update on public.valentinos to budgeter_api;
grant delete on public.valentinos to budgeter_api;
grant usage, select on valentinos_id_seq to budgeter_api;

create table days (
    id serial primary key,
    date date not null unique,
    hours_worked jsonb,
    worked_today jsonb,
    current_employees jsonb,
    food_costs jsonb
);

grant select on public.days to budgeter_api;
grant insert on public.days to budgeter_api;
grant update on public.days to budgeter_api;
grant delete on public.days to budgeter_api;
grant usage, select on days_id_seq to budgeter_api;

create table weeks (
    id serial primary key,
    start_date date not null unique,
    end_date date not null unique,
    start_food_cost_budget numeric(10, 2) not null,
    start_kitchen_pay_budget numeric(10, 2) not null,
    start_floor_pay_budget numeric(10, 2) not null,
    end_food_cost_budget numeric(10, 2),
    end_kitchen_pay_budget numeric(10, 2),
    end_floor_pay_budget numeric(10, 2),
    food_cost_expense numeric(10, 2),
    kitchen_pay_expense numeric(10, 2),
    floor_pay_expense numeric(10, 2)
);
grant usage, select on weeks_id_seq to budgeter_api;

