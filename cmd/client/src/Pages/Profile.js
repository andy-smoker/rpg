import React from 'react'
import { Header } from './Components/Header'
import {MyComponent as CharShitTable } from '../JSON'

export function Profile(){
    return(
        <div>
            <Header />
            <label>Таблица персонажей</label>
            
            <CharShitTable />
        </div>
    )
}