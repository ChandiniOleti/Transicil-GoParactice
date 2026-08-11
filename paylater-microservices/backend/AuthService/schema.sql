-- Auth Service schema for SQLC (shared paylaterdb).
-- Tables used for admin ownership + temporary login lookups.

CREATE TABLE admins (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    credit_limit DECIMAL(10,2) NOT NULL DEFAULT 2000.00,
    current_due DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

CREATE TABLE merchants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    merchant_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(15) UNIQUE NOT NULL,
    commission DECIMAL(5,2) NOT NULL
);
