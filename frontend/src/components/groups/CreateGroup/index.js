import React, { useState } from 'react';
import './index.css';
//import Button from 'react-bootstrap/Button';
import { DatePicker, TimePicker, Space, Form, message, Button, Input } from 'antd';
import dayjs from 'dayjs';
import customParseFormat from "dayjs/plugin/customParseFormat";
dayjs.extend(customParseFormat);
const { RangePicker } = DatePicker;
const disabledDate = (current) => {
  // Can not select days before today and today
  return current && current < dayjs().endOf('day');
};

const CreateGroup = (props) => {
  const [groupItemName, setGroupItemName] = useState("");

  /*   const handleCreateGroupBtnClick = () => {
      setShowCreateGroup(true);
    };
  
    const handleCreateGroupCancelBtnClick = () => {
      setShowCreateGroup(false);
    };
  
    const handleCreateGroupSubmitBtnClick = (e, description, photo) => {
      e.preventDefault();
      setShowCreateGroup(false);
      //send message here
  
       console.log(selectedItem);
      console.log(description);
      console.log(photo); 
<Form.Item className="mb-3 createGroupText">
          
          <Form.Label>What do you want to achieve?</Form.Label>
          <Form.Control placeholder="eg:Get up before 8:00" required="required" /* onChange={handleGroupItemNameChange} */


  const { handleCreateGroupSubmitBtnClick, handleCreateGroupCancelBtnClick } = props;
  const [form] = Form.useForm();
  const config = {
    rules: [
      {
        type: 'object',
        required: true,
        message: 'Please select time!',
      },
    ],
  };
  const onFinish = () => {
    message.success('Submit success!');
  };
  const onFinishFailed = () => {
    message.error('Submit failed!');
  };
  return (
    <div className='createGroup'>
      {/* <Form onSubmit={(e) => { handleCreateGroupSubmitBtnClick(e, { groupItemName }); }}> */}
      <Form form={form}
        layout="vertical"
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off">
        <Form.Item
          name="goal"
          label="Set a goal to achieve!"
        >
          <Input placeholder="eg: Start working before 8:00" />
        </Form.Item>

        <Form.Item
          name="evidence"
          label="Completion evidence:"

        >
          <Input placeholder="eg: a photo of library." />
        </Form.Item>

        <Form.Item name="completeTime" label="Daily completion time:" >
          <TimePicker />
        </Form.Item>

        <Form.Item
          name="startEndTime"
          label="Duration:">
          <Space direction="vertical" size={12}>
            <RangePicker disabledDate={disabledDate} />
          </Space>
        </Form.Item>

        <Form.Item>
          <Space>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
            <Button type="primary" onClick={handleCreateGroupCancelBtnClick}>
              Cancel
            </Button>
          </Space>
        </Form.Item>


      </Form>

    </div>
  );
};

export default CreateGroup;