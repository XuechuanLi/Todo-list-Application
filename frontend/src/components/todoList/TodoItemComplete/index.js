import React, { useState } from 'react'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import './index.css'

export default function TodoItemComplete(props) {

    const [description, setDescription] = useState("");
    const [photo, setPhoto] = useState();

    const handleDescriptionChange = (e) => {
        setDescription(e.target.value);
    }

    const handlePhotoChange = (e) => {
        setPhoto(e.target.value);
    }

    const {handleTodoItemSubmitBtnClick, handleTodoItemCancelBtnClick} = props

    return (
        <div className="todoItemComplete">
            <Form onSubmit={(e) => {handleTodoItemSubmitBtnClick(e, {description}, {photo})}}>
                <Form.Group className="mb-3 todoItemCompleteText">
                    <Form.Label>Description</Form.Label>
                    <Form.Control placeholder="Enter description" required="required" onChange={handleDescriptionChange}/>
                </Form.Group>
                <Form.Group className="mb-3 todoItemCompletePhoto" type="file">
                    <Form.Label>Photo</Form.Label>
                    <br></br>
                    <input type="file" required="required" onChange={handlePhotoChange}/>
                </Form.Group>
                <Button className="todoListCompleteSubmit" variant="primary" type="submit">
                    Submit
                </Button>
                <Button className="todoListCompleteCancel" variant="primary" onClick={handleTodoItemCancelBtnClick}>
                    Cancel
                </Button>
                {/* <button className="btn1">test</button> */}
            </Form>
        </div>
    )
}