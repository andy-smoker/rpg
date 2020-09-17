import React, { useState, useEffect } from 'react'

export function MyComponent() {
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState()
    const [components, setComponent] = useState([])

    useEffect(() => {
        fetch("http://localhost:8080/sw/chars").then(
            res => res.json()
        ).then(
            (result) => {
                setIsLoaded(true)
                setComponent(result)
            },
            (error) => {
                setIsLoaded(true)
                setError(error)
            }
        )
    }, [])

    if (error) {
    return <div>Error: {error.message}</div>
    } else if (!isLoaded){
        return <div> Загрузка...</div>
    } else {
        return(
            <table>
                <tr>
                    <th hidden>ID</th><th>Имя</th> <th>Ранг</th> 
                </tr>
                {components.map(component =>(
                    <tr>
                        <td hidden>
                        {component.ID}
                        </td>
                        <td>
                        {component.name}
                        </td>
                        <td>
                        {component.rank}
                        </td>
                        <a href={"profile/#ch?ch="+component.ID} > adout</a>
                    </tr>
                ))}
            </table>
        )
    }
}