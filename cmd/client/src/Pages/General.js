import React, { useState } from 'react'
import { Redirect } from 'react-router-dom'

export function General() {
    
    return(
        <MyComponent />
    )
}


function MyComponent(params){
    const [user] = useState({
        login: null,
        username: null,
        password: null
    })
    const [data] = useState({
        user: null,
        token: null
    })
    function setData(params) {
        data.user = params.user
        data.token = params.token
    }
    

    async function clk(){
        await fetch("http://localhost:8080/auth",{
            method: 'POST',
            body: JSON.stringify(user)
        }).then(resp => resp.json()).then(json => setData(json) )
        
        localStorage.setItem("token", data.token)

    } 
    
    return(
        <div>
            <label>
                Login
                <input name="login" onChange={(e) => {user.login = (e.target.value)}}/>
            </label> <br />
            
            <label>
                Password
                <input type="password" name="password" onChange={(e) => {user.password = (e.target.value)}}/>
            </label> <br />
            <a href="/reg"> регистрация</a> <br />
            <button onClick={clk}>
                
                <a href="/profile" > Login </a>
            </button>
        </div>
    )
}