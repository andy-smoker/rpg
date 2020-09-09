import React from 'react'


export class Header extends React.Component{
    render(){
    return(
        <header>
            Header <br />
            <button type="button">
                <a href="/profile"> Профиль </a>
            </button>
            <button type="button">
                <a href="/create"> Создать </a>
            </button>
        </header>
    )}
}