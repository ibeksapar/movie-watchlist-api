DO $$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='genres' AND column_name='description') THEN
      ALTER TABLE genres ADD COLUMN description TEXT;
   END IF;
END $$;