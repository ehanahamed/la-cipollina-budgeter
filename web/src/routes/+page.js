import { PUBLIC_API_URL } from "$env/static/public";
export async function load({ fetch }) {
    let authToken = localStorage.getItem("budgeter:auth");
    if (!authToken) {
        return {
            authed: false,
            PUBLIC_API_URL: PUBLIC_API_URL
        };
    }
    try{
        const res = await (
            await fetch(PUBLIC_API_URL + "/authed-user", {
                method: "GET",
                headers: {
                    "Authorization": "Bearer " + authToken
                }
            })
        ).json();
        if (res?.id != null) {
            return {
                authed: true,
                authedUser: res,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        } else {
            return {
                authed: false,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        }
    } catch (err) {
        console.error("Error while getting users in settings' load func: ", err);
        return {
            authed: false,
            users: true,
            PUBLIC_API_URL: PUBLIC_API_URL
        };
    }
}
