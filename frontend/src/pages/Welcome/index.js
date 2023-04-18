import Button from 'react-bootstrap/Button';
import React from 'react';
import {Link} from 'react-router-dom'
import './index.css'

function Welcome (){
    return (
        <div className="welcome">
            <h1 className="content">Welcome!</h1>
            <div className="buttons">
                <Link to="/login"><button type="button" className="btn btn-primary welcomeToLogin">Login</button></Link>
                <Link to="/register"><button type="button" className="btn btn-primary welcomeToRegister">Register</button></Link>
            </div>
        </div>
    )
}

export default Welcome;