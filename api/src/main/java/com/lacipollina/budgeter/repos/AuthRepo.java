package com.lacipollina.budgeter.repos;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.stereotype.Repository;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;

@Repository
public class AuthRepo {
    @Autowired
    private JdbcTemplate jdbcTemplate;
   
    private RowMapper<String> usernameRowMapper = new RowMapper<String>() {
        public String mapRow(ResultSet resultSet, int rowNum) throws SQLException {
            return resultSet.getString("username");
        }
    };
    private RowMapper<String> tokenRowMapper = new RowMapper<String>() {
        public String mapRow(ResultSet resultSet, int rowNum) throws SQLException {
            return resultSet.getString("token");
        }
    };
    public List<String> getAllUsers() {
        try {
            return jdbcTemplate.query(
                "SELECT name FROM public.usernames",
                rowMapper
            );
        } catch (EmptyResultDataAccessException e) {
            return null;
        }
    }
    public String verifyPasswordAndGenAuthToken(String username, String password) {
        String verifiedUsername;
        try {
             verifiedUsername = jdbcTemplate.queryForObject(
                """
                SELECT username FROM auth.users
                WHERE username = ? and encrypted_password = crypt(?, encrypted_password)
                """,
                usernameRowMapper,
                username,
                password
            );
        } catch (EmptyResultDataAccessException e) {
            return null;
        }
        if (verifiedUsername != null) {
            return jdbcTemplate.queryForObject(
                """
                INSERT INTO auth.sessions (username)
                VALUES (?) RETURNING token
                """,
                tokenRowMapper,
                verifiedUsername
            )
        }
    }
    // public String verifyToken(String token) {
    //     jdbcTemplate.queryForObject()
    // }
}

