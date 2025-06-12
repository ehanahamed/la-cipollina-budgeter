import { PUBLIC_API_URL } from "$env/static/public";
export function load() {
    return {
        PUBLIC_API_URL: PUBLIC_API_URL,
        employees: [
            {
                "name": "Carson",
                "type": "floor"
            },
            {
                "name": "Cristian",
                "type": "floor"
            },
            {
                "name": "Renato",
                "type": "kitchen"
            },
            {
                "name": "Manuel",
                "type": "kitchen"
            },
            {
                "name": "Caitlyn",
                "type": "floor"
            },
        ]
    }
}
