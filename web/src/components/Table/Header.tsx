import { cn } from "../../lib/index.ts";

type Props = {
    columnWidth: number;
    className?: string;
    children?: React.ReactNode;
    isFirst?: boolean;
    isLast?: boolean;
    colSpan?: number;
    [key: string]: unknown;
};

export const HeaderCell: React.FC<Props> = ({
    columnWidth,
    className,
    children,
    isFirst,
    isLast,
    colSpan,
    ...props
}) => {
    return (
        <th
            colSpan={colSpan}
            className={cn(
                "text-xl font-semibold px-2 py-1 text-left bg-indigo-400 text-white border border-gray-200 border-b-0 border-l-0 break-all",
                isFirst ? "rounded-tl-lg border-l" : "",
                isLast ? "rounded-tr-lg" : "",
                className
            )}
            style={{ width: columnWidth }}
            {...props}
        >
            {children}
        </th>
    );
};
