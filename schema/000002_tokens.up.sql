CREATE TABLE tokens (
   id SERIAL PRIMARY KEY,
   user_id INT REFERENCES users (id),
   token 
);




CREATE TABLE todo_lists (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id),
    title VARCHAR NOT NULL,
    description TEXT
);
