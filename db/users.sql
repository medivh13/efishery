CREATE TABLE efishery.users (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    phone character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(255) NOT NULL
);

CREATE SEQUENCE efishery.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE efishery.users_id_seq OWNED BY efishery.users.id;

ALTER TABLE ONLY efishery.users ALTER COLUMN id SET DEFAULT nextval('efishery.users_id_seq'::regclass);

ALTER TABLE ONLY efishery.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE INDEX usersindex ON efishery.users USING btree (deleted_at, phone);


INSERT INTO efishery.users (name, phone, role, password) VALUES
('name 1', '11111', 'user', 'jaw9'),
('name 2', '22222', 'user', 'jaw9'),
('name 3', '33333', 'admin', 'MpLm'),
('name 4', '44444', 'admin', 'bwYG'),
('name 5', '55555', 'admin', 'XVlB');