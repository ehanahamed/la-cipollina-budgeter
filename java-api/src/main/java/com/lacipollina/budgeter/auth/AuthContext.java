package com.lacipollina.budgeter.auth;

import java.util.UUID;

public class AuthContext {
    private boolean authed = false;
    private String username;
    
    public AuthContext(boolean authed) {
        this.authed = authed;
    }
    public AuthContext(
        boolean authed,
        String username
    ) {
        this.authed = authed;
        this.username = username;
    }

    public boolean isAuthed() {
        return authed;
    }
    public AuthedUser getUsername() {
        return username;
    }
}
