



-- Create a function to handle task update
CREATE OR REPLACE FUNCTION handle_task_update()
RETURNS trigger AS $$
BEGIN
    -- Ensure valid state transitions
    IF OLD.state = 'create' AND NEW.state = 'creating' THEN
        RETURN NEW;
    ELSIF OLD.state = 'creating' AND (NEW.state = 'created' OR NEW.state = 'failed') THEN
        RETURN NEW;
    ELSIF OLD.state = 'created' AND NEW.state = 'deleting' THEN
        RETURN NEW;
    ELSIF OLD.state = 'deleting' AND (NEW.state = 'deleted' OR NEW.state = 'failed') THEN
        RETURN NEW;
    ELSE
        RAISE EXCEPTION 'Invalid state transition from % to %', OLD.state, NEW.state;
    END IF;
END;
$$ LANGUAGE plpgsql;



-- Create trigger for update
CREATE TRIGGER trigger_task_update
BEFORE UPDATE ON tasks
FOR EACH ROW
WHEN (OLD.state IS DISTINCT FROM NEW.state)
EXECUTE FUNCTION handle_task_update();
