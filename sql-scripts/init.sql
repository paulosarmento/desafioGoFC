-- Criação do banco de dados se ainda não existir
CREATE DATABASE IF NOT EXISTS routesdb;

-- Usar o banco de dados
USE routesdb;

-- Criação da tabela routes
CREATE TABLE IF NOT EXISTS routes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    source_lat DOUBLE,
    source_lng DOUBLE,
    destination_lat DOUBLE,
    destination_lng DOUBLE
);
