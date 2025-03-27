import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import { Project } from "./types.ts";
import { Link } from "react-router-dom";
import DataTable from "../../../components/DataTable.tsx";
import { ColumnDef, createColumnHelper } from "@tanstack/react-table";
import { Drawer } from "../../../components/Drawer.tsx";
import { TextField } from "../../../components/TextField.tsx";
import { useState } from "react";
import { useForm } from "react-hook-form";
import TextArea from "../../../components/TextArea.tsx";
import { useAtom } from "jotai";
import { authStateAtom } from "../../../state/auth.ts";

type ProjectListResponse = APIResponse<Project[]>;

const columns: ColumnDef<Project>[] = [
    {
        header: "Project Name",
        accessorKey: "project_name",
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
    const [authState] = useAtom(authStateAtom);
    const queryClient = useQueryClient();
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
    const drawerForm = useForm();
    const [drawerOpen, setDrawerOpen] = useState(false);

    const projectMutation = useMutation({
        mutationFn: async (data: any) => {
            const response = await httpClient.post<APIResponse<Project>>(
                "projects",
                {
                    json: data,
                }
            );
            if (!response.ok) {
                throw new Error("Failed to create project");
            }
            return (await response.json()).data;
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["projects"] });
        },
    });
    if (!isSuccess) {
        return <div>Loading...</div>;
    }

    const { register, handleSubmit } = drawerForm;

    const onSubmit = async (data: any) => {
        const payload = { ...data, added_by: authState.user?.email };
        console.log(payload);
        await projectMutation.mutateAsync(payload);
    };
    return (
        <div className="w-full">
            <div className="flex mb-4">
                <h2 className="text-2xl font-semibold">Projects</h2>
                <button
                    onClick={() => setDrawerOpen(true)}
                    type="button"
                    className="rounded-xl p-2 text-white font-bold hover:bg-blue-500 bg-blue-400 ml-auto hover:cursor-pointer"
                >
                    Create Project
                </button>
            </div>
            <hr className="my-4" />
            <DataTable columns={columns} data={data} contained />
            <Drawer isOpen={drawerOpen} onDismiss={() => setDrawerOpen(false)}>
                <div className="bg-white shadow-md p-4 h-full max-h-dvh overflow-y-auto min-w-xl">
                    <h2 className="text-2xl font-semibold">Create Project</h2>
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="mb-4">
                            <TextField
                                label="Project Name"
                                register={register("project_name")}
                                stacked
                                placeholder="Name of your project"
                            />
                        </div>
                        <div className="mb-4">
                            <TextArea
                                label="Description"
                                register={register("description")}
                                placeholder="Description of your project"
                                stacked
                            />
                        </div>
                        <div className="mb-4">
                            <TextField
                                label="Channel ID"
                                register={register("channel_id")}
                                placeholder="Channel ID of target Discord Text channel"
                                stacked
                            />
                        </div>
                        {/* <div className="mb-4">
                            <TextArea
                                label="Source Repository"
                                register={register("project_source_url")}
                                placeholder="Source repository URL"
                            />
                        </div> */}
                        <button
                            type="submit"
                            className="rounded-xl p-2 text-white font-bold hover:bg-blue-500 bg-blue-400 ml-auto hover:cursor-pointer"
                        >
                            Create
                        </button>
                    </form>
                </div>
            </Drawer>
        </div>
    );
};

export default ProjectListPage;
