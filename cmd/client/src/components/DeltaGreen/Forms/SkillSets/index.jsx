import React, { useState } from 'react'
import style from './index.module.css'
import fStyle from '../index.module.css'


export const SkillSets = () => {
    let skills = [
        {name:'Бух.учет', tag:'accounting', con:10},
        {name:'Бдительность', tag:'alertness', con:20},
        {name:'Антропология', tag:'anthrop', con:0},
        {name:'Археология', tag:'archeology', con:0},
        {name:'Искусство', tag:'art', con:0},
        {name:'Артиллерия', tag:'artilery', con:0},
        {name:'Атлетика', tag:'athletics', con:30},
        {name:'Бюрократия', tag:'bureauc', con:10},
        {name:'Компьютерные науки', tag:'compscince', con:0},
        {name:'Ремесло', tag:'craft', con:0},
        {name:'Криминология', tag:'crimlog', con:10},
        {name:'Разрушение', tag:'demolitions', con:0},
        {name:'Маскировка', tag:'disguise', con:10},
        {name:'Уклонение', tag:'dodge', con:30},
        {name:'Вождение', tag:'drive', con:20},
        {name:'Тяжелая техника', tag:'havymachine', con:10},
        {name:'Стрельба', tag:'firearms', con:20},
        {name:'Первая помощ', tag:'firstaid', con:10},
        {name:'Криминалистика', tag:'forensic', con:0},
        {name:'Тяжелое оружие', tag:'havyweapon', con:10},
        {name:'История', tag:'history', con:10},
        {name:'Проницательность', tag:'humint', con:0},
        {name:'Ложь', tag:'law', con:0},
        {name:'Медицина', tag:'midicine', con:0},
        {name:'Холодное оружее', tag:'melee', con:30},
        {name:'Военное дело', tag:'military', con:0},
        {name:'Навигация', tag:'navigate', con:10},
        {name:'Оккультизм', tag:'occult', con:10},
        {name:'Убеждение', tag:'persuade', con:20},
        {name:'Фармацевтика', tag:'farmacy', con:0},
        {name:'Пилотирование', tag:'pilot', con:0},
        {name:'Психотерапия', tag:'psycho', con:10},
        {name:'Верховая езда', tag:'ride', con:10},
        {name:'Наука', tag:'Science', con:0},
        {name:'Поиск', tag:'Search', con:20},
        {name:'Криптография', tag:'sigint', con:0},
        {name:'Скртыность', tag:'stealth', con:10},
        {name:'Хирургия', tag:'surgery', con:0},
        {name:'Выживания', tag:'survival', con:10},
        {name:'Плаванье', tag:'anthrop', con:20},
        {name:'Драка', tag:'combat', con:40},
        {name:'Необъяснимое', tag:'unnatural', con:0},
        
    ]
    return (
        <div className={`${fStyle.skill} ${style.main}`}>
            { skills.map(e => {
                return (
                    <SkillRow text={e.name} tag={e.tag} con={e.con} />
                )
            })}
            <dev>
                <dev className={`${style.row} ${fStyle.cell}`} >
                    <p> Языки и другие навыки </p>
                </dev>
                <SkillRow text={<input type='text' />} />
                <SkillRow text={<input type='text' />} />
            </dev>
        </div>
    )
}

const SkillRow = (prop) => {
    const [val, setVal] = useState(0)
    return (
        <dev className={style.row} >
            <dev className={`${fStyle.cell}`}>
                <p>{prop.text}
                    <input type='text' name={prop.tag} onChange={e => { setVal(e.target.value) }} defaultValue='0'/>
                </p>
            </dev>
            <dev className={`${fStyle.cell}`}>
                <p>{prop.con + parseInt(val)}</p>
            </dev>
        </dev>
    )
}