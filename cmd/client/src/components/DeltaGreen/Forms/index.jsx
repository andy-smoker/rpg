import React from 'react'
import {PerData} from './PerData'
import style from './index.module.css'
import { StatData } from './StatData'
import { PsyData } from './PsyData'
import { SkillSets } from './SkillSets'

export const Form = () =>{
   
    return (
        <div className={style.chsh}>
            <Title text='ПЕРСОНАЛЬНЫЕ ДАННЫЕ' />
            <PerData />
            <Title text='СТАТИСТИЧЕСКАЯ ИНФОРМАЦИЯ' />
            <StatData />
            <Title text='ПСИХОЛОГИЧЕСКАЯ ИНФОРМАЦИЯ' />
            <PsyData />
            <Title text='НАБОР НАВЫКОВ' />
            <SkillSets />
            <button className={style.down}>
              Dowload
            </button>
        </div>
    )
}


const Title = (prop) => {
    return (
        <div className={style.title}>
                <p>{prop.text}</p>
            </div>
    )
}

