import {
  Box,
  Button,
  Flex,
  FormLabel,
  Heading,
  Input, VStack
} from '@chakra-ui/react'
import { createStandaloneToast } from '@chakra-ui/toast'
import React, { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { API_LOGIN } from '../api/auth'
import useStore from '../provider/zustand/store'
import { localSaveToken } from '../utils/token'

export default function Login() {
  const navigate = useNavigate()
  const setUser = useStore((state) => state.setUser);
  const { toast } = createStandaloneToast()
  const [loginForm, setLoginForm] = useState({
      email: '',
      password: '',
      loading: false
  })

  const disabledButtonLogin = () => {
      if(!loginForm.email  || !loginForm.password) 
          return true
      return false
  }

  const handleSubmitLogin = async (e) => {
      e.preventDefault();
      setLoginForm({
          ...loginForm,
          loading: true
      })
      const res = await API_LOGIN({
          email: loginForm.email,
          password: loginForm.password
      })
      setLoginForm({
          ...loginForm,
          loading: false
      })
      if(res.status === 200) {
          localSaveToken(res.data.data.token)
          setUser(res.data.data)
          clearLoginForm();
          navigate('/')
      } else {
          toast({
              position: 'bottom',
              title: 'Error Login.',
              description: res.message,
              status: 'error',
              duration: 9000,
              isClosable: true,
          })
      }
  }

  const onChangeLoginForm = (e) => {
      setLoginForm({
          ...loginForm,
          [e.target.name]: e.target.value
      })
  }

  const clearLoginForm = () => {
      setLoginForm({
          email: '',
          password: '',
          loading: false
      })
  }
  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
      <Box width="60%" minheight="100%" display="flex" alignItems="center">
        <Box m={10} width="100%">
          <Box textAlign="center" as="h1" fontSize="2xl" fontWeight="bold" mb={3}>
            <Heading as="h2" size="2xl">
              Login Akun
            </Heading>
          </Box>
          <Box as="span" fontSize="m" color="grey">
            Silahkan Masukkan Email Dan Password Anda
          </Box>
          <Box maxWidth="80%" m={5}>
            <form onSubmit={handleSubmitLogin}>
              <VStack spacing={4} align='stretch'>
                <Box>
                  <FormLabel htmlFor='email' fontWeight='bold'>Email address</FormLabel>
                  <Input onChange={onChangeLoginForm} value={loginForm.email} name="email" id='email' type='email' maxWidth="full" height={50} placeholder='Masukkan Alamat Email Anda' />
                </Box>
                <Box>
                  <FormLabel htmlFor='email' fontWeight='bold'>Password</FormLabel>
                  <Input onChange={onChangeLoginForm} value={loginForm.password} name="password" id='password' type='password' colorScheme="red" maxWidth="full" height={50} placeholder='Masukkan Password Anda' />
                </Box>
                <Box>
                  <VStack spacing={3} mt={5}>
                    <Button isLoading={loginForm.loading} disabled={disabledButtonLogin()} onClick={handleSubmitLogin} colorScheme="red" width="full" p={5} type="submit">
                      Login
                    </Button>
                    <Box as='p' fontSize='m' color="grey" textAlign="center">
                      Atau Anda Sudah Memiliki Akun
                    </Box>
                    <Button as={Link} to="/register" colorScheme="blue" variant="outline" width="full" p={5}>
                      Daftar Sekarang
                    </Button>
                  </VStack>
                </Box>
              </VStack>
            </form>                
          </Box>
        </Box>
      </Box>
      <Box width="40%" height="100vh" bg="#6A67CE" display="flex" alignItems="center" position="sticky" top="0" left="0" overflowY="auto">
        <Box m={10} width="100%">
          <Box as="h1" fontSize="6xl" fontWeight="bold" mb={3} color="#EEF3D2">
            Teenager
          </Box>
          <Box as="span" fontSize="lg" color="#EEF3D2">
            Tempat mengajar dan berbagi kecerdasan
          </Box>
          <Box as="p" fontSize="m" color="grey"></Box>
        </Box>
      </Box>
    </Flex>
  );
}
