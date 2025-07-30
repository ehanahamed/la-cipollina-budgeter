import { PUBLIC_API_URL } from "$env/static/public";
export async function load({ fetch, url }) {
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
                await fetch(PUBLIC_API_URL + "/weeks", {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            return {
                authed: true,
                weeks: res,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        } catch (err) {
            console.error("Error while getting weeks in reports' load func: ", err);
            return {
                authed: false,
                error: true,
                PUBLIC_API_URL: PUBLIC_API_URL
            };
        }
    }
}
