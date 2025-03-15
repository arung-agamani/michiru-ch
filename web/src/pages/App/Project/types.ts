export interface Project {
    id: string;
    project_name: string;
    description: string;
    channel_id: string;
    added_by: string;
    created_at: string;
    updated_at: string;
    webhook_origin?: string;
    webhook_url?: string;
    webhook_secret?: string;
}
