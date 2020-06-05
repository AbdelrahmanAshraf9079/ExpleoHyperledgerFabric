import React from 'react'
import "../MainPage"

function Row() {
    return (
        <div className= "row"  >
        <div style={{color:"white"}}>
        <i class="fas fa-file-alt" style={{fontSize:"25px"}}></i>
        </div>
        <h3 style={{color:"white"}} className="subject" > Document.subject</h3>
        <h3 style={{color:"white"}} className="creationDate" > Document.creationDate</h3>
        </div>
    )
}


export default Row ;