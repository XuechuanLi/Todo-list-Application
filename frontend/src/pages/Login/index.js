import React, { useState } from 'react';
// import './index.css'
import axios from "axios";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

import { useNavigate } from 'react-router-dom'

import {loginUser, onLogin, logout} from '../../cookie'

import './index.css'

function Login() {
    
    const [useremail, setUseremail] = useState("");
    const [password, setPassword] = useState("");
    const [formSuccess, setFormSuccess] = useState(-1);
    const navigate = useNavigate();

    function handleEmailChange(e) {
        setUseremail(e.target.value);
    }

    function handlePasswordChange(e) {
        setPassword(e.target.value);
    }
  
    function handleSubmit(e) {
        e.preventDefault();
        // read email and password
        console.log("useremail: " + useremail);
        console.log("Password: " + password);
        // axios.post("/login", {
        //     useremail: useremail,
        //     password: password
        // }).then(function (response) {
        //     onLogin(response.id);
        //     console.log(response);
        // });
        if (useremail.length < 5) {
            setFormSuccess(0);
        } else {
            setFormSuccess(1);
            setTimeout(() => {
                navigate('/todoList');
            }, 1000);
            
        }
        
    }

    return (
        <>
            <Form className="login" onSubmit={handleSubmit} >
                <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control type="email" placeholder="Enter email" onChange={handleEmailChange} />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="Password" onChange={handlePasswordChange} />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Login
                </Button>
            </Form>
            <div className="loginResult">
                {
                    formSuccess === 0 ? <p>Fail!</p> : 
                    formSuccess === 1 ? <p>Succeed!</p> : null
                }
            </div>
        </>
    );
}

export default Login;