interface CompulsoryChildren {
    children: React.ReactNode;
}

export function SectionTitle({ children }: CompulsoryChildren) {
    return <h2 className="text-2xl">{children}</h2>;
}

export function SectionDescription({ children }: CompulsoryChildren) {
    return <p className="text-gray-500">{children}</p>;
}

export function PageTitle({ children }: CompulsoryChildren) {
    return <h1 className="text-3xl font-bold">{children}</h1>;
}
