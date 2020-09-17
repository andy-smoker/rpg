import React from 'react'
import {Switch, Route, Redirect} from 'react-router-dom'

import { General } from './Pages/General'
import { Profile } from './Pages/Profile'
import { Register } from './Pages/Register'
import { CreateChar } from './Pages/CreateChar'


export const useRoutes = isAuth => {
    if (isAuth) {
        return (
            <Switch>
                <Route path="/profile" exact>
                    <Profile />
                </Route>
                <Route path="/create" exact>
                    <CreateChar />
                </Route>
                <Route path="/profile/ch:id">
                    <CreateChar />
                </Route>
                <Redirect to="/profile" />
            </Switch>
        )
    }
    return (
        <Switch>
            <Route path="/" exact>
                    <General />
                </Route>
            <Route path="/reg" exact>
                    <Register />
            </Route>
            <Redirect to="/" />
        </Switch>
    )
}