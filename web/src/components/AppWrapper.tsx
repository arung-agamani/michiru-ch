import "@radix-ui/themes/styles.css";
import { Theme } from "@radix-ui/themes";

interface Props {
    children: React.ReactNode;
}
export default function AppWrapper({ children }: Props) {
    return (
        <Theme>
            <div className="main-wrapper bg-gray-50 w-full flex">
                {children}
            </div>
        </Theme>
    );
}
