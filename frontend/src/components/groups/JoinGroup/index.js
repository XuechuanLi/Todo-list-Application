import React, { useState } from 'react'
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import './index.css'

export default function JoinGroup(props) {

    const [joinPoints, setJoinPoints] = useState("");

    const handleJoinPointsChange = (e) => {
        setJoinPoints(e.target.value);
    }

    

    const {handleJoinGroupSubmitBtnClick, handleJoinGroupCancelBtnClick} = props

    return (
        <div className="joinGroup">
            <Form onSubmit={(e) => {handleJoinGroupSubmitBtnClick(e, {joinPoints})}}>
                <Form.Group className="mb-3 joinGroupText">
                    <Form.Label>How many points would you cost to join this group?</Form.Label>
                    <Form.Control placeholder="10" required="required" onChange={handleJoinPointsChange}/>
                </Form.Group>
                <div>Your points: 100</div>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
                <Button variant="primary" onClick={handleJoinGroupCancelBtnClick}>
                    Cancel
                </Button>
            </Form>
        </div>
    )
}