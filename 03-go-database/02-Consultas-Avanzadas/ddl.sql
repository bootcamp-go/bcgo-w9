/* ---------------------------- */
-- DDL
DROP DATABASE IF EXISTS empresa_db;

CREATE DATABASE IF NOT EXISTS empresa_db;

USE empresa_db;

CREATE TABLE IF NOT EXISTS departamentos (
	`depto_nro` varchar(7) NOT NULL,
	`nombre_depto` varchar(50) NOT NULL,
    `localidad` varchar(150) NULL,
    PRIMARY KEY (`depto_nro`)
);

CREATE TABLE IF NOT EXISTS empleados (
	`cod_emp` varchar(6) NOT NULL,
    `nombre` varchar(50) NOT NULL,
    `apellido` varchar(50) NOT NULL,
    `puesto` varchar(50) NOT NULL,
    `fecha_alta` date NULL,
    `salario` int NULL,
    `comision` int DEFAULT 0,
    `depto_nro` varchar(7) NULL,
    PRIMARY KEY (`cod_emp`),
    KEY `idx_employee_puesto` (`puesto`),
    KEY `idx_employee_depto_nro` (`depto_nro`),
    CONSTRAINT `fk_employee_depto_nro` FOREIGN KEY (`depto_nro`) REFERENCES `departamentos` (`depto_nro`)
);

/* ---------------------------- */
-- DML
INSERT INTO departamentos (`depto_nro`, `nombre_depto`, `localidad`) VALUES
('D-000-1', 'Software', 'Los Tigres'),
('D-000-2', 'Sistemas', 'Guadalupe'),
('D-000-3', 'Contabilidad', 'La Roca'),
('D-000-4', 'Ventas', 'La Plata'),
('D-000-5', 'Marketing', 'Mar del Plata');


INSERT INTO empleados (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES
('E-0001', 'César', 'Piñero', 'Vendedor', '2018-05-12', 80000, 15000, 'D-000-4'),
('E-0002', 'Yosep', 'Kowaleski', 'Analista', '2015-07-14', 140000, 0, 'D-000-2'),
('E-0003', 'Mariela', 'Barrios', 'Director', '2014-06-05', 185000, 0, 'D-000-3'),
('E-0004', 'Jonathan', 'Aguilera', 'Vendedor', '2015-06-03', 85000, 10000, 'D-000-4'),
('E-0005', 'Daniel', 'Brezezicki', 'Vendedor', '2018-03-03', 83000, 10000, 'D-000-4'),
('E-0006', 'Mito', 'Barchuk', 'Presidente', '2014-06-05', 190000, 0, 'D-000-3'),
('E-0007', 'Emilio', 'Galarza', 'Desarrollador', '2014-08-02', 60000, 0, 'D-000-1');
