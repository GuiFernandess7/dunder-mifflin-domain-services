# Dunder Mifflin Domain Services
Repository containing the business logic layer for a fictional application handling data from a local PostgreSQL database based on Dunder Mifflin, intended for educational purposes involving PostgreSQL, Golang, and SQLC.
<hr>

## Database Tables:

<img src="https://github.com/GuiFernandess7/dunder-mifflin-domain-services/assets/63022500/48a71f19-a413-4754-b091-ed78076079d7" alt="company-database" width="500" />

<hr>

### Some of the queries

```
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
```