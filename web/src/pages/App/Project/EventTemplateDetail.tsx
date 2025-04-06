import { Link, useParams } from "react-router-dom";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import { EventTemplate } from "../../../types.ts";
import Loader from "../../../components/Loader.tsx";
import TextArea from "../../../components/TextArea.tsx";
import { FormProvider, useForm } from "react-hook-form";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import {
    PageTitle,
    SectionDescription,
    SectionTitle,
} from "../../../components/Typography.tsx";

import "github-markdown-css/github-markdown.css";

const payloadRef: Record<string, string> = {
    push: "https://pkg.go.dev/github.com/go-playground/webhooks/v6@v6.4.0/github#PushPayload",
};

const EventTemplateDetailPage = () => {
    const { projectId, eventType } = useParams();
    const queryClient = useQueryClient();
    const form = useForm<EventTemplate>();
    const { data, isSuccess } = useQuery({
        queryKey: ["projects", projectId, "event-templates", eventType],
        queryFn: async () => {
            const res = await httpClient.get<APIResponse<EventTemplate>>(
                `projects/${projectId}/templates/event-type/${eventType}`
            );
            if (!res.ok) {
                throw new Error("Failed to fetch event template");
            }
            const data = await res.json();
            form.reset(data.data);
            return data.data;
        },
        refetchOnWindowFocus: false,
        refetchOnReconnect: false,
        refetchOnMount: false,
    });

    const onSubmit = async (data: EventTemplate) => {
        const payload = { ...data, project_id: projectId };
        const res = await httpClient.post<APIResponse<EventTemplate>>(
            `projects/${projectId}/templates`,
            {
                json: payload,
            }
        );
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
        form.reset(responseData.data);
    };

    if (!projectId || !eventType) {
        return <div>Project ID and Event Type are required</div>;
    }
    if (!isSuccess) {
        return <Loader />;
    }
    return (
        <FormProvider {...form}>
            <form className="p-4" onSubmit={form.handleSubmit(onSubmit)}>
                <PageTitle>Event Template Detail</PageTitle>
                <nav className="mb-4 text-sm text-gray-500">
                    <Link to="../../.." relative="path">
                        Projects
                    </Link>{" "}
                    /{" "}
                    <Link to="../.." relative="path">
                        Project Detail
                    </Link>
                    {" / "}
                    <span>Event Template</span>
                </nav>
                <div className="grid grid-cols-2">
                    <div className="bg-white shadow-md p-4 rounded">
                        <SectionTitle>{data.event_type}</SectionTitle>
                        <SectionDescription>
                            For detail on payload shape in respect to Go string
                            template, please visit{" "}
                            <a
                                href={payloadRef[data.event_type]}
                                target="_blank"
                                className="text-blue-500 hover:underline"
                            >
                                this page
                            </a>
                        </SectionDescription>
                        <div className="my-2"></div>
                        <TextArea
                            label="Description"
                            stacked
                            register={form.register("description")}
                        />
                        <TextArea
                            label="Template Data"
                            register={form.register("template")}
                            stacked
                        />
                        <div className="flex gap-x-4 mt-4">
                            <button
                                className="bg-blue-400 text-white font-bold p-2 rounded hover:bg-blue-500 hover:cursor-pointer mb-2 w-[100px]"
                                type="submit"
                            >
                                Save
                            </button>
                        </div>
                    </div>
                    <div className="bg-white shadow-md p-4 rounded">
                        <SectionTitle>Template Preview.</SectionTitle>
                        <SectionDescription>
                            Note: this preview uses Github-Flavored Markdown
                            which is not an accurate representation of how the
                            message will render inside Discord. Rendering based
                            on latest received event payload will be implemented
                            later. Additionally, Discord styling will be
                            implemented later also.
                        </SectionDescription>
                        <hr className="my-2" />
                        <div className="markdown-body p-4">
                            <Markdown remarkPlugins={[remarkGfm]}>
                                {form.watch("template")}
                            </Markdown>
                        </div>
                    </div>
                </div>
                <div className="mt-4">
                    <h3 className="text-lg font-semibold">Created At</h3>
                    <p>{new Date(data.created_at).toLocaleString()}</p>
                </div>
                <div className="mt-4">
                    <h3 className="text-lg font-semibold">Updated At</h3>
                    <p>{new Date(data.updated_at).toLocaleString()}</p>
                </div>
            </form>
        </FormProvider>
    );
};

export default EventTemplateDetailPage;
