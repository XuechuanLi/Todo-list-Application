import React, { useState } from 'react'
import './index.css'
import {Link} from 'react-router-dom'

export default function LoginForm() {
    return (
        <div className="loginForm">
            <ul>
                <li>username</li>
                <li>password</li>
            </ul>
            <Link to="/register">register</Link>
        </div>
    )
}
