import React from 'react';
import { Box, Flex, VStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Select, NumberInput, InputGroup, InputLeftAddon, Heading, Hide, createStandaloneToast } from '@chakra-ui/react';
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import {API_REGISTER} from '../api/auth'

export default function Register() {
  const navigate = useNavigate();
  const { toast } = createStandaloneToast()
  const [registerForm, setRegisterForm] = useState({
    fullname: '',
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    gender: 0,
    disability: 0,
    role: 2,
    phoneNumber: '',
    address: '',
    birthdate: '',
    loading: false
  })

  const onChangeRegisterForm = (e) => {
    setRegisterForm({
      ...registerForm,
      [e.target.name]: e.target.value
    })
  }



  const disabledButtonRegister = () => {
    if(!registerForm.fullname || !registerForm.email  || !registerForm.password || !registerForm.birthdate || !registerForm.password || !registerForm.phoneNumber) 
        return true
    return false
  }

  const handleSubmitRegister = async (e) => {
      e.preventDefault();
      setRegisterForm({
          ...registerForm,
          loading: true
      })
      const res = await API_REGISTER({
        name: registerForm.fullname,
        username: registerForm.username,
        email: registerForm.email,
        password: registerForm.password,
        role: registerForm.role,
        gender: Number(registerForm.gender),
        type_of_disability: Number(registerForm.disability),
        birthdate: registerForm.birthdate,
      })
      setRegisterForm({
          ...registerForm,
          loading: false
      })
      console.log('res', res)
      if(res.status === 201) {      
          clearRegisterForm();   
          gotoLoginPage();   
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

  const gotoLoginPage = () => {
    toast({
      position: 'bottom',
      title: 'Berhasil Mendaftar Akun.',
      description: 'Anda akan diahrakan kehalaman login dalam 3 detik',
      status: 'success',
      duration: 9000,
      isClosable: false,    
    })
    setTimeout(() => {
      navigate('/login')
    },3000)
  }

  const clearRegisterForm = () => {
    setRegisterForm({
      fullname: '',
      username: '',
      email: '',
      password: '',
      confirmPassword: '',
      gender: 0,
      disability: 0,
      role: 0,
      phoneNumber: '',
      address: '',
      birthdate: '',
      loading: false
    })
  
  }

  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
      <Hide below='md'>
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
      </Hide>
      <Box width="60%" minheight="100%" display="flex" alignItems="center">
        <Box m={10} width="100%">
          <Box textAlign="center" as="h1" fontSize="2xl" fontWeight="bold" mb={3}>
            <Heading as="h2" size="2xl">
              Register Akun
            </Heading>
          </Box>
          <Box as="span" fontSize="m" color="grey">
            Silahkan Masukkan Data Diri Anda
          </Box>
          <Box maxWidth="80%" m={5}>
            <VStack spacing={4} align="stretch">
              <Box>
                <FormLabel htmlFor="full-name" fontWeight="bold">
                  Full Name
                </FormLabel>
                <Input id="fullname" name="fullname" type="text" maxWidth="full" height={50} placeholder="Full Name" value={registerForm.fullname} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="username" fontWeight="bold">
                  Username
                </FormLabel>
                <Input id="username" name="username" type="text" maxWidth="full" height={50} placeholder="User Name" value={registerForm.username} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Email address
                </FormLabel>
                <Input id="email" name="email" type="email" maxWidth="full" height={50} placeholder="Masukkan Alamat Email Anda" value={registerForm.email} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Password
                </FormLabel>
                <Input id="password" type="password" name="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" value={registerForm.password} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Confirm Password
                </FormLabel>
                <Input id="co-password" name="confirmPassword" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" value={registerForm.confirmPassword} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Gender
                </FormLabel>
                <Select id="gender" placeholder="Select Gender" name="gender" value={registerForm.gender} onChange={onChangeRegisterForm}>
                  <option value={1} >Pria</option>
                  <option value={2}>Wanita</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Disabilitas
                </FormLabel>
                <Select id="disability" name="disability" placeholder="Select Disabilitas" value={registerForm.disability} onChange={onChangeRegisterForm}>
                  <option value={1}>Tunanetra</option>
                  <option value={2}>Tunarungu</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Phone Number
                </FormLabel>
                <InputGroup mt={5}>
                  <InputLeftAddon children="+62" />
                  <Input type="tel" placeholder="phone number" name="phoneNumber" value={registerForm.phoneNumber} onChange={onChangeRegisterForm} />
                </InputGroup>
              </Box>
              <Box>
                <FormLabel htmlFor="birthdate" fontWeight="bold">
                  Tanggal Lahir
                </FormLabel>
                <InputGroup mt={5}>
                  <Input type="date" placeholder="Tanggal Lahir" name="birthdate" value={registerForm.birthdate} onChange={onChangeRegisterForm} />
                </InputGroup>
              </Box>
              <Box>
                <VStack spacing={3} mt={5}>
                  <Button disabled={disabledButtonRegister()} onClick={handleSubmitRegister} as={Link} to="/login" colorScheme="blue" loading={registerForm.loading} variant="outline" width="100%" p={5}>
                    Daftar Sekarang
                  </Button>
                </VStack>
              </Box>
            </VStack>
          </Box>
        </Box>
      </Box>
    </Flex>
  );
}
