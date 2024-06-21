-- Create tasks table with additional columns for baremetal private IP and baremetal ID
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    -- user_id VARCHAR(20) NOT NULL,
    description TEXT NOT NULL,
    state VARCHAR(20) NOT NULL CHECK (state IN ('create', 'creating', 'created', 'failed', 'deleting', 'deleted')),
    baremetal_private_ip VARCHAR(45) NOT NULL,
    baremetal_id VARCHAR(100) NOT NULL
);
