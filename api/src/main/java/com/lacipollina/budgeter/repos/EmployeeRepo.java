package com.lacipollina.budgeter.repos;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.stereotype.Repository;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.UUID;
import com.lacipollina.budgeter.Employee;

@Repository
public class EmployeeRepo {
    @Autowired
    private JdbcTemplate jdbcTemplate;
   
    private RowMapper<Employee> employeeRowMapper = new RowMapper<Employee>() {
        public Employee mapRow(ResultSet resultSet, int rowNum) throws SQLException {
            return new Employee(
                resultSet.getString("name"),
                resultSet.getDouble("wage"),
                resultSet.getString("type")
            );
        }
    };
    public List<Employee> getAllEmployees() {
        try {
            return jdbcTemplate.query(
                "SELECT name, wage, type FROM employees",
                employeeRowMapper
            );
        } catch (EmptyResultDataAccessException e) {
            return null;
        }
    }
}

