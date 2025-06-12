package com.lacipollina.budgeter.auth;

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
    public List<String> getAllUsernames() {
        try {
            return jdbcTemplate.query(
                "SELECT name FROM public.usernames",
                usernameRowMapper
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
            );
        }
    }

    public AuthContext authContextUsingToken(String authToken) {
        if (authToken != null && !authToken.isEmpty()) {
            try {
                return jdbcTemplate.queryForObject(
                    "SELECT u.username " +
                    "FROM auth.users u JOIN auth.sessions s ON u.username = s.username " +
                    "WHERE s.token = ? and s.expire_at > (select now())",
                    new Object[] { authToken },
                    (resultSet, rowNum) -> new AuthContext(
                        true,
                        resultSet.getString("username")
                    )
                );
            } catch (EmptyResultDataAccessException e) {
                return new AuthContext(false);
            }
        } else {
            return new AuthContext(false);
        } 
    }
}

