import React from 'react'
import style from './index.module.css'
import fStyle from '../../index.module.css'

export const Bonds = () => {
    return (
        <div className={fStyle.bonds}>
            <div className={style.main}>
                <div className={`${fStyle.cell}`} >
                    <p> СВЯЗИ </p>
                </div>
                <div className={`${fStyle.cell}`} >
                    <p> ОЧКИ </p>
                </div>
            </div>
            < BondsRow />
            < BondsRow />
            < BondsRow />

        </div>
    )
}

const BondsRow = (prop) => {
    return(
        <div className={style.main}>
            <div className={`${fStyle.cell}`} >
                <input type='text' name={'bond' + prop.tag} />
            </div>
            <div className={`${fStyle.cell}`} >
                <p > {prop.score} </p>
            </div>
        </div>
    )
}