INSERT INTO todos (title, description, completed, category_id, priority, due_date)
VALUES
  ('Finish Backend Challenge', 'Complete all tasks', false, 1, 'high', NOW() + INTERVAL '2 days'),
  ('Buy groceries', 'Milk, eggs, bread', false, 3, 'medium', NOW() + INTERVAL '1 day'),
  ('Workout', 'Morning jog for 30 mins', false, 4, 'low', NOW() + INTERVAL '3 days');
