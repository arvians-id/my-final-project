import React, { useState } from 'react';
import { Box, Flex, VStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Select, NumberInput, InputGroup, InputLeftAddon, Heading } from '@chakra-ui/react';
import { Link, useNavigate } from 'react-router-dom';
import { API_REGISTER } from '../api/auth';
import { createStandaloneToast } from '@chakra-ui/toast'

export default function Register() {
  const { toast } = createStandaloneToast()
  const [registerForm, setRegisterForm] = useState({
    name: "",
    username: "",
    email: "",
    password: "",
    role: 2,
    gender: 0,
    type_of_disability: 0,
    birthdate: "",
    loading: false
  })
  const navigate = useNavigate()

  const disabledButtonRegister = () => {
    if(!registerForm.email  || !registerForm.password || !registerForm.gender || !registerForm.type_of_disability || !registerForm.birthdate ) 
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
        name: registerForm.name,
        username: registerForm.username,
        email: registerForm.email,
        password: registerForm.password,
        role: registerForm.role,
        gender: registerForm.gender,
        type_of_disability: registerForm.type_of_disability,
        birthdate: registerForm.birthdate,
      })
      setRegisterForm({
          ...registerForm,
          loading: false
      })
      if(res.status === 200) {      
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

  const onChangeRegisterForm = (e) => {
      setRegisterForm({
          ...registerForm,
          [e.target.name]: e.target.value
      })
  }

  const clearRegisterForm = () => {
      setRegisterForm({
        name: "",
        username: "",
        email: "",
        password: "",
        role: 0,
        gender: 0,
        type_of_disability: 0,
        birthdate: "",
        loading: false
      })
  }
   
  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
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
                <FormLabel htmlFor="first-name" fontWeight="bold">
                  First Name
                </FormLabel>
                <Input id="first-name" type="text" maxWidth="full" height={50} placeholder="First Name" />
              </Box>
              <Box>
                <FormLabel htmlFor="last-name" fontWeight="bold">
                  Last Name
                </FormLabel>
                <Input id="last-name" type="text" maxWidth="full" height={50} placeholder="Last Name" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Email address
                </FormLabel>
                <Input id="email" type="email" maxWidth="full" height={50} placeholder="Masukkan Alamat Email Anda" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Repeat Email
                </FormLabel>
                <Input id="re-email" type="email" maxWidth="full" height={50} placeholder="Repeat Email Anda" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Password
                </FormLabel>
                <Input id="password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Confirm Password
                </FormLabel>
                <Input id="co-password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Gender
                </FormLabel>
                <Select id="gender" placeholder="Select Gender">
                  <option>Male</option>
                  <option>Female</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Disabilitas
                </FormLabel>
                <Select id="gender" placeholder="Select Disabilitas">
                  <option>Tunanetra</option>
                  <option>Tunarungu</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Phone Number
                </FormLabel>
                <InputGroup mt={5}>
                  <InputLeftAddon children="+62" />
                  <Input type="tel" placeholder="phone number" />
                </InputGroup>
              </Box>
              <Box>
                <VStack spacing={3} mt={5}>
                  <Button as={Link} to="/login" colorScheme="blue" variant="outline" width="150px" p={5}>
                    Login
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
