import React from 'react'
import {
    Box,
    Flex,
    VStack,
    FormControl,
    FormLabel,
    FormErrorMessage,
    FormHelperText,
    Input,
    Button,
} from '@chakra-ui/react'
import { Link } from 'react-router-dom'

export default function Login() {
    return (
        <Flex minHeight='100vh' width='full' flexDirection="row">
            <Box width="60%" minheight="100%" display="flex" alignItems="center">
                <Box m={10} width="100%">
                    <Box as='h1' fontSize='2xl' fontWeight='bold' mb={3}>
                        <h1>Login Akun</h1>
                    </Box>
                    <Box as='span' fontSize='m' color="grey">Silahkan Masukkan Email Dan Password Anda</Box>
                    <Box maxWidth="80%" m={5}>
                        <VStack spacing={4} align='stretch'>
                            <Box>
                                <FormLabel htmlFor='email' fontWeight='bold'>Email address</FormLabel>
                                <Input id='email' type='email' maxWidth="full" height={50} placeholder='Masukkan Alamat Email Anda' />
                            </Box>
                            <Box>
                                <FormLabel htmlFor='email' fontWeight='bold'>Password</FormLabel>
                                <Input id='password' type='password' colorScheme="red" maxWidth="full" height={50} placeholder='Masukkan Password Anda' />
                            </Box>
                            <Box>
                                <VStack spacing={3} mt={5}>
                                    <Button as={Link} to="/" colorScheme="red" width="full" p={5}>
                                        Login
                                    </Button>
                                    <Box as='p' fontSize='m' color="grey" textAlign="center">
                                        Atau Anda Sudah Memiliki Akun
                                    </Box>
                                    <Button as={Link} to="/register" colorScheme="red" variant="outline" width="full" p={5}>
                                        Daftar Sekarang
                                    </Button>
                                </VStack>
                            </Box>
                        </VStack>
                    </Box>
                </Box>
            </Box>
            <Box width="40%" height="100vh" bg="red.100" display="flex" alignItems="center" position="sticky" top="0" left="0" overflowY="auto">
                <Box m={10} width="100%">
                    <Box as='h1' fontSize='4xl' fontWeight='bold' mb={3}>
                        TEENAGER
                    </Box>
                    <Box as='span' fontSize='m'>
                        TEmpat mENgajar dan berbAGi kecERdasan
                    </Box>
                    <Box as='p' fontSize='m' color="grey">

                    </Box>
                </Box>
            </Box>
        </Flex>
    )
}
