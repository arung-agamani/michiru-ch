import { useFormContext } from "react-hook-form";
import TextField from "../TextField.tsx";
import { Project } from "../../types.ts";
import { useState } from "react";
import httpClient, { APIResponse } from "../../lib/httpClient.ts";
import { useQueryClient } from "@tanstack/react-query";

const ProjectSource = () => {
    const form = useFormContext<Project>();
    const [editMode, setEditMode] = useState(false);
    const queryClient = useQueryClient();

    const onSubmit = async (data: Project) => {
        const payload = { ...data };
        const res = await httpClient.put<APIResponse<Project>>(
            `projects/${data.id}/source`,
            {
                json: payload,
            }
        );
        if (!res.ok) {
            console.error("Failed to update project source");
            return;
        }
        // Handle success
        const responseData = await res.json();
        console.log("Project source updated:", responseData);
        queryClient.invalidateQueries({
            queryKey: ["projects", data.id],
        });
        form.reset(responseData.data);
        setEditMode(false);
    };
    return (
        <form className="lg:max-w-1/2" onSubmit={form.handleSubmit(onSubmit)}>
            <TextField
                label="Source"
                register={form.register("project_source_url")}
                stacked
                editMode={editMode}
                viewAsLink
            />
            <TextField
                label="Auth Token"
                register={form.register("project_source_auth_token")}
                stacked
                editMode={editMode}
            />
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
        </form>
    );
};

export default ProjectSource;
