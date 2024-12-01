import "./ListToDoLists.css";
import { useRef } from "react";
import { BiSolidTrash } from "react-icons/bi";

function ListToDoLists({
    listSummaries,
    handleSelectList,
    handleDeleteToDoList,
    handleNewToDoList,
}) {
    const labelRef = useRef();

    if (listSummaries === null) {
        return <div className="ListToDoLists loading">Loading... to-do lists</div>;
    } else if (listSummaries.length === 0) {
        return (
            div
        );
    }
}
