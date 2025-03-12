import { useQuery } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../lib/httpClient.ts";

interface Project {
    id: string;
    project_name: string;
    description: string;
    channel_id: string;
    added_by: string;
    created_at: string;
    updated_at: string;
    webhook_origin?: string;
    webhook_url?: string;
    webhook_secret?: string;
}

type ProjectResponse = APIResponse<Project[]>;

const AppIndexPage = () => {
    const { data, isSuccess } = useQuery({
        queryKey: ["projects"],
        queryFn: async () => {
            const response = await httpClient.get<ProjectResponse>("projects");
            if (!response.ok) {
                throw new Error("Failed to fetch projects");
            }
            return (await response.json()).data;
        },
    });
    if (!isSuccess) {
        return <div>Loading...</div>;
    }
    return (
        <p>
            <table className="min-w-full bg-white border border-gray-200">
                <thead>
                    <tr className="bg-gray-100 border-b">
                        <th className="py-2 px-4 text-left">Project Name</th>
                        <th className="py-2 px-4 text-left">Description</th>
                        <th className="py-2 px-4 text-left">Added By</th>
                        <th className="py-2 px-4 text-left">Created At</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((project) => (
                        <tr
                            key={project.id}
                            className="border-b hover:bg-gray-50"
                        >
                            <td className="py-2 px-4">
                                {project.project_name}
                            </td>
                            <td className="py-2 px-4">{project.description}</td>
                            <td className="py-2 px-4">{project.added_by}</td>
                            <td className="py-2 px-4">{project.created_at}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </p>
    );
};

export default AppIndexPage;
