import { Link } from "react-router-dom";

const IndexPage = () => {
    return (
        <div className="container mx-auto py-8 justify-center flex flex-col items-center">
            <h1 className="text-4xl font-bold">Michiru Ch.</h1>
            <p className="text-2xl italic font-light">
                Imagine Discord for CI/CD
            </p>
            <Link to="/app">
                <button
                    type="button"
                    className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 hover:cursor-pointer"
                >
                    Go to App
                </button>
            </Link>
            <section className="mt-8 container mx-auto">
                <h2 className="text-2xl">Roadmap</h2>
                <ul className="list-disc list-inside">
                    <li>Authentication</li>
                    <ul className="list-decimal list-inside ml-4">
                        <li>Websockets</li>
                        <li>Chat</li>
                    </ul>
                    <li>CI/CD</li>
                    <li>Notifications</li>
                </ul>
            </section>
        </div>
    );
};

export default IndexPage;
