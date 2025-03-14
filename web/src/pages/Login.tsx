import { useLocation } from "react-router-dom";

const LoginPage = () => {
    const location = useLocation();
    const handleGoogleSignIn = () => {
        const currentURI =
            location.state?.from ||
            `${globalThis.window.location.hostname}/app`;
        console.log("Redirecting to:", currentURI);
        globalThis.window.location.href = `/auth/login?redirect_uri=${encodeURIComponent(
            currentURI
        )}`;
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md text-center">
                <h2 className="text-2xl font-bold mb-6">Login</h2>
                <button
                    type="button"
                    onClick={handleGoogleSignIn}
                    className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                >
                    Sign in with Google
                </button>
            </div>
        </div>
    );
};

export default LoginPage;
