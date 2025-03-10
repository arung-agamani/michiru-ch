import ky from "ky";

const httpClient = ky.create({
    prefixUrl: "/api/v1",
    headers: {
        "Content-Type": "application/json",
    },
    timeout: 10000,
    hooks: {
        beforeRequest: [
            (req) => {
                const token = localStorage.getItem("authToken");
                if (token) {
                    req.headers.set("Authorization", `Bearer ${token}`);
                }
            },
        ],
    },
});

export default httpClient;
