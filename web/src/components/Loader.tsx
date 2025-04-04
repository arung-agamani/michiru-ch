export function Loader() {
    return <p className="text-center text-gray-500">Loading...</p>;
}

export function FullscreenLoader() {
    return (
        <div className="flex w-full h-full p-4">
            <p className="text-xl">Loading...</p>
        </div>
    );
}

export default Loader;
