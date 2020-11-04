import React from 'react'
import style from './index.module.css'
import fStyle from '../index.module.css'

export const PerData = () => {
    return (
        <div className={`${style.per_data} ${fStyle.per}`}>
            < PerCell name='ФИО' tag='name' style={style.name} />
            < PerCell name='КУРАТОР' tag='curator' style={style.curator} />
            < PerCell name='ПРОФЕССИЯ' tag='propf' style={style.propf} />
            < PerCell name='НАЦИОНАЛЬНОСТЬ' tag='nation' style={style.nation} />
            < PerCell name='ВОЗРАСТ' tag='age' style={style.age} />
            <div className={`${style.gender} ${fStyle.cell}`} >
                <p> ПОЛ </p>
                М<input type='radio' name='gender' value='M' />
                Ж<input type='radio' name='gender' value='F' />
            </div>
            <div className={`${style.occup} ${fStyle.cell}`}>
                <p> ОБРАЗОВАНИЕ </p>
                <textarea name='occup'/>
            </div> 
        </div>
    )
}

const PerCell = (prop) => {
    return (
        <div className={`${style.cell} ${fStyle.cell} ${prop.style}`}>
            <p> {prop.name} </p>
            <input type='text' name={prop.tag} />
        </div>
    )
}