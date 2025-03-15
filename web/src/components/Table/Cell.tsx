import { cn } from "../../lib/index.ts";

type Props = {
    columnWidth: number;
    className?: string;
    children?: React.ReactNode;
    isFirst?: boolean;
    isLast?: boolean;
    [key: string]: any;
};

export const BodyCell: React.FC<Props> = ({
    columnWidth,
    className,
    children,
    isFirst,
    ...props
}) => {
    return (
        <td
            className={cn(
                "px-2 py-2 text-left border border-gray-200 border-t-0 border-l-0 break-words md:break-all",
                isFirst ? "border-l" : "",
                className
            )}
            style={{ width: columnWidth }}
            {...props}
        >
            {children}
        </td>
    );
};
