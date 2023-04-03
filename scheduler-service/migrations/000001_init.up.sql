CREATE TABLE IF NOT EXISTS chats (
   id BIGINT PRIMARY KEY,
   newsletter_interval VARCHAR(80),
   created_at TIMESTAMPTZ DEFAULT now(),
   updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trigger_set_timestamp
   BEFORE UPDATE ON chats
   FOR EACH ROW
   EXECUTE PROCEDURE set_timestamp();