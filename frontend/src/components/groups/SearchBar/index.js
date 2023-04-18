import React, { useState, useEffect } from 'react'
import './index.css'

export default function SearchBar(props) {

    const [keyWord, setKeyWord] = useState("");
    const { handleSearchBarClick } = props;

    const handleChange = (e) => {
        setKeyWord(e.target.value);
    }

    return (
        <div className="searchBar">
            <input type="text" className="searchInput" placeholder="Search groups" onChange={handleChange}/>
            <button type="button" className="searchBtn" onClick={() => handleSearchBarClick(keyWord)}>Search</button>
        </div>
    )
}