-- MySQL Schema
CREATE TABLE sections (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL COMMENT 'Human-readable section identifier',
    section_type VARCHAR(50) NOT NULL COMMENT 'Type of section (hero, features, etc)',
    display_order INT DEFAULT 0 COMMENT 'Sorting order',
    is_active BOOLEAN DEFAULT TRUE COMMENT 'Visibility control',
    details JSON NOT NULL COMMENT 'Flexible content storage',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_section_type (section_type),
    INDEX idx_is_active (is_active),
    INDEX idx_display_order (display_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;