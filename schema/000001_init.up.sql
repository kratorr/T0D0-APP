CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   login VARCHAR UNIQUE NOT NULL,
   password VARCHAR NOT NULL
);


CREATE TABLE states (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE todo_lists (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    title VARCHAR NOT NULL,
    description TEXT
);


CREATE TABLE todo_element (
    id SERIAL PRIMARY KEY,
    todo_list_id INT REFERENCES todo_lists (id),
    title VARCHAR NOT NULL
);