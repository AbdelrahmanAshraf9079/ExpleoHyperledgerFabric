import React, { useState } from "react";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import './LoginPage.css'
import axios from 'axios'
import logo from './logo.png';
import { Button } from "@material-ui/core";

function LoginPage() {

const [userName, setUserName] = useState();
const [teamName, setTeamName] = useState();
const [channelName, setChannelName] = useState();

const channelOptions = [{label :"Main Channel", value : "mainchannel"},{label :"Client Channel",value : "clientchannel"},{label :"Offer Channel",value :"offerchannel"}];
const teamOptions = [
    {label :"Sales", value : "sales"},
    {label :"Resourcing",value: "resourcing"},
    {label :"Engagement Management", value:"engManagement"},
    {label :"Client",value:"client"}
    ];



async function handleClick() {

let userReq = {
        "username": userName,
        "orgName" : teamName
    }

let result = await sendUserReq(userReq);

console.log(result)
}


return (
<div style={{ display: "flex", justifyContent: "center" ,width:"100%",height:'100%'}}>


    <div className= "centerHover" style={{ display: "flex", justifyContent: "center" }}style={{ display: "flex", justifyContent: "center" }}>


        <img src={logo} alt="Logo" className="logo" />



        <input
            style={{ display: "flex" }}
            className="text"
            type="text"
            name="name"
            placeholder="Username"
            onChange={(val) => {
            setUserName(val.target.value);
            }}
        />


        <Dropdown
            className="teams"
            options={teamOptions}
            value={teamName}
            font
            placeholder="Select an option"
            onChange={(val) => {
                setTeamName(val.value);
                }}
        />

        <Dropdown
            className="channels"
            options={channelOptions}
            value={channelName}
            placeholder="Select an option"
            onChange={(val) => {
                setChannelName(val.value);
                }}
        />

        <Button  variant="contained" color="primary"  onClick={handleClick} size='large'>
        Login
        </Button>

    </div>


</div>
);
}

export function sendUserReq(body) {

    return new Promise((resolve) => {
        console.log("body: ",body)
        axios.post("http://localhost:4000/users",body)
            .then(res => {
                resolve(res.data)
            })
            .catch(error => {
            console.log(error)
            })
    })
}


export default LoginPage;
