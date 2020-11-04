import React from 'react'
import { NavLink } from 'react-router-dom'
import classes from './index.module.css'

export const SideBar = () => {
    return (
        <div className={classes.sidebar} >
            <NavLink className={classes.item} to='/general' >
                <button> GENERAL </button>
            </NavLink>
            <NavLink className={classes.item} to='/dg/about' >
                <button> DELTA </button>
            </NavLink>
            <NavLink className={classes.item} to='/dg/about' >
                <button> DND </button>
            </NavLink>
            <NavLink className={classes.item} to='/dg/about' >
                <button> SAWAGE </button>
            </NavLink>
            <NavLink className={classes.item} to='/dg/about' >
                <button> TAIL </button>
            </NavLink>


        </div>
    )
}
