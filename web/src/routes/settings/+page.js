import { PUBLIC_API_URL } from "$env/static/public";

export async function load() {
    if (window.localStorage) {
        let authToken = localStorage.getItem("auth");
        if (!authToken) {
            return {
                authed: false
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
            return {
                authed: true,
                employees: res
            };
        } catch (err) {
            console.error("Error while getting employees in settings' load func: ", err);
            return {
                authed: false,
                error: true
            };
        }
    }
}
