import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "./pages/Index.tsx";
import LoginPage from "./pages/Login.tsx";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<IndexPage />} />
                <Route path="/login" element={<LoginPage />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
