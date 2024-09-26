CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    status boolean,
    created_date timestamp with time zone DEFAULT now(),
    last_login timestamp with time zone DEFAULT now()
);

CREATE TABLE session (
    id integer NOT NULL,
    user_id bigint,
    user_token character varying(300) NOT NULL,
    is_expire boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);