import { PUBLIC_API_URL } from "$env/static/public";
export async function load({ fetch, params }) {
    if (window.localStorage) {
        let authToken = localStorage.getItem("budgeter:auth");
        if (!authToken) {
            return {
                authed: false,
                PUBLIC_API_URL: PUBLIC_API_URL,
                date: params.date
            };
        }
        try{
            const res = await (
                await fetch(`${PUBLIC_API_URL}/days/${params.date}`, {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken
                    }
                })
            ).json();
            return {
                authed: true,
                PUBLIC_API_URL: PUBLIC_API_URL,
                date: params.date,
                dayData: res
            };
        } catch (err) {
            console.error("Error while getting day data: ", err);
            return {
                authed: false,
                error: true,
                PUBLIC_API_URL: PUBLIC_API_URL,
                date: params.date
            };
        }
    }
}
