import { Link, useNavigate, useParams } from "react-router-dom";
import {
    PageTitle,
    SectionDescription,
    SectionTitle,
} from "../../../components/Typography.tsx";
import { FormProvider, useForm } from "react-hook-form";
import TextField from "../../../components/TextField.tsx";
import TextArea from "../../../components/TextArea.tsx";
import httpClient from "../../../lib/httpClient.ts";
import { useQueryClient } from "@tanstack/react-query";

export default function AddProjectEventTemplate() {
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const form = useForm();
    const { register } = form;
    const { projectId } = useParams();

    const onSubmit = async (data: any) => {
        const payload = { ...data, project_id: projectId };
        const res = await httpClient.post(`projects/${projectId}/templates`, {
            json: payload,
        });
        if (!res.ok) {
            console.error("Failed to create event template");
            return;
        }
        // Handle success
        const responseData = await res.json();
        console.log("Event template created:", responseData);
        queryClient.invalidateQueries({
            queryKey: ["projects", projectId, "event-templates"],
        });
        navigate(`../..`, { relative: "path" });
    };
    if (!projectId) {
        return <div>Project ID is required</div>;
    }
    return (
        <div className="p-4">
            <PageTitle>Add Event Template</PageTitle>
            <nav className="mb-4 text-sm text-gray-500">
                <Link to="../../.." relative="path">
                    Projects
                </Link>{" "}
                /{" "}
                <Link to="../.." relative="path">
                    Project Detail
                </Link>
                {" / "}
                <span>Add Event Template</span>
            </nav>
            <FormProvider {...form}>
                <form
                    className="flex flex-col bg-white shadow-md p-4 rounded"
                    onSubmit={form.handleSubmit(onSubmit)}
                >
                    <SectionTitle>Event Template</SectionTitle>
                    <SectionDescription>
                        Create a template that will be used to render message in
                        respect to webhook payload received
                    </SectionDescription>
                    <div className="my-2" />
                    <TextField
                        label="Event Type"
                        register={register("event_type")}
                        stacked
                    />
                    <TextArea
                        label="Template"
                        register={register("template")}
                        stacked
                    />
                    <div className="flex gap-x-4 mt-4">
                        <button
                            className="bg-blue-400 text-white font-bold p-2 rounded hover:bg-blue-500 hover:cursor-pointer mb-2 w-[100px]"
                            type="submit"
                        >
                            Save
                        </button>
                        <Link
                            to="../.."
                            relative="path"
                            className="bg-red-400 text-white font-bold p-2 rounded hover:bg-red-500 hover:cursor-pointer mb-2 w-[100px] flex justify-center items-center"
                        >
                            Cancel
                        </Link>
                    </div>
                </form>
            </FormProvider>
        </div>
    );
}
