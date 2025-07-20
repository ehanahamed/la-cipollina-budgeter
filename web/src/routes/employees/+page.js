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
                await fetch(PUBLIC_API_URL + "/employees", {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            const valentinosRes = await (
                await fetch(PUBLIC_API_URL + "/valentinos", {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            return {
                authed: true,
                employees: res,
                valentinos: valentinosRes,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        } catch (err) {
            console.error("Error while getting employees in settings' load func: ", err);
            return {
                authed: false,
                error: true,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        }
    }
}
