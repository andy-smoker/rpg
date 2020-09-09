import React from'react'
import { Header } from './Components/Header'

export function CreateChar() {
    return(
        <div>
            < Header />
            <p> Page for cratin charshit</p>
            <form>
            <p>Имя <input /></p>
            <p>Концепт <input /></p>
            <p>Внешность <input /></p>
            <p>Ранг <select>
                <option value="newbie">новичек</option>
                <option value="middle">бывалый</option>
                <option value="veteran">ветеран</option>
                </select></p>
            <p><textarea>о персонаже</textarea></p>
            <button type="submit"> создать </button>
            </form>
            
        </div>
    )
}