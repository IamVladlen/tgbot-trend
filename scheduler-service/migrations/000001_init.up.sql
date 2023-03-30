CREATE TABLE IF NOT EXISTS chat (
   id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
   chat_id BIGINT,
   newsletter_interval VARCHAR(80),
   created_at TIMESTAMPTZ DEFAULT now(),
   updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER IF NOT EXISTS set_timestamp
   BEFORE UPDATE ON chat
   FOR EACH ROW
   EXECUTE PROCEDURE trigger_set_timestamp();