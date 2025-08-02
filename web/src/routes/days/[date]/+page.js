import { PUBLIC_API_URL } from "$env/static/public";
export async function load({ fetch, params, url }) {
    if (window.localStorage) {
        let authToken = localStorage.getItem("budgeter:auth");
        try{
            const weekRes = await (
                await fetch(`${PUBLIC_API_URL}/weeks/${params.date}`, {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken,
                        "Eh-Share-Link-Token": url.searchParams.get("s")
                    }
                })
            ).json();
            const daysRes = await (
                await fetch(`${PUBLIC_API_URL}/weeks/${params.date}/days`, {
                    method: "GET",
                    headers: {
                        "Authorization": "Bearer " + authToken,
                        "Eh-Share-Link-Token": url.searchParams.get("s")
                    }
                })
            ).json();
            return {
                authed: true,
                PUBLIC_API_URL: PUBLIC_API_URL,
                date: params.date,
                week: weekRes,
                days: daysRes,
                usingShareLink: authToken == null && url.searchParams.has("s")
            };
        } catch (err) {
            console.error("Error while getting data: ", err);
            return {
                authed: false,
                error: true,
                PUBLIC_API_URL: PUBLIC_API_URL,
                date: params.date,
                usingShareLink: authToken == null && url.searchParams.has("s")
            };
        }
    }
}
