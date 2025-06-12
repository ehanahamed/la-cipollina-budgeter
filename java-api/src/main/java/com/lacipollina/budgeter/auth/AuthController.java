package com.lacipollina.budgeter;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.SchemaMapping;
import org.springframework.stereotype.Controller;

import graphql.schema.DataFetchingEnvironment;

import java.util.List;

@Controller
public class AuthController {
    @Autowired
    private AuthRepo authRepo;

    @QueryMapping
    public List<String> allUsernames() {
        return authRepo.getAllUsernames();
    }

    @MutationMapping
    public boolean signIn(
        @Argument String username,
        @Argument String password
    ) {
        /* returns token */
        return authRepo.verifyPasswordAndGenAuthToken(
            username,
            password
        );
    }
}
