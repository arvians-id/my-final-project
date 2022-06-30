import {
  Box,
  Button,
  Flex,
  FormLabel,
  Heading,
  Input,
  VStack,
} from '@chakra-ui/react';
import { createStandaloneToast } from '@chakra-ui/toast';
import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { API_LOGIN } from '../api/auth';
import useStore from '../provider/zustand/store';
import { adapterUserToFE } from '../utils/adapterToFE';
import { localSaveToken } from '../utils/token';
import { useHoverTextToSpech } from '../hooks/useHover';

let utterThis;
var synth = window.speechSynthesis;
export default function Login() {
  const navigate = useNavigate();
  const setUser = useStore((state) => state.setUser);
  const { toast } = createStandaloneToast();
  const [refTitle, isHoverTitle] = useHoverTextToSpech('Login Akun');
  const [refInfo, hoverInfo] = useHoverTextToSpech(
    'Silahkan Masukkan Email Dan Password Anda'
  );
  const [refFieldEmail, isEmailField] = useHoverTextToSpech('Email address');
  const [refPasswordField, isPasswordField] = useHoverTextToSpech('Password');
  const [refButtonLogin, isButtonLogin] = useHoverTextToSpech('Tombol Login');
  const [infoLogin, isInfoLogin] = useHoverTextToSpech(
    'Atau Anda Sudah Memiliki Akun'
  );
  const [inforTeenager, setInforTeenager] = useHoverTextToSpech('Teenager');
  const [infoDetailTeenager, setInfoDetailTeenager] = useHoverTextToSpech(
    'Tempat mengajar dan berbagi kecerdasan'
  );
  const [refButtonRegister, isButtonRegister] =
    useHoverTextToSpech('Tombol Daftar');
  const [loginForm, setLoginForm] = useState({
    email: '',
    password: '',
    loading: false,
  });

  const disabledButtonLogin = () => {
    if (!loginForm.email || !loginForm.password) return true;
    return false;
  };

  const handleSubmitLogin = async (e) => {
    e.preventDefault();
    setLoginForm({
      ...loginForm,
      loading: true,
    });
    const res = await API_LOGIN({
      email: loginForm.email,
      password: loginForm.password,
    });
    setLoginForm({
      ...loginForm,
      loading: false,
    });
    if (res.status === 200 && res.data) {
      localSaveToken(res.data.token);
      utterThis = new SpeechSynthesisUtterance('Login Akun Berhasil');
      synth.speak(utterThis);
      setUser(adapterUserToFE(res.data.data));
      clearLoginForm();
      navigate('/');
      utterThis = new SpeechSynthesisUtterance(
        'Password atau Email tidak ditemukan'
      );
      synth.speak(utterThis);
    } else if (res.status !== 500) {
      toast({
        position: 'bottom',
        title: 'Error Login.',
        description: 'Password atau Email tidak ditemukan',
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    } else {
      utterThis = new SpeechSynthesisUtterance('Login Akun Gagal');
      synth.speak(utterThis);
      toast({
        position: 'bottom',
        title: 'Error Login.',
        description: res.messsage,
        status: 'error',
        duration: 3000,
        isClosable: true,
      });
    }
  };

  const onChangeLoginForm = (e) => {
    setLoginForm({
      ...loginForm,
      [e.target.name]: e.target.value,
    });
  };

  const clearLoginForm = () => {
    setLoginForm({
      email: '',
      password: '',
      loading: false,
    });
  };

  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
      <Box width="60%" minheight="100%" display="flex" alignItems="center">
        <Box m={10} width="100%">
          <Box
            textAlign="center"
            as="h1"
            fontSize="2xl"
            fontWeight="bold"
            mb={3}
          >
            <Heading ref={refTitle} as="h2" size="2xl">
              Login Akun
            </Heading>
          </Box>
          <Box as="span" ref={refInfo} fontSize="m" color="grey">
            Silahkan Masukkan Email Dan Password Anda
          </Box>
          <Box maxWidth="80%" m={5}>
            <form onSubmit={handleSubmitLogin}>
              <VStack spacing={4} align="stretch">
                <Box>
                  <FormLabel
                    ref={refFieldEmail}
                    htmlFor="email"
                    fontWeight="bold"
                  >
                    Email address
                  </FormLabel>
                  <Input
                    onChange={onChangeLoginForm}
                    value={loginForm.email}
                    name="email"
                    id="email"
                    type="email"
                    maxWidth="full"
                    height={50}
                    placeholder="Masukkan Alamat Email Anda"
                  />
                </Box>
                <Box>
                  <FormLabel
                    ref={refPasswordField}
                    htmlFor="email"
                    fontWeight="bold"
                  >
                    Password
                  </FormLabel>
                  <Input
                    onChange={onChangeLoginForm}
                    value={loginForm.password}
                    name="password"
                    id="password"
                    type="password"
                    colorScheme="red"
                    maxWidth="full"
                    height={50}
                    placeholder="Masukkan Password Anda"
                  />
                </Box>
                <Box>
                  <VStack spacing={3} mt={5}>
                    <Button
                      ref={refButtonLogin}
                      isLoading={loginForm.loading}
                      disabled={disabledButtonLogin()}
                      onClick={handleSubmitLogin}
                      colorScheme="red"
                      width="full"
                      p={5}
                      type="submit"
                    >
                      Login
                    </Button>
                    <Box
                      ref={infoLogin}
                      as="p"
                      fontSize="m"
                      color="grey"
                      textAlign="center"
                    >
                      Atau Anda Sudah Memiliki Akun
                    </Box>
                    <Button
                      as={Link}
                      to="/register"
                      colorScheme="red"
                      variant="outline"
                      width="full"
                      p={5}
                      ref={refButtonRegister}
                    >
                      Daftar Sekarang
                    </Button>
                  </VStack>
                </Box>
              </VStack>
            </form>
          </Box>
        </Box>
      </Box>
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
            ref={inforTeenager}
            as="h1"
            fontSize="6xl"
            fontWeight="bold"
            mb={3}
            color="#EEF3D2"
          >
            Teenager
          </Box>
          <Box as="span" fontSize="lg" ref={infoDetailTeenager} color="#EEF3D2">
            Tempat mengajar dan berbagi kecerdasan
          </Box>
          <Box as="p" fontSize="m" color="grey"></Box>
        </Box>
      </Box>
    </Flex>
  );
}
