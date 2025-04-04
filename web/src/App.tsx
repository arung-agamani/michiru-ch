import { BrowserRouter, Route, Routes } from "react-router-dom";
import IndexPage from "./pages/Index.tsx";
import LoginPage from "./pages/Login.tsx";
import ProtectedRoute from "./components/ProtectedRoute.tsx";
import AppIndexPage from "./pages/App/index.tsx";
import DashboardPage from "./pages/App/Dashboard.tsx";
import ProjectDetailPage from "./pages/App/Project/ProjectDetail.tsx";
import ProjectListPage from "./pages/App/Project/ProjectList.tsx";
import PredefinedTemplatesPage from "./pages/App/PredefinedTemplates/index.tsx";
import AddTemplatePage from "./pages/App/PredefinedTemplates/AddTemplate.tsx";
import TemplateDetailPage from "./pages/App/PredefinedTemplates/TemplateDetail.tsx";
import AddProjectEventTemplate from "./pages/App/Project/AddProjectEventTemplate.tsx";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<IndexPage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/app" element={<ProtectedRoute />}>
                    <Route index element={<AppIndexPage />} />
                    <Route path="dashboard" element={<DashboardPage />} />
                    <Route path="projects" element={<ProjectListPage />} />
                    <Route path="projects/:projectId" element={<ProjectDetailPage />} />
                    <Route path="projects/:projectId/event-templates/add" element={<AddProjectEventTemplate />} />
                    <Route path="predefined-templates" element={<PredefinedTemplatesPage />} />
                    <Route path="predefined-templates/add" element={<AddTemplatePage />} />
                    <Route path="predefined-templates/:id" element={<TemplateDetailPage />} />
                </Route>
            </Routes>
        </BrowserRouter>
    );
}

export default App;
