import React from 'react'
import style from './index.module.css'
import logo from '../../images/logo.webp'
import Auth from '../Auth'

export const Header = () => {
    return(
        <header className={style.header}>
            <div className={style.img} >
                <img className={style.img} src={logo} />
            </div>
            <div className={style.middle}> Middle </div>
            <div className={style.sigin}>
                <Auth />
            </div>
        </header>
    )
}