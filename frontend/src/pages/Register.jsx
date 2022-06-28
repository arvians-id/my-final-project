import React from 'react';
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
  Select,
  NumberInput,
  InputGroup,
  InputLeftAddon,
  Heading,
  Hide,
  createStandaloneToast,
  Text,
} from '@chakra-ui/react';
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { API_REGISTER } from '../api/auth';
import {
  strongRegexPatternPassword,
  mediumRegexPatternPassword,
  usernameRegexPattern,
} from '../utils/reqex';
import { checkIsValidUsername } from '../utils/user';

export default function Register() {
  const navigate = useNavigate();
  const { toast } = createStandaloneToast();
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
    loading: false,
  });

  const onChangeRegisterForm = (e) => {
    setRegisterForm({
      ...registerForm,
      [e.target.name]: e.target.value,
    });
  };

  const onChangeGender = (e) => {
    setRegisterForm({
      ...registerForm,
      gender: Number(e.target.value),
    });
  };

  const onChangeDisabilitas = (e) => {
    setRegisterForm({
      ...registerForm,
      disability: Number(e.target.value),
    });
  };

  const checkIsValidRegister = () => {
    if (
      !registerForm.fullname ||
      !registerForm.email ||
      !registerForm.password ||
      !registerForm.birthdate ||
      !registerForm.password ||
      !registerForm.phoneNumber ||
      registerForm.gender === 0 ||
      registerForm.disability === 0 ||
      registerForm.password !== registerForm.confirmPassword ||
      !checkIsValidPassword() ||
      !checkIsValidUsername(registerForm.username)
    )
      return true;
    return false;
  };

  const handleSubmitRegister = async (e) => {
    e.preventDefault();
    if (!checkIsValidRegister()) {
      setRegisterForm({
        ...registerForm,
        loading: true,
      });
      const res = await API_REGISTER({
        name: registerForm.fullname,
        username: registerForm.username,
        email: registerForm.email,
        password: registerForm.password,
        role: registerForm.role,
        gender: Number(registerForm.gender),
        disability:
          Number(registerForm.disability) === 3
            ? 0
            : Number(registerForm.disability),
        birthdate: registerForm.birthdate,
      });
      setRegisterForm({
        ...registerForm,
        loading: false,
      });
      if (res.status === 201) {
        clearRegisterForm();
        gotoLoginPage();
      } else {
        toast({
          position: 'bottom',
          title: 'Error Login.',
          description: res.message,
          status: 'error',
          duration: 3000,
          isClosable: true,
        });
      }
    }
  };

  const gotoLoginPage = () => {
    toast({
      position: 'bottom',
      title: 'Berhasil Mendaftar Akun.',
      description: 'Anda akan diahrakan kehalaman login dalam 3 detik',
      status: 'success',
      duration: 3000,
      isClosable: false,
    });
    setTimeout(() => {
      navigate('/login');
    }, 3000);
  };

  const clearRegisterForm = () => {
    setRegisterForm({
      fullname: '',
      username: '',
      email: '',
      password: '',
      confirmPassword: '',
      gender: 0,
      disability: -1,
      role: 0,
      phoneNumber: '',
      address: '',
      birthdate: '',
      loading: false,
    });
  };

  const checkIsValidPassword = () => {
    if (
      strongRegexPatternPassword.test(registerForm.password) ||
      mediumRegexPatternPassword.test(registerForm.password)
    ) {
      return true;
    }
    return false;
  };

  const renderPasswordStatus = () => {
    if (registerForm.password) {
      if (strongRegexPatternPassword.test(registerForm.password)) {
        return (
          <Box
            bgColor="green.600"
            p="1"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
          >
            <Text color="white">Password sangat kuat</Text>
          </Box>
        );
      } else if (mediumRegexPatternPassword.test(registerForm.password)) {
        return (
          <Box
            bgColor="orange.600"
            p="1"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
          >
            <Text color="white">Password cukup kuat</Text>
          </Box>
        );
      } else {
        return (
          <Box
            bgColor="red.600"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
            p="1"
          >
            <Text color="white">Password lemah</Text>
          </Box>
        );
      }
    }
  };

  const renderUsernameStatus = () => {
    if (registerForm.username) {
      if (
        usernameRegexPattern.test(registerForm.username) &&
        registerForm.username.length > 2
      ) {
        return (
          <Box
            bgColor="green.600"
            p="1"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
          >
            <Text color="white">Username valid</Text>
          </Box>
        );
      } else {
        return (
          <Box
            bgColor="red.600"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
            p="1"
          >
            <Text color="white">
              Minimal 3 karakter terdiri dari huruf besar, kecil, angka,
              karakter . _
            </Text>
          </Box>
        );
      }
    }
  };

  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
      <Hide below="md">
        <Box
          width="40%"
          height="100vh"
          bg="#6A67CE"
          display="flex"
          alignItems="center"
          position="sticky"
          top="0"
          left="0"
          overflowY="auto"
        >
          <Box m={10} width="100%">
            <Box
              as="h1"
              fontSize="6xl"
              fontWeight="bold"
              mb={3}
              color="#EEF3D2"
            >
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
          <Box
            textAlign="center"
            as="h1"
            fontSize="2xl"
            fontWeight="bold"
            mb={3}
          >
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
                <Input
                  id="fullname"
                  name="fullname"
                  type="text"
                  maxWidth="full"
                  height={50}
                  placeholder="Full Name"
                  value={registerForm.fullname}
                  onChange={onChangeRegisterForm}
                />
              </Box>
              <Box>
                <FormLabel htmlFor="username" fontWeight="bold">
                  Username
                </FormLabel>
                <Input
                  id="username"
                  name="username"
                  type="text"
                  maxWidth="full"
                  height={50}
                  placeholder="User Name"
                  value={registerForm.username}
                  onChange={onChangeRegisterForm}
                />
                {renderUsernameStatus()}
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Email address
                </FormLabel>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  maxWidth="full"
                  height={50}
                  placeholder="Masukkan Alamat Email Anda"
                  value={registerForm.email}
                  onChange={onChangeRegisterForm}
                />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Password
                </FormLabel>
                <Input
                  id="password"
                  type="password"
                  name="password"
                  colorScheme="red"
                  maxWidth="full"
                  height={50}
                  placeholder="Masukkan Password Anda"
                  value={registerForm.password}
                  onChange={onChangeRegisterForm}
                />
                {renderPasswordStatus()}
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Confirm Password
                </FormLabel>
                <Input
                  id="co-password"
                  name="confirmPassword"
                  type="password"
                  colorScheme="red"
                  maxWidth="full"
                  height={50}
                  placeholder="Masukkan Password Anda"
                  value={registerForm.confirmPassword}
                  onChange={onChangeRegisterForm}
                />
                {registerForm.password &&
                  (registerForm.password !== registerForm.confirmPassword ? (
                    <Box
                      bgColor="red.600"
                      borderBottomLeftRadius="4px"
                      borderBottomRightRadius="4px"
                      p="1"
                    >
                      <Text color="white">Password tidak sama</Text>
                    </Box>
                  ) : (
                    <Box
                      bgColor="green.600"
                      borderBottomLeftRadius="4px"
                      borderBottomRightRadius="4px"
                      p="1"
                    >
                      <Text color="white">Password sama</Text>
                    </Box>
                  ))}
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Gender
                </FormLabel>
                <Select
                  id="gender"
                  placeholder="Select Gender"
                  name="gender"
                  value={registerForm.gender}
                  onChange={onChangeGender}
                >
                  <option value={1}>Pria</option>
                  <option value={2}>Wanita</option>
                </Select>
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Disabilitas
                </FormLabel>
                <Select
                  id="disability"
                  name="disability"
                  placeholder="Select Disabilitas"
                  value={registerForm.disability}
                  onChange={onChangeDisabilitas}
                  required
                >
                  <option value={3}>None</option>
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
                  <Input
                    type="tel"
                    placeholder="phone number"
                    name="phoneNumber"
                    value={registerForm.phoneNumber}
                    onChange={onChangeRegisterForm}
                  />
                </InputGroup>
              </Box>
              <Box>
                <FormLabel htmlFor="birthdate" fontWeight="bold">
                  Tanggal Lahir
                </FormLabel>
                <InputGroup mt={5}>
                  <Input
                    type="date"
                    placeholder="Tanggal Lahir"
                    name="birthdate"
                    value={registerForm.birthdate}
                    onChange={onChangeRegisterForm}
                  />
                </InputGroup>
              </Box>
              <Box>
                <VStack spacing={3} mt={5}>
                  <Button
                    disabled={checkIsValidRegister()}
                    onClick={handleSubmitRegister}
                    colorScheme="blue"
                    loading={registerForm.loading}
                    variant="outline"
                    width="100%"
                    p={5}
                    type="submit"
                  >
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
