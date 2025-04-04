import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { FormProvider, useForm } from "react-hook-form";
import httpClient, { APIResponse } from "../../../lib/httpClient.ts";
import TextField from "../../../components/TextField.tsx";
import TextArea from "../../../components/TextArea.tsx";

const AddTemplatePage = () => {
    const form = useForm();
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = form;
    const [apiErrors, setApiErrors] = useState<string[]>([]);
    const navigate = useNavigate();

    const onSubmit = async (data: any) => {
        try {
            const response = await httpClient.post<APIResponse<any>>(
                "predefined-templates",
                {
                    json: data,
                }
            );

            const result = await response.json();

            if (!response.ok) {
                setApiErrors(result.error || ["Failed to add template"]);
                return;
            }

            navigate("/app/predefined-templates");
        } catch (error) {
            setApiErrors(["An unexpected error occurred"]);
        }
    };

    return (
        <div className="p-4 bg-white shadow-md rounded">
            <h1 className="text-2xl font-semibold mb-4">
                Add Predefined Template
            </h1>
            <FormProvider {...form}>
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-4">
                        <TextArea
                            label="Template Description"
                            register={register("description", {
                                required: "Description is required",
                            })}
                            // error={errors.description?.message}
                        />
                    </div>
                    <div className="mb-4">
                        <TextField
                            label="Event Type"
                            register={register("event_type", {
                                required: "Event Type is required",
                            })}
                            // error={errors.event_type?.message}
                        />
                    </div>
                    <div className="mb-4">
                        <TextArea
                            label="Template Content"
                            register={register("template", {
                                required: "Template content is required",
                            })}
                            // error={errors.description?.message}
                        />
                    </div>
                    {apiErrors.length > 0 && (
                        <div className="mb-4">
                            {apiErrors.map((error, index) => (
                                <p key={index} className="text-red-500 text-sm">
                                    {error}
                                </p>
                            ))}
                        </div>
                    )}
                    <button
                        type="submit"
                        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                    >
                        Save
                    </button>
                </form>
            </FormProvider>
        </div>
    );
};

export default AddTemplatePage;
