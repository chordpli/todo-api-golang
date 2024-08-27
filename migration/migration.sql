create table todos
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255)                              NOT NULL,
    description TEXT,
    status      enum ('PENDING', 'COMPLETED', 'PROGRESS') NOT NULL,
    created_at  DATETIME                                  NOT NULL,
    updated_at  DATETIME                                  NOT NULL,
    deleted_at  DATETIME
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;