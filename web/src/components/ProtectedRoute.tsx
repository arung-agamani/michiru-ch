import React, { useEffect } from "react";
import { Navigate, Outlet, useLocation, useNavigate } from "react-router-dom";
import { useAtom } from "jotai";
import { authStateAtom } from "../state/auth.ts";
import Sidebar from "./Sidebar.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

const queryClient = new QueryClient();

const ProtectedRoute: React.FC = () => {
    const [authState, setAuthState] = useAtom(authStateAtom);
    const location = useLocation();
    const navigate = useNavigate();

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const response = await fetch("/auth/me", {
                    redirect: "manual",
                });
                if (response.ok) {
                    // console.log(response);
                    const userData = await response.json();
                    const expiry = Date.now() + 3600 * 1000; // 1 hour expiry
                    setAuthState({ user: userData, expiry });
                } else {
                    setAuthState({ user: null, expiry: null });
                    navigate("/login", {
                        state: {
                            from:
                                globalThis.window.location.protocol +
                                "//" +
                                globalThis.window.location.host +
                                location.pathname +
                                location.search,
                        },
                    });
                }
            } catch (error) {
                console.error("Failed to fetch user data:", error);
                setAuthState({ user: null, expiry: null });
                navigate("/login", {
                    state: {
                        from:
                            globalThis.window.location.protocol +
                            "//" +
                            globalThis.window.location.host +
                            location.pathname +
                            location.search,
                    },
                });
            }
        };
        if (
            authState.user === null ||
            (authState.expiry !== null && authState.expiry < Date.now())
        ) {
            checkAuth();
        }
    }, [location]);

    if (authState.user === null) {
        return <div>Loading...</div>;
    }

    return (
        <QueryClientProvider client={queryClient}>
            <div className="flex">
                <Sidebar />
                <div className="main-wrapper p-4 bg-gray-50 w-full">
                    <Outlet />
                </div>
            </div>
            <ReactQueryDevtools />
        </QueryClientProvider>
    );
};

export default ProtectedRoute;
