import {
    useState,
    useEffect,
    useId,
    type FC,
    type ReactNode,
    PropsWithChildren,
    useCallback,
    Fragment,
} from "react";
import { createPortal } from "react-dom";
import { Transition, TransitionChild } from "@headlessui/react";
import { cn } from "../lib/index.ts";

export interface PortalProps {
    children: ReactNode;
}

export const Portal: FC<PortalProps> = ({ children }) => {
    const id = useId();
    const [containerAttached, setContainerAttached] = useState(false);

    useEffect(() => {
        if (!containerAttached) {
            const el = document.createElement("div");
            el.id = id;
            document.body.appendChild(el);
            setContainerAttached(true);
        }
        return () => {
            containerAttached && document.getElementById(id)?.remove();
        };
    }, [id, containerAttached]);

    return (
        containerAttached &&
        createPortal(children, document.getElementById(id)!, id)
    );
};

export interface DrawerProps {
    isOpen?: boolean;
    className?: string;
    onDismiss?: () => void;
}

export const Drawer: FC<PropsWithChildren<DrawerProps>> = ({
    isOpen,
    className,
    onDismiss,
    children,
}) => {
    const handleBackdropClick = useCallback(() => {
        onDismiss?.();
    }, [onDismiss]);

    return (
        <Portal>
            <Transition show={isOpen} appear>
                <TransitionChild>
                    <div
                        onClick={handleBackdropClick}
                        className="fixed inset-0 bg-gray-800 opacity-50 data-[closed]:opacity-0 transition-opacity duration-300 ease-in-out"
                    />
                </TransitionChild>
                <TransitionChild>
                    <div
                        className={cn(
                            className,
                            "fixed bottom-0 right-0 h-dvh",
                            "transition duration-300 ease-in-out data-[closed]:translate-x-full"
                        )}
                    >
                        {children}
                    </div>
                </TransitionChild>
            </Transition>
        </Portal>
    );
};
