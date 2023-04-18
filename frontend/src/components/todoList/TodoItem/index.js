import React, { useState, useEffect } from 'react'
import './index.css'

export default function TodoItem(props) {

    const [completed, setCompleted] = useState(false);

    const {handleTodoItemCompleteBtnClick, item, selectedItem} = props;

    useEffect(() => {
        if (item.completed || (item.ID === selectedItem.ID && selectedItem.completed)) {
            setCompleted(true);
        }
    });

    return (
        <div className="todoItem">
            <div className="todoItemName">{item.title}</div>
            <div className="todoItemDescription">{item.description}</div>
            <div className="todoItemCompleteDeadline">Please finish before {item.finishDeadline}</div>
            <div className="todoItemEndTime">End time: {item.groupEndTime}</div>
            <div className="todoItemRemainingPeople">Remaining people: {item.remainingPeople}</div>
            <div className="todoItemJackPot">Jack pot: {item.jackPot}</div>
            {completed ? <div>Completed</div> : <button type="button" className="todoItemCompleteBtn" onClick={() => handleTodoItemCompleteBtnClick(item.ID)}>Complete</button>}
        </div>
    )
}