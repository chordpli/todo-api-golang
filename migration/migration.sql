
create table todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status enum('PENDING', 'COMPLETED', 'PROGRESS') NOT NULL,
)