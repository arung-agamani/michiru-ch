import React, { useEffect } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useAtom } from "jotai";
import { authStateAtom } from "../state/auth.ts";

const ProtectedRoute: React.FC = () => {
    const [authState, setAuthState] = useAtom(authStateAtom);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const response = await fetch("/auth/me");
                if (response.ok) {
                    const userData = await response.json();
                    const expiry = Date.now() + 3600 * 1000; // 1 hour expiry
                    setAuthState({ user: userData, expiry });
                } else {
                    setAuthState({ user: null, expiry: null });
                }
            } catch (error) {
                console.error("Failed to fetch user data:", error);
                setAuthState({ user: null, expiry: null });
            }
        };

        if (
            authState.user === null ||
            (authState.expiry !== null && authState.expiry < Date.now())
        ) {
            checkAuth();
        }
    }, [authState, setAuthState]);

    if (authState === null) {
        return <div>Loading...</div>;
    }

    return authState ? (
        <>
            <Outlet />
        </>
    ) : (
        <Navigate to="/login" />
    );
};

export default ProtectedRoute;
