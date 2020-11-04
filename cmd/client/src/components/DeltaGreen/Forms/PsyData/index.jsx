import React from 'react'
import style from './index.module.css'
import fStyle from '../index.module.css'
import { Bonds } from './Bonds'
import { Flaws } from './Flaws'
import { Motiv } from './Motiv'

export const PsyData = () => {
    return (
        <div className={`${fStyle.psy} ${style.main}`}>
            <Bonds />
            <Motiv />
            <Flaws />
        </div>
    )
}