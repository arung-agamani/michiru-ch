import { atom } from "jotai";

interface User {
    id: string;
    email: string;
    username: string;
    created_at: string;
}

interface AuthState {
    user: User | null;
    expiry: number | null;
}

export const authStateAtom = atom<AuthState>({
    user: null,
    expiry: null,
});
