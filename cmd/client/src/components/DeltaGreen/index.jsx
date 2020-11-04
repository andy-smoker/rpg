import React from 'react'
import {  NavLink, Route } from 'react-router-dom'
import { Form } from './Forms'
import About from './About'
import style from './index.module.css'

export const DGgeneral = () => {
    const prefix = "dg"
    return (
        <div className={style.mane}>
            <div className={style.nav}>
                <NavLink className={style.nav_btn} to={`/${prefix}`}>
                    <button> About </button>
                </NavLink>
                <NavLink to={`/${prefix}/form`}>
                    <button> AddForm </button>
                </NavLink>
                <NavLink to={`/${prefix}/list`}>
                    <button> LIST </button>
                </NavLink>
            </div>
            <Route path={`/${prefix}/form`} component={Form} />
            <Route path={`/${prefix}/about`} component={About} />
        </div>
    )
}