import React from 'react'
import { Button, ButtonGroup } from '@chakra-ui/react'
import { Link, Route, Routes } from 'react-router-dom'
import Home from './Home'

export default function Login() {
    return (
        <div>
            <h1>Selamat Datang Di Halaman Login</h1>
            <Routes>
                <Route path="/" element={Login}>
                    <Button as={Link}>Home</Button>
                </Route>
                <Route path="/home" element={Home}>
                    <Button as={Link} to="/home">Home</Button>
                </Route>
            </Routes>
        </div>
    )
}
