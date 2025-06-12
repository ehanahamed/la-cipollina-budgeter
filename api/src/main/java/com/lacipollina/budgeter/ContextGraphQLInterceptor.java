package com.lacipollina.budgeter;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.server.WebGraphQlInterceptor;
import org.springframework.graphql.server.WebGraphQlRequest;
import org.springframework.graphql.server.WebGraphQlResponse;
import org.springframework.stereotype.Component;
import reactor.core.publisher.Mono;
import org.springframework.http.HttpCookie;
import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;
import com.lacipollina.budgeter.auth.AuthRepo;

@Component
public class ContextGraphQLInterceptor implements WebGraphQlInterceptor {
    @Autowired
    private AuthRepo authRepo;

    @Override
    public Mono<WebGraphQlResponse> intercept(WebGraphQlRequest request, Chain chain) {
        String authToken = null;

        String authHeader = request.getHeaders().getFirst("Authorization");
        if (
            authHeader != null &&
            authHeader.length() > 7 &&
            authHeader.substring(
                0, 7 /* "bearer " is first 7 characters */
            ).toLowerCase().equals("bearer ")
        ) {
            authToken = authHeader.substring(7);
        }

        String finalAuthToken = authToken;
        request.configureExecutionInput((executionInput, builder) ->
            builder.graphQLContext(contextBuilder ->
                contextBuilder.of("authContext", authRepo.authContextUsingToken(finalAuthToken))
            ).build()
        );

        return chain.next(request);
    }
}

