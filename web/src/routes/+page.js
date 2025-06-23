import { PUBLIC_API_URL } from "$env/static/public";
export async function load() {
    let result = {
        authed: false,
        PUBLIC_API_URL: PUBLIC_API_URL
    }
    try {
        // const authToken = localStorage.getItem("auth");
        // if (authToken) {
        //     const req = await fetch(PUBLIC_API_URL + "/users/", {
        //         method: "POST",
        //         headers: {
        //             "Authorization": "Bearer " + authToken
        //         },
        //         body: JSON.stringify({
        //             query: "{ amAuthed }"
        //         })
        //     })
        //     if (req) {
        // 
        //     }
        // }
    } finally {
        return result;
    }
}
