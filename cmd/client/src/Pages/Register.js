import React ,{useState}from 'react'

export function Register(){
    return (
        <div>
        < RegForm />
        </div>
    )
}

function RegForm() {
    const [user, setUser] = useState({
        login: null,
        username: null,
        password: null
    })

    function sendForm() {
        fetch("http://localhost:8080/reg",{
            method: "POST",
            body: JSON.stringify(user)
        }).then(resp => (console.log(resp.headers.get("Sattus-Code"))))
    }

    return(
        <div>
            RegisterPage  <br />
            <button> <a href="/">Главная</a> </button> <br />
            <label>
                Login
                <input name="login" onChange={(e) => {user.login = (e.target.value)}} />
            </label> <br />
            <label>
                Username
                <input name="username" onChange={(e) => {user.username = (e.target.value)}} />
            </label> <br />
            <label>
                Password
                <input type="password" name="password" onChange={(e) => {user.password = (e.target.value)}} />
            </label> <br />
            <button onClick={sendForm}>Registrate</button>
        </div>
    )
}