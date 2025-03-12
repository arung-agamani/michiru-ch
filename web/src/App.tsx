import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "./pages/Index.tsx";
import LoginPage from "./pages/Login.tsx";
import ProtectedRoute from "./components/ProtectedRoute.tsx";
import AppIndexPage from "./pages/App/index.tsx";
import DashboardPage from "./pages/App/Dashboard.tsx";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<IndexPage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/app" element={<ProtectedRoute />}>
                    <Route index element={<AppIndexPage />} />
                    <Route path="dashboard" element={<DashboardPage />} />
                </Route>
            </Routes>
        </BrowserRouter>
    );
}

export default App;
