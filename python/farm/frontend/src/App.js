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

    function handleNewToDoList(newName) {
        const updateData = async () => {
            const newListData = {
                name: newName,
            };

            await axios.post("/api/lists", newListData);
            reloadData().catch(console.error)
        }
        updateData();
    }

    function handleDeleteToDoList(id) {
        const updateData = async () => {
            await axios.delete(`/api/lists/${id}`);
            reloadData().catch(console.error);
        };
        updateData();
    }

    function backToList() {
        setSelectedItem(null);
        reloadData().catch(console.error);
    }

    if (selectedItem === null) {
        return (
            <div className='App'>
                <ListToDoLists
                    listSummaries={listSummaries}
                    handleSelectList={handleSelectList}
                    handleNewToDoList={handleNewToDoList}
                    handleDeleteToDoList={handleDeleteToDoList}
                />
            </div>
        );
    } else {
        return (
            <div className='App'>
                <ToDoList
                    selectedItem={selectedItem}
                    handleBackButton={backToList}
                />
            </div>
        );
    }
}

export default App;
