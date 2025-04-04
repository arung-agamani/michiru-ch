import { useQuery } from "@tanstack/react-query";
import { ColumnDef } from "@tanstack/react-table";
import httpClient, { APIResponse } from "../../lib/httpClient.ts";
import { Link } from "react-router-dom";
import DataTable from "../DataTable.tsx";

interface Props {
    projectId: string;
}

interface EventTemplate {
    id: string;
    project_id: string;
    event_type: string;
    template: string;
    description?: string;
    created_at: string;
    updated_at: string;
}

const columns: ColumnDef<EventTemplate>[] = [
    {
        header: "Event Type",
        accessorKey: "event_type",
        cell: ({ row }) => (
            <Link
                to={`/app/projects/${row.original.project_id}/event-templates/${row.original.id}`}
                className="text-blue-700 hover:underline"
            >
                {row.original.event_type}
            </Link>
        ),
    },
    {
        header: "Description",
        accessorKey: "description",
    },
];

const AvailableEvents = [
    "push",
    "pull_request",
    "issue",
    "issue_comment",
] as const;

const EventTemplateMapping: React.FC<Props> = ({ projectId }) => {
    const { data } = useQuery({
        queryKey: ["projects", projectId, "event-templates"],
        queryFn: async () => {
            const res = await httpClient.get<APIResponse<EventTemplate[]>>(
                `projects/${projectId}/templates`
            );
            if (!res.ok) {
                throw new Error("Failed to fetch event templates");
            }
            const data = (await res.json()).data;
            return data;
        },
    });
    return (
        <div className="w-full py-4 gap-y-4 flex flex-col">
            {/* <DataTable columns={columns} data={data || []} contained /> */}
            {AvailableEvents.map((event) => {
                const eventTemplate = data?.find(
                    (template) => template.event_type === event
                );
                if (!eventTemplate) {
                    return (
                        <div key={event} className="">
                            <p className="text-lg ">
                                Event Type:{" "}
                                <span className="font-semibold">{event}</span>
                            </p>
                            <p className="text-sm text-gray-500">
                                No template available for this event type.
                            </p>
                        </div>
                    );
                }
                return (
                    <div key={event}>
                        <p className="text-lg ">
                            Event Type:{" "}
                            <span className="font-semibold">
                                {eventTemplate?.event_type}
                            </span>
                        </p>
                        <p className="text-sm text-gray-500">
                            {eventTemplate?.description}
                        </p>
                        <div className="w-full bg-gray-100 p-4 rounded-md mt-2">
                            <p className="text-sm text-gray-700">
                                {eventTemplate?.template}
                            </p>
                        </div>
                    </div>
                );
            })}
        </div>
    );
};

export default EventTemplateMapping;
