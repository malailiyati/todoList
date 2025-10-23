INSERT INTO categories (name, color)
VALUES
  ('Work', '#FF0000'),      
  ('Personal', '#0000FF'),  
  ('Shopping', '#00FF00'),  
  ('Fitness', '#FFA500')    
ON CONFLICT (name) DO NOTHING;
