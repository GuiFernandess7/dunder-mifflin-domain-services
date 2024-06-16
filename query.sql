-- name: ListAllEmployees :many
-- Get all employees including their branches
SELECT *
FROM employee
LEFT JOIN branch
ON employee.emp_id = branch.mgr_id;

-- name: ListAllEmployeesBy :many
-- Get All employees by limit in order desc
SELECT *
FROM employee
LEFT JOIN branch
ON employee.emp_id = branch.mgr_id
ORDER BY emp_id DESC
LIMIT $1;

-- name: ListSalariesASC :many
SELECT *
FROM employee
ORDER BY salary ASC;

-- name: ListSalariesDESC :many
SELECT *
FROM employee
ORDER BY salary DESC;

-- name: FindEmployeeByName :many
SELECT * FROM employee WHERE first_name = @n;

-- name: FilterEmployeeBySalary :many
SELECT *
FROM employee
WHERE salary >= @min AND salary < @max;

-- name: FilterEmployeeBySex :many
SELECT *
FROM employee
WHERE sex = @sex;

-- name: FilterEmployeeBySalaryAndSex :many
SELECT *
FROM employee
WHERE salary >= @min AND salary < @max
AND (@sex::VARCHAR IS NULL OR sex = @sex::VARCHAR);

-- name: EmployeesCount :one
SELECT COUNT(emp_id)
FROM employee;

-- name: GetEmployeeByBranch :many
SELECT employee.emp_id, employee.first_name, employee.last_name, branch.branch_name
FROM employee
JOIN branch
ON branch.branch_id = employee.branch_id
WHERE branch.branch_name = @b_name::VARCHAR;

-- name: ListManagers :many
SELECT employee.emp_id, employee.first_name, branch.branch_name
FROM employee
JOIN branch
ON employee.emp_id = branch.mgr_id;

-- name: GetClientbyBranch :many
SELECT client.client_id, client_name, branch.branch_name
FROM client
JOIN branch
ON branch.branch_id = client.branch_id;

-- name: GetEmployeeBySalesQtd :many
SELECT employee.first_name, employee.last_name
FROM employee
WHERE employee.emp_id IN (
    SELECT works_with.emp_id
    FROM works_with
    WHERE works_with.total_sales > @min::INT
    AND works_with.total_sales < @max::INT
);

-- name: GetClientsByEmployee :many
SELECT c.client_name
FROM client c
JOIN branch b ON c.branch_id = b.branch_id
JOIN employee e ON b.mgr_id = e.emp_id
WHERE e.first_name = @first_name::VARCHAR;

-- name: GetSupplierByType :many
SELECT branch_supplier.supplier_name, branch.branch_name, branch_supplier.supply_type
FROM branch_supplier
JOIN branch
ON branch_supplier.branch_id = branch.branch_id
WHERE branch_supplier.supply_type = @type::VARCHAR;
