import React, { useState } from 'react'
import style from './index.module.css'
import sdStyle from '../index.module.css'
import fStyle from '../../index.module.css'

const stats = [
    { name: "Сила (STR)", tag: "strg" },
    { name: "Тело (CON)", tag: "con" },
    { name: "Ловкость (DEX)", tag: "dex" },
    { name: "Интеллект (INT)", tag: "int"},
    { name: "Воля (POW)", tag: "pow"},
    { name: "Харизма (CHA)", tag: "cha" }
]

export const Stat = (prop) => {
    
    const val = {strg:'0', con:'0', dex:'0',int:'0', pow:'0',cha:'0'}
    

    return {
        val,
        code: (<div className={sdStyle.stat}>
            <div className={style.stat_row}>
                <div className={`${style.cell} ${fStyle.cell}`}> STATISTIC </div>
                <div className={`${style.cell} ${fStyle.cell}`}> SCORE </div>
                <div className={`${style.cell} ${fStyle.cell}`}> x5=% </div>
                <div className={`${style.dis} ${fStyle.cell}`}> DISTINGTING DEATURE </div>
            </div>
            {stats.map(e => {
                let v, c = StatRow(e)
                val[e.tag] = c.val
                console.log(val)
                return (
                    c.code
                )
            })}

        </div>)
    }
}

const StatRow = (prop) => {
    const [percent, setPercent] = useState(0)
    return {
        val: percent,
        code: (<div className={style.stat_row}>
            <div className={`${style.cell} ${fStyle.cell}`}> <p>{prop.name}</p> </div>
            <div className={`${style.score} ${fStyle.cell}`}> <input type="number" min='0' name={prop.tag} onChange={e => { setPercent(e.target.value) }} /></div>
            <div className={`${style.cell} ${fStyle.cell}`}>  {percent * 5}%  </div>
            <div className={`${style.dis} ${fStyle.cell}`}>  <input type='text' name={'dis_' + prop.tag} /> </div>
        </div>)
    }

}