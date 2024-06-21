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
    -- Ensure the initial state is 'create'
    IF NEW.state IS DISTINCT FROM 'create' THEN
        RAISE EXCEPTION 'Invalid initial state: %', NEW.state;
    END IF;

    -- Change state to 'creating'
    NEW.state := 'creating';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for insert
CREATE TRIGGER trigger_task_insert
BEFORE INSERT ON tasks
FOR EACH ROW
EXECUTE FUNCTION handle_task_insert();

-- Create trigger for notify_task
CREATE TRIGGER tasks_notify
AFTER INSERT ON tasks
FOR EACH ROW EXECUTE FUNCTION notify_task();