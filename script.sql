CREATE DATABASE IF NOT EXISTS tuticketdb;

USE tuticketdb;

CREATE TABLE evento (
	id INT NOT NULL AUTO_INCREMENT,
	nombre VARCHAR(25) NULL,
	descripcion VARCHAR(50) NULL,
	fecha DATE NULL,
	sede_id INT NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (sede_id) REFERENCES sede(id)
);

CREATE TABLE sede (
	id INT NOT NULL AUTO_INCREMENT,
	nombre VARCHAR(50) NULL,
	tipo VARCHAR(25) NULL,
	ciudad VARCHAR(25) NULL,
	pais VARCHAR(25) NULL,
	direccion VARCHAR(50) NULL,

	PRIMARY KEY(id)
);

CREATE TABLE asiento (
	id INT NOT NULL AUTO_INCREMENT,
	numero INT NULL,
	disponible BOOlEAN NULL,
	evento_id INT NULL,
	localidad_id INT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (evento_id) REFERENCES evento(id),
	FOREIGN KEY (localidad_id) REFERENCES localidad(id)
);

CREATE TABLE localidad (
	id INT NOT NULL AUTO_INCREMENT,
	nombre VARCHAR(25) NULL,
	precio DECIMAL(5,2) NULL,

	PRIMARY KEY(id),
);

CREATE TABLE boleto (
	id INT NOT NULL AUTO_INCREMENT,
	asiento_id INT NULL,
	usuario_id INT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (asiento_id) REFERENCES asiento(id),
	FOREIGN KEY (usuario_id) REFERENCES usuario(id)
);

CREATE TABLE usuario (
	id INT NOT NULL AUTO_INCREMENT,
	usuario VARCHAR(25) NULL,
	nombre VARCHAR(25) NULL,
	contrasena VARCHAR(25) NULL,

	PRIMARY KEY(id)
);

INSERT INTO evento (id, nombre, descripcion, fecha, sede_id) VALUES
	(1, 'Aerosmith', 'Concierto de Aerosmith', '2019-06-06', 1);

INSERT INTO sede (id, nombre, tipo, ciudad, pais, direccion) VALUES
	(1, 'Estadio Monumental', 'Estadio', 'Guayaquil', 'Ecuador', 'Av. Barcelona');

INSERT INTO asiento (id, numero, disponible, evento_id, localidad_id) VALUES
	(1, 1, false, 1, 1),
	(2, 2, false, 1, 1),
	(3, 3, true, 1, 1),
	(4, 4, true, 1, 1),
	(5, 1, true, 1, 2),
	(6, 2, true, 1, 2),
	(7, 3, true, 1, 2),
	(8, 4, true, 1, 2);

INSERT INTO localidad (id, nombre, precio) VALUES
	(1, 'Palco', '100,00'),
	(2, 'Tribuna', '50,00');

INSERT INTO boleto (id, asiento_id, usuario_id) VALUES
	(1, 1, 1),
	(2, 2, 1);

INSERT INTO usuario (id, usuario, nombre, contrasena) VALUES
	(1, 'user1', 'Allan Alarcón', '12345'),
	(2, 'user2', 'Juan Pérez', '12345');