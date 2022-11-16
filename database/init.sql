CREATE TABLE users
(
    id uint NOT NULL,
    name text,
    balance int,
    CONSTRAINT user_pkey PRIMARY KEY(id)
);

INSERT INTO users(id, name) VALUES
                                (1,'U1',100),
                                (2,'U2',200),
                                (3,'U3',100);