import React from 'react';

import './index.css'
import {Link} from 'react-router-dom'


function Register (){
    return (
        <div className="register">
            This is Register...
            <div className="buttons">
                <Link to="/login"><button type="button" className="btn btn-primary registerToLogin">Register</button></Link>
            </div>
        </div>
    )
}

export default Register;