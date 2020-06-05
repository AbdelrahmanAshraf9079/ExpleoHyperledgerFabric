import React from 'react'
import "../MainPage.css"
import logo from "../expleo.png"

function Appbar() {
  return (
    <div className="Appbar">
      <img src={logo} alt="Logo" className="logo" />
  </div>
  )
}


export default Appbar ;