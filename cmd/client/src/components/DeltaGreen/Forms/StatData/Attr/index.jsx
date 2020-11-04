import React from 'react'
import style from './index.module.css'
import sdStyle from '../index.module.css'
import fStyle from '../../index.module.css'

export const Attr = (props) => {
    return (
        <div className={sdStyle.attr}>
            <div className={style.attr_row}>
                <div className={`${style.cell} ${fStyle.cell}`} > DEVIDED ATTRIBUTES</div>
                <div className={`${style.cell} ${fStyle.cell}`} > MAXIMUM </div>
                <div className={`${style.cell} ${fStyle.cell}`} > CURRENT </div>
            </div>
            < AttrRow name='ЖИЗНЬ (HP)' v={parseInt(props.val.con)}/>
            < AttrRow name='СИЛА ВОЛИ (WP)' v={parseInt(props.val.pow)}/>
            < AttrRow name='РАССУДОК (SAN)' v={parseInt(props.val.pow)*5}/>
            < AttrRow name='ТОЧКА СЛОВМА (BP)' v={(parseInt(props.val.pow)*5)-(parseInt(props.val.pow))}/>
        </div>
    )
}

const AttrRow = (prop) => {
    return (
        <div className={style.attr_row}>
            <div className={` ${fStyle.cell}`}> <p>{prop.name}</p> </div>
            <div className={` ${fStyle.cell}`}> <p name={'max'+prop.tag} > {prop.v} </p> </div>
            <div className={`${fStyle.cell}`}> <input type='text'name={'curr'+prop.tag}/> </div>
        </div>
    )
}
