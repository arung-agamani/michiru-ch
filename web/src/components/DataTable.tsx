import React from "react";
import {
    useReactTable,
    ColumnDef,
    getCoreRowModel,
    flexRender,
} from "@tanstack/react-table";
import { cn } from "../lib/index.ts";
import { HeaderCell } from "./Table/Header.tsx";
import { BodyCell } from "./Table/Cell.tsx";

export type RowData = Record<string, any>;

export interface DataTableProps<TData extends RowData> {
    columns: ColumnDef<TData>[];
    data: TData[];
    contained?: boolean;
}

function DataTable<TData extends RowData>({
    columns,
    data,
    contained = false,
}: DataTableProps<TData>) {
    const table = useReactTable({
        columns,
        data,
        getCoreRowModel: getCoreRowModel(),
        defaultColumn: {
            size: 200,
        },
    });
    return (
        <table className={cn(contained ? "w-full" : "")}>
            <thead className="">
                {table.getHeaderGroups().map((headerGroup) => (
                    <tr
                        key={headerGroup.id}
                        className={cn(contained ? "" : "flex")}
                    >
                        {headerGroup.headers.map((header) => {
                            return (
                                <HeaderCell
                                    key={header.id}
                                    colSpan={header.colSpan}
                                    columnWidth={header.column.getSize()}
                                    className={cn(
                                        contained
                                            ? ""
                                            : "flex justify-between items-center"
                                    )}
                                    isFirst={header.index === 0}
                                    isLast={
                                        header.index ===
                                        headerGroup.headers.length - 1
                                    }
                                >
                                    <div className="flex justify-between w-full">
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(
                                                  header.column.columnDef
                                                      .header,
                                                  header.getContext()
                                              )}
                                        <span>P</span>
                                    </div>
                                </HeaderCell>
                            );
                        })}
                    </tr>
                ))}
            </thead>
            <tbody>
                {table.getRowModel().rows.length === 0 ? (
                    <tr>
                        <td
                            colSpan={columns.length}
                            className="text-center p-4 text-gray-500"
                        >
                            No data
                        </td>
                    </tr>
                ) : (
                    table.getRowModel().rows.map((row) => (
                        <tr
                            key={row.id}
                            className={cn(
                                contained ? "" : "flex",
                                "hover:bg-gray-200"
                            )}
                        >
                            {row.getVisibleCells().map((cell, idx) => (
                                <BodyCell
                                    key={cell.id}
                                    className={cn(
                                        "p-2",
                                        contained
                                            ? ""
                                            : "flex items-center justify-between"
                                    )}
                                    columnWidth={cell.column.getSize()}
                                    isFirst={idx === 0}
                                >
                                    {flexRender(
                                        cell.column.columnDef.cell,
                                        cell.getContext()
                                    )}
                                </BodyCell>
                            ))}
                        </tr>
                    ))
                )}
            </tbody>
        </table>
    );
}

export default DataTable;
