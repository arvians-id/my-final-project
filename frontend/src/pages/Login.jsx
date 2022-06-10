import React from 'react'
import {
    Box,
    Flex,
    FormControl,
    FormLabel,
    FormErrorMessage,
    FormHelperText,
    Input,
    Button,
} from '@chakra-ui/react'

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
                        <FormControl>
                            <FormLabel htmlFor='email' fontWeight='bold'>Email address</FormLabel>
                            <Input id='email' type='email' maxWidth="full" height={50} placeholder='Masukkan Alamat Email Anda' />
                        </FormControl>
                        <FormControl mt={5} >
                            <FormLabel htmlFor='email' fontWeight='bold'>Password</FormLabel>
                            <Input id='password' type='password' colorScheme="red" maxWidth="full" height={50} placeholder='Masukkan Password Anda' />
                        </FormControl>
                        <Button colorScheme='red' width="full" mt={4} p={5}>Masuk Sekarang</Button>
                        <Button colorScheme='red' variant="outline" width="full" mt={4} p={5}>Daftar Akun</Button>
                    </Box>
                </Box>
            </Box>
            <Box width="40%" minheight="100%" bg="red.100" display="flex" alignItems="center">
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
