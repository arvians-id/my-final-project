import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Login from './pages/Login'

export default function Main() {
    return (
        <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={<Home />} />
        </Routes>
    )
}
