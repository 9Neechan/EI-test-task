-- Вставляем 10 сервисов
INSERT INTO services (name, description) VALUES
('AI Chatbot', 'Conversational AI assistant'),
('Image Recognition', 'AI-powered image analysis'),
('Speech to Text', 'Convert speech into text'),
('Text to Speech', 'Convert text into natural voice'),
('Machine Translation', 'AI-driven language translation'),
('Sentiment Analysis', 'Analyze emotions in text'),
('Recommendation System', 'Personalized content suggestions'),
('Fraud Detection', 'Detect anomalies in transactions'),
('Chat Moderation', 'Automated chat filtering'),
('Stock Prediction', 'AI-based stock price forecasting');

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
