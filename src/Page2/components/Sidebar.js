import React from 'react'

function Sidebar() {
    return (
        <div className="sideBar">

            <div className='userIcon'>
            <i class="fas fa-user-tie" style={{fontSize:'40px'}}></i>
            </div>
            <h3> User Name </h3>

            <br></br>
            <div className='userIcon' >
            <i class="fas fa-sitemap" style={{fontSize:'40px'}} ></i>
            </div>
            <h3> Team  </h3>

            <br></br>
            <div className='userIcon' >
            <i class="fas fa-project-diagram" style={{fontSize:'40px'}}></i>
            </div>
            <h3> Cahnnel  </h3>
        </div>
    )
}

export default Sidebar ;