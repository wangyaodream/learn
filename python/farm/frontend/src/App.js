import {userEffect, useState} from 'react';
import axios from 'axios';
import "./App.css";
import ListToDoLists from "./ListToDoList";
import ToDoList from "./ToDoList";

function App() {
    const [listSummaries, setListSummaries] = useState(null);
    const [selectedItem, setSelectedItem] = useState(null);

    useEffect(() => {
        reloadData().catch(console.error);
    }, []);

    async function reloadData() {
        const response = await axios.get("/api/lists");
        const data = await response.data;
        setListSummaries(data);
    }


}
