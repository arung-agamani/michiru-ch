import { Link } from "react-router-dom";
import { authStateAtom } from "../../state/auth.ts";
import { useAtom } from "jotai";

const AppIndexPage = () => {
    const [authData] = useAtom(authStateAtom);
    return (
        <div>
            <h1 className="text-4xl font-semibold">
                Good Day, {authData.user?.username}
            </h1>
            <hr className="my-2 grey" />
            <div className="mt-4 p-4 bg-blue-100 rounded-lg">
                <Link to="projects" className="text-blue-700 hover:underline">
                    Go to Projects
                </Link>
            </div>
            <div className="mt-4 p-4 bg-green-100 rounded-lg">
                <Link
                    to="predefined-templates"
                    className="text-green-700 hover:underline"
                >
                    Predefined Templates
                </Link>
            </div>
        </div>
    );
};

export default AppIndexPage;
