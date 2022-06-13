import React from 'react';
import { Box, Flex, VStack, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Select, NumberInput, InputGroup, InputLeftAddon } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export default function Register() {
  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
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
      <Box width="60%" minheight="100%" display="flex" alignItems="center">
        <Box m={10} width="100%">
          <Box as="h1" fontSize="2xl" fontWeight="bold" mb={3}>
            <h1>Register Akun</h1>
          </Box>
          <Box as="span" fontSize="m" color="grey">
            Silahkan Masukkan Data Diri Anda
          </Box>
          <Box maxWidth="80%" m={5}>
            <VStack
              spacing={4}
              align='stretch'
            >
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
                  <Button colorScheme="red" width="full" p={5}>
                    Daftar Sekarang
                  </Button>
                  <Box as='p' fontSize='m' color="grey" textAlign="center">
                    Atau Anda Sudah Memiliki Akun
                  </Box>
                  <Button as={Link} to="/login" colorScheme="red" variant="outline" width="full" p={5}>
                    Login Sekarang
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
