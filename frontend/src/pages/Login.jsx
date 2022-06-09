import React from 'react'
import { Box, Button} from '@chakra-ui/react'
import { Link } from 'react-router-dom'

export default function Login() {
    return (
        <div>
            <Box m={10}>
                <h1>Selamat Datang Di Halaman Login</h1>
                <Box mr={20} mt={10}>
                    <Link to="/home"><Button variantColor="teal"m={3}>Home</Button></Link>
                    <Link to="/"><Button color={'red'}>Login</Button></Link>
                </Box>
            </Box>
        </div>
    )
}
