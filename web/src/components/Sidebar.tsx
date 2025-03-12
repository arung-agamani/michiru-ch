import React from "react";
import { Link } from "react-router-dom";

const Sidebar: React.FC = () => {
    return (
        <div className="h-screen w-64 bg-gray-800 text-white flex flex-col">
            <div className="p-4 text-2xl font-bold">My App</div>
            <nav className="flex-1 p-4">
                <ul>
                    <li className="mb-2">
                        <Link
                            to="/app"
                            className="block p-2 rounded hover:bg-gray-700"
                        >
                            Home
                        </Link>
                    </li>
                    <li className="mb-2">
                        <Link
                            to="/app/dashboard"
                            className="block p-2 rounded hover:bg-gray-700"
                        >
                            Dashboard
                        </Link>
                    </li>
                    <li className="mb-2">
                        <Link
                            to="/profile"
                            className="block p-2 rounded hover:bg-gray-700"
                        >
                            Profile
                        </Link>
                    </li>
                    <li className="mb-2">
                        <Link
                            to="/settings"
                            className="block p-2 rounded hover:bg-gray-700"
                        >
                            Settings
                        </Link>
                    </li>
                </ul>
            </nav>
            <div className="p-4">
                <button
                    type="button"
                    className="w-full p-2 bg-red-600 rounded hover:bg-red-700"
                >
                    Logout
                </button>
            </div>
        </div>
    );
};

export default Sidebar;
