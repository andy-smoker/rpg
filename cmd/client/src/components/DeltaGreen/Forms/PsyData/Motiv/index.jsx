import React from 'react'
import style from './index.module.css'
import fStyle from '../../index.module.css'

export const Motiv = () => {
    return (
        <div>
            <div className={style.main}>
                <div className={fStyle.cell}>
                    <p>МОТИВАЦИЯ И ПСИХИЧЕСКИЕ РАССТРОЙТВА</p>
                </div>
                <div className={`${fStyle.cell} ${style.cell}`}>
                    <textarea />
                </div>
            </div>
        </div>
    )
}