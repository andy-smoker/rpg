import React from 'react'
import style from './index.module.css'
import fStyle from '../../index.module.css'

export const Flaws = () => {
    return (
        <div className={style.main}>
            <div className={`${fStyle.cell} ${style.title}`}>
                <p>13. ИНЦИДЕНТЫ ПОТЕРИ РАССУДКА БЕЗ БЕЗУМИЯ</p>
            </div>
            <div className={`${fStyle.cell} ${style.cell}`}>
                <p>НАСИЛИЕ
                    <input type='checkbox' name='voilence' />
                    <input type='checkbox' name='voilence' />
                    <input type='checkbox' name='voilence' />
                </p>
            </div>
            <div className={`${fStyle.cell} ${style.cell}`}>
                <p>БЕСПОМОЩНОСТЬ
                    <input type='checkbox' name='helpness' />
                    <input type='checkbox' name='helpness' />
                    <input type='checkbox' name='helpness' />
                </p>
            </div>
        </div>
    )
}