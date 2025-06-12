package com.lacipollina.budgeter.repos;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.stereotype.Repository;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import com.fasterxml.jackson.databind.ObjectMapper;

@Repository
public class DayReportRepo {
    @Autowired
    private JdbcTemplate jdbcTemplate;
   
    private RowMapper<String> usernameRowMapper = new RowMapper<String>() {
        public String mapRow(ResultSet resultSet, int rowNum) throws SQLException {
            return resultSet.getString("username");
        }
    };
    public Employee addDayInput(Employee employee) {
        try {
            return jdbcTemplate.queryForObject(
                """
                INSERT INTO employees (name, wage, type) VALUES
                (?, ?, ?) RETURNING name, wage, type
                """,
                employeeRowMapper,
                employee.getName(),
                employee.getWage(),
                employee.getType()
            );
        } catch (EmptyResultDataAccessException e) {
            return null;
        }
    }
    public Employee updateEmployee(String name, Employee employee) {
        try {
            return jdbcTemplate.queryForObject(
                """
                UPDATE employees SET name = ?, wage = ?, type = ?
                WHERE name = ?
                RETURNING name, wage, type
                """,
                employeeRowMapper,
                employee.getName(),
                employee.getWage(),
                employee.getType(),
                name
            );
        } catch (EmptyResultDataAccessException e) {
            return null;
        }
    }
    public boolean deleteEmployee(String name) {
        return jdbcTemplate.update(
            """
            DELETE FROM employees
            WHERE name = ?
            """,
            name
        ) > 0;
    }
}

