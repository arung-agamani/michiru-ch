import { useQuery } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import { Project } from "./types.ts";
import { Link } from "react-router-dom";
import DataTable from "../../../components/DataTable.tsx";
import { ColumnDef, createColumnHelper } from "@tanstack/react-table";

type ProjectListResponse = APIResponse<Project[]>;

const columns: ColumnDef<Project>[] = [
    {
        header: "Project Name",
        accessorKey: "project_name",
        size: 400,
        cell: (data) => (
            <Link to={`${data.row.original.id}`}>
                {data.row.original.project_name}
            </Link>
        ),
    },
    {
        header: "Description",
        accessorKey: "description",
    },
    // {
    //     header: "Added By",
    //     accessorKey: "added_by",
    // },
    {
        header: "Created At",
        accessorKey: "created_at",
    },
];

const ProjectListPage = () => {
    const { data, isSuccess } = useQuery({
        queryKey: ["projects"],
        queryFn: async () => {
            const response = await httpClient.get<ProjectListResponse>(
                "projects"
            );
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
        <div className="w-full">
            <h2 className="text-2xl font-semibold mb-4">Projects</h2>
            <hr className="my-4" />
            <DataTable columns={columns} data={data} contained />
            <DataTable columns={columns} data={data} />
        </div>
    );
};

export default ProjectListPage;
