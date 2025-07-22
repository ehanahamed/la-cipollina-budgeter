import { PUBLIC_API_URL } from "$env/static/public";

export async function load({ fetch }) {
    if (window.localStorage) {
        let authToken = localStorage.getItem("budgeter:auth");
        if (!authToken) {
            return {
                authed: false,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        }
        try{
            const res = await (
                await fetch(PUBLIC_API_URL + "/users", {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            const authedUserRes = await (
                await fetch(PUBLIC_API_URL + "/authed-user", {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            return {
                authed: true,
                users: res,
                authedUser: authedUserRes,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        } catch (err) {
            console.error("Error while getting users in settings' load func: ", err);
            return {
                authed: false,
                users: true,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        }
    }
}
