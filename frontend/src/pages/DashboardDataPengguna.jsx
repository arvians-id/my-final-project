import {
  Button,
  Table,
  TableCaption,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import MainAppLayout from '../components/layout/MainAppLayout';

export default function DashboardDataPengguna() {
  let num = 1
  let userList = [
    {
      id: 1,
      name: "Irfan Kurniawan",
      type_of_disability: 0,
      role: 1
    },
    {
      id: 2,
      name: "Rahmalina",
      type_of_disability: 2,
      role: 2
    }
  ]

  let disability_type = (type) => {
    if (type === 0) {
      return "None"
    }
    else if (type === 1) {
      return "Tuna Netra"
    }
    else {
      return "Tuna Rungu"
    }
  }

  let role_type = (role) => {
    if (role === 1) {
      return "Guru"
    }
    else {
      return "Siswa"
    }
  }
  return (
    <MainAppLayout>
      <Box m={5} width="full">
        <Stack spacing={6}>
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Manajemen User
            </Box>
            <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
              Manajemen Data User dengan Mudah Dan Cepat
            </Box>
          </Box>
          {/* End Header */}
          {/* Content */}
          <Button variant="solid" colorScheme="green" width="30%">Tambah User</Button>
          <Box>
            <TableContainer>
              <Table variant='striped' colorScheme='blue'>
                <Thead>
                  <Tr>
                    <Th>No</Th>
                    <Th>Nama</Th>
                    <Th>Disabilitas</Th>
                    <Th>Role</Th>
                    <Th>Aksi</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {
                    userList.map((user) => {
                      return (
                        <Tr>
                          <Td>{num++}</Td>
                          <Td>{user.name}</Td>
                          <Td>{disability_type(user.type_of_disability)}</Td>
                          <Td>{role_type(user.role)}</Td>
                          <Td>
                            <Stack direction="row" spacing={3}>
                              <Button variant="solid" colorScheme="blue" size="sm">Edit</Button>
                              <Button variant="solid" colorScheme="red" size="sm">Hapus</Button>
                            </Stack>
                          </Td>
                        </Tr>
                      )
                    })
                  }
                </Tbody>
              </Table>
            </TableContainer>
          </Box>
          {/* End Content */}
        </Stack>
      </Box>
    </MainAppLayout>
  );
}
