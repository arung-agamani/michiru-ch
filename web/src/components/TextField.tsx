import React, { useEffect } from "react";
import { useFormContext, UseFormRegisterReturn } from "react-hook-form";
import { Input } from "@headlessui/react";
import { cn } from "../lib/index.ts";
interface TextFieldProps extends React.InputHTMLAttributes<HTMLInputElement> {
    label: string;
    editMode?: boolean;
    register: UseFormRegisterReturn;
    stacked?: boolean;
}

export const TextField: React.FC<TextFieldProps> = ({
    label,
    editMode = true,
    register,
    stacked = false,
    ...props
}) => {
    const { formState } = useFormContext();
    return (
        <div
            className={cn(
                "grid grid-cols-2 items-center",
                stacked && "flex flex-col items-start pb-4"
            )}
        >
            <label className="font-bold">{label}</label>
            {!editMode ? (
                <span className="text-xl">
                    {formState.defaultValues?.[register.name] ?? "No data"}
                </span>
            ) : (
                <Input
                    {...register}
                    className={cn(
                        "px-2 py-1 w-full rounded-md",
                        "border border-blue-200",
                        "data-focus:border-blue-500 data-focus:outline-0"
                    )}
                    {...props}
                />
            )}
        </div>
    );
};

export default TextField;
