import React from 'react';
import { Box, Flex, VStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Heading } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export default function Login() {
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
            <VStack spacing={4} align="stretch">
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Email address
                </FormLabel>
                <Input id="email" type="email" maxWidth="full" height={50} placeholder="Masukkan Alamat Email Anda" />
              </Box>
              <Box>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Password
                </FormLabel>
                <Input id="password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" />
              </Box>
              <Box>
                <VStack spacing={3} mt={10}>
                  <Button as={Link} to="/" colorScheme="blue" width="150px" p={5}>
                    Login
                  </Button>
                  <Box as="p" fontSize="m" color="#6A67CE" textAlign="center">
                    Atau Belum Punya Akun?
                  </Box>
                  <Button as={Link} to="/register" colorScheme="blue" variant="outline" width="150px" p={5}>
                    Daftar Sekarang
                  </Button>
                </VStack>
              </Box>
            </VStack>
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
