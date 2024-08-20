# locoAssignment
In the above assignment, I aimed to keep the solution as straightforward as possible. To maintain simplicity, I did not use any schema management tools and env variable focusing instead on the core functionality.

Before testing, please update the database URL(dbUrl) in the code to match your local setup. Additionally, you will need to manually create the transactions table. Here is the SQL query to create the table:
# for creating table
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    type VARCHAR(255) NOT NULL,
    parent_id INTEGER REFERENCES transactions(id)
);
