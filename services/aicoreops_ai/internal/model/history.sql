CREATE TABLE IF NOT EXISTS history (
    id bigint AUTO_INCREMENT,
    session_id varchar(255) NOT NULL,
    question text NOT NULL,
    answer text NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'history table';

# goctl model mysql ddl --src history.sql --dir .