import React from 'react'
import style from './index.module.css'
import fStyle from '../index.module.css'
import { Stat } from './Stat'
import { Attr } from './Attr'

export const StatData = () => {
    let stat = Stat() 
    
    return (
        <div className={`${style.statdata} ${fStyle.stat}`}>
            {stat.code}
            <Attr val={stat.val}/>
            <div className={`${style.look} ${fStyle.cell}`}> 
            <p>Внешность</p>
            <textarea name={fStyle.title} />

            </div>
        </div>
    )
}




