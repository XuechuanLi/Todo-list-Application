import React, { useState } from 'react';
import TotoItem from '../../components/todoList/TodoItem';
import TodoItemComplete from '../../components/todoList/TodoItemComplete';
import './index.css';

function TodoList () {

    const [showTodoItemComplete, setShowTodoItemComplete] = useState(false);
    const [selectedItem, setSelectedItem] = useState({ ID: -1, completed: false });

    //items should be got from back-end daabase
    const items = [
        {
            ID: 11,
            completed: false,
            title: "Study for an hour",
            description: "Please upload a photo of study and describe what you learned.",
            finishDeadline: "22:00",
            groupEndTime: "12/30/2022",
            jackPot: 100,
            remainingPeople: 20
        },
        {
            ID: 12,
            completed: false,
            title: "Tennis",
            description: "Please upload a photo of tennis court.",
            finishDeadline: "20:00",
            groupEndTime: "1/31/2023",
            jackPot: 150,
            remainingPeople: 10
        },
        {
            ID: 13,
            completed: true,
            title: "Get up",
            description: "Please upload a photo such as breakfast to prove you got up.",
            finishDeadline: "8:00",
            groupEndTime: "1/10/2023",
            jackPot: 180,
            remainingPeople: 5
        },
        {
            ID: 14,
            completed: false,
            title: "Memorize 30 English words",
            description: "Please upload a photo English words.",
            finishDeadline: "23:59",
            groupEndTime: "12/25/2022",
            jackPot: 310,
            remainingPeople: 22
        }
    ];

    const handleTodoItemCompleteBtnClick = (itemID) => {
        setShowTodoItemComplete(true);
        setSelectedItem({ ID: itemID, completed: false });
    };

    const handleTodoItemCancelBtnClick = () => {
        setShowTodoItemComplete(false);
    };

    const handleTodoItemSubmitBtnClick = (e, description, photo) => {
        e.preventDefault();
        setShowTodoItemComplete(false);
        setSelectedItem({ ID: selectedItem.ID, completed: true });
        //send message here
        console.log(selectedItem);
        console.log(description);
        console.log(photo);
    };

    return (
        <div className="todoList">
            This is TodoList...
            {
                showTodoItemComplete === true ? <TodoItemComplete className="todoItemComplete" handleTodoItemSubmitBtnClick={handleTodoItemSubmitBtnClick} handleTodoItemCancelBtnClick={handleTodoItemCancelBtnClick}></TodoItemComplete> : null
            }
            {
                items.map(item => (<TotoItem handleTodoItemCompleteBtnClick={handleTodoItemCompleteBtnClick} item={item} selectedItem={selectedItem}></TotoItem>))
            }
        </div>
    );
}

export default TodoList;