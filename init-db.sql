-- Create tasks table with additional columns for baremetal private IP and baremetal ID
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    state VARCHAR(20) NOT NULL,
    baremetal_private_ip VARCHAR(45) NOT NULL,
    baremetal_id VARCHAR(100) NOT NULL
);

-- Create a notification function
CREATE OR REPLACE FUNCTION notify_task()
RETURNS trigger AS $$
BEGIN
    PERFORM pg_notify('task_channel', row_to_json(NEW)::text);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a function to handle task insert
CREATE OR REPLACE FUNCTION handle_task_insert()
RETURNS trigger AS $$
BEGIN
    -- Ensure the initial state is 'creating'
    IF NEW.state IS DISTINCT FROM 'creating' THEN
        RAISE EXCEPTION 'Invalid initial state: %', NEW.state;
    END IF;

    -- Simulate some processing (e.g., set state to 'created')
    NEW.state := 'created';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a function to handle task update
CREATE OR REPLACE FUNCTION handle_task_update()
RETURNS trigger AS $$
BEGIN
    -- Ensure the state transition is valid
    IF OLD.state = 'created' AND NEW.state = 'deleting' THEN
        -- Simulate some processing for deleting
        RETURN NEW;
    ELSIF OLD.state = 'deleting' AND NEW.state = 'deleted' THEN
        -- Simulate some processing for deleted
        RETURN NEW;
    ELSE
        RAISE EXCEPTION 'Invalid state transition from % to %', OLD.state, NEW.state;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for insert
CREATE TRIGGER trigger_task_insert
BEFORE INSERT ON tasks
FOR EACH ROW
EXECUTE FUNCTION handle_task_insert();

-- Create trigger for update
CREATE TRIGGER trigger_task_update
BEFORE UPDATE ON tasks
FOR EACH ROW
WHEN (OLD.state IS DISTINCT FROM NEW.state)
EXECUTE FUNCTION handle_task_update();
