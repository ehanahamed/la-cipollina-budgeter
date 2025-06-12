import { PUBLIC_API_URL } from "$env/static/public";
export async function load() {
    let result = {
        authed: false,
        PUBLIC_API_URL: PUBLIC_API_URL
    }
    try {
        if (localStorage.getItem("auth")) {
            const req = await fetch(PUBLIC_API_URL + "/la-cipollina-budgeter/graphql", {
                method: "POST",
                headers: {
                    "Authorization": "Bearer " + 
                }
                body: JSON.stringify({
                    query: "{ amAuthed }"
                })
            })
            if (req) {

            }
        }
    } finally {
        return result;
    }
}

