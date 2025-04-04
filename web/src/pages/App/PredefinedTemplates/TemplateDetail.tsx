import { useParams, Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";

interface PredefinedTemplate {
    id: string;
    event_type: string;
    template: string;
    description?: string;
    created_at: string;
    updated_at: string;
}

const TemplateDetailPage = () => {
    const { id } = useParams();
    const { data: template, isSuccess } = useQuery({
        queryKey: ["predefinedTemplates", id],
        queryFn: async () => {
            const response = await httpClient.get<
                APIResponse<PredefinedTemplate>
            >(`predefined-templates/${id}`);
            if (!response.ok) {
                throw new Error("Failed to fetch template");
            }
            return (await response.json()).data;
        },
    });

    if (!isSuccess) {
        return <div>Loading...</div>;
    }
    return (
        <div className="p-4 bg-white shadow-md rounded">
            <nav className="mb-4 text-sm text-gray-500">
                <Link
                    to="/app/predefined-templates"
                    className="hover:underline"
                >
                    Predefined Templates
                </Link>{" "}
                / <span>Template Detail</span>
            </nav>
            <h1 className="text-2xl font-semibold mb-4">Template Detail</h1>
            <div className="mb-4">
                <p>
                    <strong>Event Type:</strong> {template.event_type}
                </p>
                <p>
                    <strong>Description:</strong> {template.description}
                </p>
                <p>
                    <strong>Template Content:</strong> {template.template}
                </p>
            </div>
        </div>
    );
};

export default TemplateDetailPage;
