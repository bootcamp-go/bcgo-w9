CREATE USER IF NOT EXISTS 'user_test'@'localhost' IDENTIFIED BY 'user_pass';

GRANT ALL PRIVILEGES ON db_app_test.* TO 'user_test'@'localhost';

FLUSH PRIVILEGES;