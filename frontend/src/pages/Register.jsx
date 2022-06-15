import React from 'react';
import { Box, Flex, VStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Select, NumberInput, InputGroup, InputLeftAddon, Heading, Hide } from '@chakra-ui/react';
import { Link } from 'react-router-dom';
import { useState } from 'react';

export default function Register() {
  const [registerForm, setRegisterForm] = useState({
    fullname: '',
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    gender: '',
    disability: '',
    role: '',
    phoneNumber: '',
    address: '',
    loading: false
  })

  const onChangeRegisterForm = (e) => {
    setRegisterForm({
      ...registerForm,
      [e.target.name]: e.target.value
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
                <Input id="full-name" type="text" maxWidth="full" height={50} placeholder="Full Name" value={registerForm.fullname} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="username" fontWeight="bold">
                  Username
                </FormLabel>
                <Input id="username" type="text" maxWidth="full" height={50} placeholder="User Name" value={registerForm.username} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Email address
                </FormLabel>
                <Input id="email" type="email" maxWidth="full" height={50} placeholder="Masukkan Alamat Email Anda" value={registerForm.email} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Password
                </FormLabel>
                <Input id="password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" value={registerForm.password} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Confirm Password
                </FormLabel>
                <Input id="co-password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" value={registerForm.confirmPassword} onChange={onChangeRegisterForm} />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Gender
                </FormLabel>
                <Select id="gender" placeholder="Select Gender" value={registerForm.gender} onChange={onChangeRegisterForm}>
                  <option>Male</option>
                  <option>Female</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Disabilitas
                </FormLabel>
                <Select id="disability" placeholder="Select Disabilitas" value={registerForm.disability} onChange={onChangeRegisterForm}>
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
                  <Input type="tel" placeholder="phone number" value={registerForm.phoneNumber} onChange={onChangeRegisterForm} />
                </InputGroup>
              </Box>
              <Box>
                <VStack spacing={3} mt={5}>
                  <Button as={Link} to="/login" colorScheme="blue" variant="outline" width="100%" p={5}>
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
