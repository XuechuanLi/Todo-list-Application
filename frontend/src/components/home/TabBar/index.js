import React, { useState } from 'react'
import  {Link} from 'react-router-dom'
import '../../../assets/css/font.css'
import './index.css'


export default function TabBar(props) {

    const { handleTabBarClick, activeTab } = props;

    return (
        <div className="tabBar">
            {/* click item and change state */}
            <div className={"tabBarItem" + (activeTab === "todoList" ? " active" : "")} onClick={() => handleTabBarClick("todoList")}>
                {/* render icons based on state */}
                <i className={"iconfont icon-bar-chart-2-" + (activeTab === "hot" ? "fill" : "line")}></i>
                <span className="tabBarSubTitle">Todo</span>
            </div>
            <div className={"tabBarItem" + (activeTab === "groups" ? " active" : "")} onClick={() => handleTabBarClick("groups")}>
                <i className={"iconfont icon-book-read-" + (activeTab === "news" ? "fill" : "line")}></i>
                <span className="tabBarSubTitle">Groups</span>
            </div>
            <div className={"tabBarItem" + (activeTab === "myPage" ? " active" : "")} onClick={() => handleTabBarClick("myPage")}>
                <i className={"iconfont icon-bookmark-" + (activeTab === "history" ? "fill" : "line")}></i>
                <span className="tabBarSubTitle">My</span>
            </div>
        </div>
    )
}