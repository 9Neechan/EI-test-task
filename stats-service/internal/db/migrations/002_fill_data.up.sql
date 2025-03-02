-- Вставляем 10 сервисов
INSERT INTO services (name, description, price) VALUES
('AI Chatbot', 'Conversational AI assistant', 0.05),
('Image Recognition', 'AI-powered image analysis', 0.10),
('Speech to Text', 'Convert speech into text', 0.08),
('Text to Speech', 'Convert text into natural voice', 0.07),
('Machine Translation', 'AI-driven language translation', 0.06),
('Sentiment Analysis', 'Analyze emotions in text', 0.04),
('Recommendation System', 'Personalized content suggestions', 0.09),
('Fraud Detection', 'Detect anomalies in transactions', 0.12),
('Chat Moderation', 'Automated chat filtering', 0.03),
('Stock Prediction', 'AI-based stock price forecasting', 0.15);

-- Вставляем 10 пользователей
INSERT INTO users (name) VALUES
('Alice Johnson'),
('Bob Smith'),
('Charlie Brown'),
('David White'),
('Emma Stone'),
('Frank Black'),
('Grace Lee'),
('Hannah Scott'),
('Ian Wright'),
('Julia Adams');

-- Вставляем 10 записей в stats (случайные user_id и service_id)
INSERT INTO stats (user_id, service_id, count) VALUES
(1, 3, 5),
(2, 5, 10),
(3, 1, 7),
(4, 8, 3),
(5, 2, 12),
(6, 7, 9),
(7, 10, 4),
(8, 6, 6),
(9, 4, 8),
(10, 9, 11);
