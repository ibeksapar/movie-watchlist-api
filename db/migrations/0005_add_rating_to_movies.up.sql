DO $$ 
BEGIN 
   IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='movies' AND column_name='rating') THEN 
      ALTER TABLE movies ADD COLUMN rating FLOAT; 
   END IF; 
END $$;
