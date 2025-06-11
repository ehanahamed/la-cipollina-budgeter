import { PUBLIC_API_URL } from "$env/static/public";

export async function load() {
    if (window.localStorage) {
        let token = localStorage.getItem("auth");
        if (token) {
            fetch(PUBLIC_API_URL + "/graphql", {
                method: "POST",
                headers: {
                    "Authorization": 
                }
                body: JSON.stringify({

                })
            })
        } else {
            alert("not logged in")
        }
    } else {
        alert("Error: Couldn't load token from localStorage");
    }
}
