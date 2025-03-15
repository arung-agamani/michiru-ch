import { ensureDir } from "https://deno.land/std@0.95.0/fs/mod.ts";

// Function to convert a string to snake case
function toSnakeCase(str: string): string {
    return str
        .replace(/([a-z])([A-Z])/g, "$1_$2")
        .replace(/[\s-]+/g, "_")
        .toLowerCase();
}

// Function to get the next migration number
async function getNextMigrationNumber(): Promise<number> {
    const migrationDir = "./db/migrations";
    await ensureDir(migrationDir);
    const files = [...Deno.readDirSync(migrationDir)];
    const migrationNumbers = files
        .map((file) => parseInt(file.name.split("_")[0], 10))
        .filter((num) => !isNaN(num));
    return migrationNumbers.length > 0 ? Math.max(...migrationNumbers) + 1 : 1;
}

// Main function to create migration files
async function createMigrationFiles(action: string) {
    const snakeCaseAction = toSnakeCase(action);
    const migrationNumber = (await getNextMigrationNumber())
        .toString()
        .padStart(4, "0");
    const upFileName = `./db/migrations/${migrationNumber}_${snakeCaseAction}.up.sql`;
    const downFileName = `./db/migrations/${migrationNumber}_${snakeCaseAction}.down.sql`;

    await Deno.writeTextFile(
        upFileName,
        "-- Write your 'up' migration SQL here\n"
    );
    await Deno.writeTextFile(
        downFileName,
        "-- Write your 'down' migration SQL here\n"
    );

    console.log(`Created migration files:\n${upFileName}\n${downFileName}`);
}

// Get user input
const action = prompt("Enter the name of the migration action:");
if (action) {
    createMigrationFiles(action);
} else {
    console.error("No action provided.");
}
