import React from 'react'

export function General() {
    return(
        <div>
            <label>
                Login
                <input name="login" />
            </label> <br />
            <label>
                Password
                <input name="password"/>
            </label> <br />
            <a href="/reg"> регистрация</a> <br />
            <button>Login</button>
        </div>
    )
}