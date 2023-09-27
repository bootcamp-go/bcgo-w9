USE empresa_db;

/* --------------------------------------------------------------------------------------------------------------------- */
/* Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores. */
SELECT e.nombre, e.puesto, d.localidad, d.nombre_depto FROM empleados as e INNER JOIN departamentos as d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = "Ventas";


/* --------------------------------------------------------------------------------------------------------------------- */
/* Visualizar los departamentos con más de cinco empleados. */
SELECT d.depto_nro, d.nombre_depto, d.localidad FROM departamentos as d INNER JOIN empleados as e ON d.depto_nro = e.depto_nro
GROUP BY d.depto_nro HAVING COUNT(e.depto_nro) > 1;

SELECT COUNT(e.depto_nro) as cantidad, d.depto_nro, d.nombre_depto, d.localidad FROM departamentos as d LEFT JOIN empleados as e ON d.depto_nro = e.depto_nro
GROUP BY d.depto_nro;

/* --------------------------------------------------------------------------------------------------------------------- */
/* Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’. */
SELECT e.nombre, e.salario, d.nombre_depto FROM empleados as e INNER JOIN departamentos as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (
	SELECT e2.puesto FROM empleados as e2 WHERE e2.nombre = "Jonathan" AND e2.apellido = "Aguilera"
);


/* --------------------------------------------------------------------------------------------------------------------- */
/* Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre. */
SELECT e.* FROM empleados as e WHERE e.depto_nro = (
	SELECT d.depto_nro FROM departamentos as d WHERE d.nombre_depto = "Contabilidad"
) ORDER BY e.nombre;


/* --------------------------------------------------------------------------------------------------------------------- */
/* Mostrar el nombre del empleado que tiene el salario más bajo. */
SELECT e.nombre FROM empleados as e
WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleados as e2
);


/* --------------------------------------------------------------------------------------------------------------------- */
/* Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’. */
SELECT e.* FROM empleados as e INNER JOIN departamentos as d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto = "Ventas" ORDER BY e.salario DESC LIMIT 1;

SELECT e.* FROM empleados as e WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleados as e2 INNER JOIN departamentos as d ON d.depto_nro = e2.depto_nro
    WHERE d.nombre_depto = "Ventas"
);
