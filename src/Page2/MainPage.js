import React from "react";
import Appbar from "./components/Appbar";
import Row from "./components/Row";
import Sidebar from "./components/Sidebar";


function MainPage() {
return (
    <div style={{height:"100%"}}>
    <Appbar />

    <div className="listView">
    <Sidebar/>
    <Row />
    </div>
    </div>
);
}

export default MainPage;
