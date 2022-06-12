import React from 'react';
import { Box, Flex, FormControl, FormLabel, FormErrorMessage, FormHelperText, Input, Button, Select, NumberInput, InputGroup, InputLeftAddon } from '@chakra-ui/react';

export default function Register() {
  return (
    <Flex minHeight="100vh" width="full" flexDirection="row">
      <Box width="40%" minheight="100%" bg="red.100" display="flex" alignItems="center">
        <Box m={10} width="100%">
          <Box as="h1" fontSize="4xl" fontWeight="bold" mb={3}>
            TEENAGER
          </Box>
          <Box as="span" fontSize="m">
            TEmpat mENgajar dan berbAGi kecERdasan
          </Box>
          <Box as="p" fontSize="m" color="grey"></Box>
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
            <FormControl mt={5}>
              <FormLabel htmlFor="first-name" fontWeight="bold">
                First Name
              </FormLabel>
              <Input id="first-name" type="text" maxWidth="full" height={50} placeholder="First Name" />
              <FormLabel htmlFor="last-name" fontWeight="bold">
                Last Name
              </FormLabel>
              <Input id="last-name" type="text" maxWidth="full" height={50} placeholder="Last Name" />
              <FormLabel htmlFor="email" fontWeight="bold">
                Email address
              </FormLabel>
              <Input id="email" type="email" maxWidth="full" height={50} placeholder="Masukkan Alamat Email Anda" />
              <FormLabel htmlFor="email" fontWeight="bold">
                Repeat Email
              </FormLabel>
              <Input id="re-email" type="email" maxWidth="full" height={50} placeholder="Repeat Email Anda" />
              <FormLabel htmlFor="email" fontWeight="bold">
                Password
              </FormLabel>
              <Input id="password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" />
              <FormLabel htmlFor="email" fontWeight="bold">
                Confirm Password
              </FormLabel>
              <Input id="co-password" type="password" colorScheme="red" maxWidth="full" height={50} placeholder="Masukkan Password Anda" />
              <FormLabel htmlFor="email" fontWeight="bold">
                Gender
              </FormLabel>
              <Select id="gender" placeholder="Select Gender">
                <option>Male</option>
                <option>Female</option>
              </Select>
              <FormLabel htmlFor="email" fontWeight="bold">
                Disabilitas
              </FormLabel>
              <Select id="gender" placeholder="Select Disabilitas">
                <option>Tunanetra</option>
                <option>Tunarungu</option>
              </Select>
              <InputGroup mt={5}>
                <FormLabel htmlFor="email" fontWeight="bold">
                  Phone Number
                </FormLabel>
                <InputLeftAddon children="+62" />
                <Input type="tel" placeholder="phone number" />
              </InputGroup>
            </FormControl>
            <Button colorScheme="red" width="full" mt={4} p={5}>
              Masuk Sekarang
            </Button>
            <Button colorScheme="red" variant="outline" width="full" mt={4} p={5}>
              Daftar Akun
            </Button>
          </Box>
        </Box>
      </Box>
    </Flex>
  );
}
