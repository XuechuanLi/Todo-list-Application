import React, {useState} from 'react';

import {Route, Routes, Link} from 'react-router-dom'

import TodoList from '../TodoList'
import Groups from '../Groups'
import MyPage from '../MyPage'
import TabBar from '../../components/home/TabBar'

function Home (){

    const [page, setPage] = useState('todoList');

    const handleTabBarClick = (active) => {
        setPage(active);
        console.log(active);
    }

    return (
        <div className="Home">
            This is Home...
            {
                page === 'todoList' ? <TodoList /> :
                page === 'groups' ? <Groups /> : <MyPage />
            }
            <TabBar handleTabBarClick={handleTabBarClick} activeTab={page}/>
                <Routes>
                    <Route path='todoList' element={<TodoList />}></Route>
                    <Route path='groups' element={<Groups />}></Route>
                    <Route path='myPage' element={<MyPage />}></Route>
                </Routes>
        </div>
    );
}

export default Home;