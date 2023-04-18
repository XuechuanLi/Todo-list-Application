import React, {useState} from 'react';

import {BrowserRouter as Router, Route, Routes, Link} from 'react-router-dom'

// import pages components
import TodoList from './pages/TodoList'
import Groups from './pages/Groups'
import MyPage from './pages/MyPage'
import Login from './pages/Login'
import Register from './pages/Register'
import Home from './pages/Home'
import Welcome from './pages/Welcome'

import {loginUser, onLogin, logout} from './cookie'

// import antd components
import { Button } from 'antd';

function App() {

    // logout()
    // onLogin()
    const userInfo = loginUser()
    console.log(userInfo)

    const [page, setPage] = useState("login")
    console.log(page)
    // if (userInfo && window.location.pathname === "/") {
    //     // window.location.pathname = '/todoList'
    //     setPage("home")
    // }
    // // window.location.href = page

    // // if (!loginUser()) page = setPage("todoList")
    // console.log(page)

    return (
        <div className="App">

            <Router>
                {/* Router config */}
                <Routes>
                    {page }
                    <Route index path='/' element={<Welcome />}></Route>
                    <Route path='/login' element={<Login />}></Route>
                    <Route path='/register' element={<Register />}></Route>
                    <Route path='/todoList' element={<Home />}></Route>
                    <Route path='/groups' element={<Home />}></Route>
                    <Route path='/myPage' element={<Home />}></Route>
                </Routes>
            </Router>
        </div>
    );
}

export default App;
