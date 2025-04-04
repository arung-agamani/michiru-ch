import { Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import DataTable from "../../../components/DataTable.tsx";
import { ColumnDef } from "@tanstack/react-table";

interface PredefinedTemplate {
    id: string;
    event_type: string;
    template: string;
    description?: string;
    created_at: string;
    updated_at: string;
}

const PredefinedTemplatesPage = () => {
    const { data: templates = [], isLoading } = useQuery({
        queryKey: ["predefinedTemplates"],
        queryFn: async () => {
            const response = await httpClient.get<
                APIResponse<PredefinedTemplate[]>
            >("predefined-templates");
            if (!response.ok) {
                throw new Error("Failed to fetch templates");
            }
            return (await response.json()).data;
        },
    });

    const columns: ColumnDef<PredefinedTemplate>[] = [
        {
            header: "Description",
            accessorKey: "description",
        },
        {
            header: "Template Name",
            accessorKey: "name",
            cell: ({ row }) => (
                <Link
                    to={`/app/predefined-templates/${row.original.id}`}
                    className="text-blue-700 hover:underline"
                >
                    {row.original.event_type}
                </Link>
            ),
        },
    ];

    return (
        <div className="w-full p-4">
            <div className="flex mb-4">
                <h2 className="text-2xl font-semibold">Predefined Templates</h2>
                <Link
                    to="/app/predefined-templates/add"
                    className="rounded-xl p-2 text-white font-bold hover:bg-blue-500 bg-blue-400 ml-auto hover:cursor-pointer"
                >
                    Add Template
                </Link>
            </div>
            <hr className="my-4" />
            {isLoading ? (
                <p>Loading...</p>
            ) : (
                <DataTable columns={columns} data={templates} contained />
            )}
        </div>
    );
};

export default PredefinedTemplatesPage;
