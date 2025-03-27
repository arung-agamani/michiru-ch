import React from "react";
import { UseFormRegisterReturn } from "react-hook-form";
import { Textarea } from "@headlessui/react";
import { cn } from "../lib/index.ts";
interface TextAreaProps
    extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
    label: string;
    editMode?: boolean;
    register: UseFormRegisterReturn;
    stacked?: boolean;
}

export const TextArea: React.FC<TextAreaProps> = ({
    label,
    defaultValue,
    editMode = true,
    register,
    stacked = false,
    ...props
}) => {
    return (
        <div
            className={cn(
                "grid grid-cols-2",
                stacked && "flex flex-col items-start gap-y-2"
            )}
        >
            <label className="font-bold">{label}</label>
            {!editMode ? (
                <span>{defaultValue}</span>
            ) : (
                <Textarea
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

export default TextArea;
