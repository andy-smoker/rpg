import React from 'react'

export function Register(){
    return(
        <div>
            RegisterPage  <br />
            <button> <a href="/">Главная</a> </button> <br />
            <label>
                Login
                <input name="login" />
            </label> <br />
            <label>
                Password
                <input name="password"/>
            </label> <br />
            <button>Registrate</button>
        </div>
    )
}