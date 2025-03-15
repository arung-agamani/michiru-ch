import { Link, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import { Project } from "./types.ts";

type ProjectResponse = APIResponse<Project>;

const ProjectDetailPage = () => {
    const { projectId } = useParams();
    const { data, isSuccess } = useQuery({
        queryKey: ["projects", projectId],
        queryFn: async () => {
            const res = await httpClient.get<ProjectResponse>(
                `projects/${projectId}`
            );
            if (!res.ok) {
                throw new Error("Failed to fetch project");
            }
            return (await res.json()).data;
        },
    });
    if (!isSuccess) {
        return <div>Loading...</div>;
    }
    return (
        <div className="p-4">
            <nav className="mb-4 text-sm text-gray-500">
                <Link to=".." relative="path">
                    Projects
                </Link>{" "}
                / <span>Project Detail</span>
            </nav>
            <h1 className="text-4xl font-semibold mb-4">
                Project Detail: {data.project_name}
            </h1>
            <div className="bg-white shadow-md rounded p-4">
                <div className="mb-4">
                    <h2 className="text-2xl font-semibold">Description</h2>
                    <p>{data.description}</p>
                </div>
                <div className="mb-4">
                    <h2 className="text-2xl font-semibold">Added By</h2>
                    <p>{data.added_by}</p>
                </div>
                <div className="mb-4">
                    <h2 className="text-2xl font-semibold">Channel ID</h2>
                    <p>{data.channel_id}</p>
                </div>
                <div className="mb-4">
                    <h2 className="text-2xl font-semibold">Created At</h2>
                    <p>{new Date(data.created_at).toLocaleString()}</p>
                </div>
                <div className="mb-4">
                    <h2 className="text-2xl font-semibold">Updated At</h2>
                    <p>{new Date(data.updated_at).toLocaleString()}</p>
                </div>
                {data.webhook_origin && (
                    <div className="mb-4">
                        <h2 className="text-2xl font-semibold">
                            Webhook Origin
                        </h2>
                        <p>{data.webhook_origin}</p>
                    </div>
                )}
                {data.webhook_url && (
                    <div className="mb-4">
                        <h2 className="text-2xl font-semibold">Webhook URL</h2>
                        <p>{data.webhook_url}</p>
                    </div>
                )}
                {data.webhook_secret && (
                    <div className="mb-4">
                        <h2 className="text-2xl font-semibold">
                            Webhook Secret
                        </h2>
                        <p>{data.webhook_secret}</p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default ProjectDetailPage;
