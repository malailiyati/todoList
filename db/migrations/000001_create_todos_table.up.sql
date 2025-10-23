CREATE TABLE public.todos (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  description TEXT,
  completed BOOLEAN DEFAULT FALSE,
  category_id INT REFERENCES categories(id) ON DELETE SET NULL,
  priority VARCHAR(10),
  due_date TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);