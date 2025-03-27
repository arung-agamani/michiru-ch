import { Link, useParams } from "react-router-dom";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { FormProvider, useForm } from "react-hook-form";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import { Project } from "./types.ts";
import TextField from "../../../components/TextField.tsx";
import TextArea from "../../../components/TextArea.tsx";
import { useState } from "react";

type ProjectResponse = APIResponse<Project>;

const ProjectDetailPage = () => {
    const { projectId } = useParams();
    const queryClient = useQueryClient();
    const [editMode, setEditMode] = useState(false);
    const [populated, setPopulated] = useState(false);
    const projectDetailForm = useForm();
    const sendMessageForm = useForm<{ template: string }>();
    const { data, isSuccess } = useQuery({
        queryKey: ["projects", projectId],
        queryFn: async () => {
            const res = await httpClient.get<ProjectResponse>(
                `projects/${projectId}`
            );
            if (!res.ok) {
                throw new Error("Failed to fetch project");
            }
            const data = (await res.json()).data;
            if (!populated) {
                projectDetailForm.reset(data);
                setPopulated(true);
            }
            return data;
        },
    });

    const { register } = projectDetailForm;
    const generateWebhook = async () => {
        const res = await httpClient.post<APIResponse<any>>(
            `projects/${projectId}/webhook`
        );
        if (!res.ok) {
            console.error("Failed to generate webhook URL");
            return;
        }
        const webhookData = (await res.json()).data;
        queryClient.invalidateQueries({ queryKey: ["projects", projectId] });
        console.log(webhookData);
    };
    const sendTestMessage = async (data: { template: string }) => {
        const res = await httpClient.post<APIResponse<any>>(
            `projects/${projectId}/send-message`,
            {
                json: data,
            }
        );
        if (!res.ok) {
            // TODO: handle error by showing a toast
            console.error("Failed to send test message");
            return;
        }
        const messageData = (await res.json()).data;
        console.log(messageData);
    };
    if (!isSuccess || !populated) {
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
            <Link to=".." relative="path">
                <span className="text-blue-500 hover:underline text-xl">
                    Go back to Project list
                </span>
            </Link>
            <FormProvider {...projectDetailForm}>
                {/* TODO: make form for updating project */}
                <div className="bg-white shadow-md rounded p-4 grid grid-cols-2 gap-x-2 mt-4 ">
                    <div className="flex flex-col">
                        <TextField
                            label="Project Name"
                            editMode={editMode}
                            register={register("project_name")}
                            stacked
                        />
                        <TextField
                            label="Description"
                            editMode={editMode}
                            register={register("description")}
                            stacked
                        />

                        <TextField
                            label="Channel ID"
                            editMode={editMode}
                            register={register("channel_id")}
                            stacked
                        />
                        <div className="flex-grow-1" />
                        {editMode ? (
                            <div className="flex gap-x-4">
                                <button
                                    className="bg-blue-400 text-white font-bold p-2 rounded hover:bg-blue-500 hover:cursor-pointer mb-2 w-[100px]"
                                    type="submit"
                                >
                                    Save
                                </button>
                                <button
                                    className="bg-red-400 text-white font-bold p-2 rounded hover:bg-red-500 hover:cursor-pointer mb-2 w-[100px]"
                                    type="button"
                                    onClick={() => {
                                        setEditMode(false);
                                        // projectDetailForm.reset(data);
                                    }}
                                >
                                    Cancel
                                </button>
                            </div>
                        ) : (
                            <button
                                className="bg-amber-200 text-amber-800 font-bold p-2 rounded hover:bg-amber-300 hover:cursor-pointer mb-2 w-[100px]"
                                onClick={() => setEditMode(!editMode)}
                                type="button"
                            >
                                Edit
                            </button>
                        )}
                    </div>
                    <div className="flex flex-col">
                        <div className="mb-4">
                            <h2 className="font-semibold">Added By</h2>
                            <p>{data.added_by}</p>
                        </div>
                        <div className="mb-4">
                            <h2 className="font-semibold">Created At</h2>
                            <p>{new Date(data.created_at).toLocaleString()}</p>
                        </div>
                        <div className="mb-4">
                            <h2 className="font-semibold">Updated At</h2>
                            <p>{new Date(data.updated_at).toLocaleString()}</p>
                        </div>
                        {data.webhook_origin && (
                            <div className="mb-4">
                                <h2 className="font-semibold">
                                    Webhook Origin
                                </h2>
                                <p>{data.webhook_origin}</p>
                            </div>
                        )}
                        {data.webhook_url && (
                            <div className="mb-4">
                                <h2 className="font-semibold">Webhook URL</h2>
                                <p>{data.webhook_url}</p>
                            </div>
                        )}
                        {data.webhook_secret && (
                            <div className="mb-4">
                                <h2 className="font-semibold">
                                    Webhook Secret
                                </h2>
                                <p>{data.webhook_secret}</p>
                            </div>
                        )}
                        <button
                            onClick={generateWebhook}
                            className="bg-blue-400 text-white font-bold p-2 rounded hover:bg-blue-500 hover:cursor-pointer"
                            type="button"
                        >
                            {data.webhook_url && "Re-"}Generate Webhook URL
                        </button>
                    </div>
                </div>
            </FormProvider>
            <div className="bg-white shadow-md rounded p-4 mt-4">
                <p className="text-2xl font-semibold">Send Message</p>
                <form
                    onSubmit={sendMessageForm.handleSubmit(sendTestMessage)}
                    className="grid grid-cols-1 gap-y-4"
                >
                    <TextArea
                        label="Template"
                        register={sendMessageForm.register("template")}
                        stacked
                    />
                    <button
                        type="submit"
                        className="bg-blue-400 text-white font-bold p-2 rounded hover:bg-blue-500 hover:cursor-pointer"
                    >
                        Send
                    </button>
                </form>
            </div>
        </div>
    );
};

export default ProjectDetailPage;
