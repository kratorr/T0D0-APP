CREATE TABLE tokens (
   user_id INT REFERENCES users (id),
   token VARCHAR(255),
   expiration_date TIMESTAMP

);
