import React, { useState, useEffect } from 'react'
import './index.css'

export default function GroupItem(props) {

    // const [completed, setCompleted] = useState(false);

    const { handleJoinBtnClick, item } = props;

    // useEffect(() => {
    //     if (item.completed || (item.ID === selectedItem.ID && selectedItem.completed)) {
    //         setCompleted(true);
    //     }
    // });

    return (
        <div className="groupItem">
            <div className="groupItemName">{item.title}</div>
            <div className="groupItemDescription">{item.description}</div>
            <div className="groupItemCompleteDeadline">Please finish before {item.finishDeadline}</div>
            <div className="groupItemEndTime">End time: {item.groupEndTime}</div>
            <div className="groupItemRemainingPeople">Joining people: {item.allPeople}</div>
            <div className="groupItemJackPot">Jack pot: {item.jackPot}</div>
            <button type="button" className="groupJoinBtn" onClick={handleJoinBtnClick}>Join</button>
        </div>
    )
}