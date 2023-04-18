import React, { useState, useEffect } from 'react';
import GroupItem from '../../components/groups/GroupItem';
import SearchBar from '../../components/groups/SearchBar';
import JoinGroup from '../../components/groups/JoinGroup';
import CreateGroup from '../../components/groups/CreateGroup';
import './index.css';

function Groups () {
    const [showCreateGroup, setShowCreateGroup] = useState(false);

    const [showJoinGroup, setShowJoinGroup] = useState(false);
    const [selectedGroup, setSelectedGroup] = useState({ ID: "" });
    const [keyWord, setKeyWord] = useState("");

    //items should be got from back-end daabase
    const allGroups = [
        {
            ID: "11",
            title: "Study for an hour",
            description: "Please upload a photo of study and describe what you learned.",
            finishDeadline: "22:00",
            groupStartTime: "12/01/2022",
            groupEndTime: "12/30/2022",
            jackPot: 100,
            allPeople: 30
        },
        {
            ID: "12",
            title: "Tennis",
            description: "Please upload a photo of tennis court.",
            finishDeadline: "20:00",
            groupStartTime: "12/10/2022",
            groupEndTime: "1/31/2023",
            jackPot: 150,
            allPeople: 39
        },
        {
            ID: "13",
            title: "Get up",
            description: "Please upload a photo such as breakfast to prove you got up.",
            finishDeadline: "8:00",
            groupStartTime: "12/21/2022",
            groupEndTime: "1/10/2023",
            jackPot: 180,
            allPeople: 51
        },
        {
            ID: "14",
            title: "Memorize 30 English words",
            description: "Please upload a photo English words.",
            finishDeadline: "23:59",
            groupStartTime: "12/01/2022",
            groupEndTime: "12/25/2022",
            jackPot: 310,
            allPeople: 100
        },
        {
            ID: "15",
            title: "Memorize 30 English words",
            description: "Please upload a photo English words.",
            finishDeadline: "23:59",
            groupStartTime: "12/01/2022",
            groupEndTime: "12/25/2022",
            jackPot: 310,
            allPeople: 100
        },
        {
            ID: "16",
            title: "Memorize 30 English words",
            description: "Please upload a photo English words.",
            finishDeadline: "23:59",
            groupStartTime: "12/01/2022",
            groupEndTime: "12/25/2022",
            jackPot: 310,
            allPeople: 100
        }
    ];

    const joinedGroups = [
        {
            ID: "11"
        },
        {
            ID: "14"
        }
    ];

    const handleSearchBarClick = (content) => {
        setKeyWord(content);
    };
    //Create Group
    const handleCreateBtnClick = () => {
        setShowCreateGroup(!showCreateGroup);
    };
    const handleCreateGroupCancelBtnClick = () => {
        setShowCreateGroup(false);
    };
    const handleCreateGroupSubmitBtnClick = (e) => {
        e.preventDefault();
        /* setShowJoinGroup(false);
        setSelectedGroup({ ID: selectedGroup.ID });
        //send message here, userID, joined one group, join points
        console.log(selectedGroup);
        console.log(joinPoints); */
    };

    //Join Group
    const handleJoinBtnClick = (itemID) => {
        setShowJoinGroup(true);
        setSelectedGroup({ ID: itemID });
    };

    const handleJoinGroupCancelBtnClick = () => {
        setShowJoinGroup(false);
    };

    const handleJoinGroupSubmitBtnClick = (e, joinPoints) => {
        e.preventDefault();
        setShowJoinGroup(false);
        setSelectedGroup({ ID: selectedGroup.ID });
        //send message here, userID, joined one group, join points
        console.log(selectedGroup);
        console.log(joinPoints);
    };
    useEffect(() => {
        console.log(keyWord);
    });

    return (
        <div className="groupPage">
            {
                showJoinGroup === true ? <JoinGroup className="joinGroup"
                    handleJoinGroupSubmitBtnClick={handleJoinGroupSubmitBtnClick}
                    handleJoinGroupCancelBtnClick={handleJoinGroupCancelBtnClick}></JoinGroup> : null
            }
            {
                showCreateGroup === true ? <CreateGroup className="createGroup"
                    handleCreateGroupSubmitBtnClick={handleCreateGroupSubmitBtnClick}
                    handleCreateGroupCancelBtnClick={handleCreateGroupCancelBtnClick}></CreateGroup> : null
            }
            <div className="groupHead">
                <div className="searchBarBox">
                    <SearchBar className="groupSearchBar" handleSearchBarClick={handleSearchBarClick}></SearchBar>
                </div>
                <div className="createBox">
                    <button className="createGroupBtn" onClick={handleCreateBtnClick}>Create</button>

                </div>

            </div>
            <div className="groupList">
                {allGroups.map(group => ((keyWord.length === 0 || group.ID.indexOf(keyWord) >= 0 || group.title.indexOf(keyWord) >= 0 || group.description.indexOf(keyWord) >= 0) === true ? <GroupItem handleJoinBtnClick={handleJoinBtnClick} item={group}></GroupItem> : null))}
            </div>
        </div>
    );
}

export default Groups;