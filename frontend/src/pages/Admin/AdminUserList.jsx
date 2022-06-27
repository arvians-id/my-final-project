import React, { useEffect } from 'react';
import {
  Box,
  Flex,
  Stack,
  HStack,
  Text,
  Spacer,
  Button,
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
  Select,
} from '@chakra-ui/react';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_GET_LIST_USER } from '../../api/user';

export default function AdminUserList() {
  a;
  const [listAllUser, setListAllUser] = useState([]);
  let num = 1;
  let userList = [
    {
      id: 1,
      name: 'Irfan Kurniawan',
      type_of_disability: 0,
      role: 1,
    },
    {
      id: 2,
      name: 'Rahmalina',
      type_of_disability: 2,
      role: 2,
    },
  ];

  let disability_type = (type) => {
    if (type === 0) {
      return 'None';
    } else if (type === 1) {
      return 'Tuna Netra';
    } else {
      return 'Tuna Rungu';
    }
  };

  let role_type = (role) => {
    if (role === 1) {
      return 'Guru';
    } else {
      return 'Siswa';
    }
  };

  const getAllUsers = async () => {
    const res = await API_GET_LIST_USER();
    if (res.status === 200) {
      setListAllUser(res.data.data ?? []);
    }
  };

  const updateRole = (type) => {
    axiosWithToken()
      .put(`${BASE_URL}/api/users/roleupdate/${type}`)
      .then((res) => {
        if (res.status === 200) {
          toast({
            status: 'success',
            title: 'Berhasil',
            description: 'Berhasil update role',
          });
          getAllUsers();
        } else {
          toast({
            status: 'error',
            title: 'Gagal',
            description: 'Gagal update role',
          });
        }
      });
  };

  const deleteUser = (id) => {
    axiosWithToken()
      .delete(`${BASE_URL}/api/users/${id}`)
      .then((res) => {
        if (res.status === 200) {
          toast({
            status: 'success',
            title: 'Berhasil',
            description: 'Berhasil update role',
          });
          getAllUsers();
        } else {
          toast({
            status: 'error',
            title: 'Gagal',
            description: 'Gagal update role',
          });
        }
      });
  };

  useEffect(() => {
    getAllUsers();
  }, []);

  return (
    <MainAppLayout>
      {/* Main */}
      <Flex
        width="80%"
        minHeight="90vh"
        bg="white"
        position="sticky"
        left="80"
        marginTop={20}
      >
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
            <Box>
              <TableContainer>
                <Table variant="striped" colorScheme="blue">
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
                    {listAllUser.map((user) => {
                      return (
                        <Tr>
                          <Td>{num++}</Td>
                          <Td>{user.name}</Td>
                          <Td>{disability_type(user.type_of_disability)}</Td>
                          <Td>{role_type(user.role)}</Td>
                          <Td>
                            <Stack direction="row" spacing={3}>
                              <Button
                                variant="solid"
                                colorScheme="blue"
                                size="sm"
                              >
                                Ganti Role
                              </Button>
                              <Select
                                id="gender"
                                placeholder="Select Gender"
                                name="gender"
                                onChange={updateRole}
                                value={user.role}
                              >
                                <option value={1}>Guru</option>
                                <option value={2}>Siswa</option>
                              </Select>
                              <Button
                                variant="solid"
                                colorScheme="red"
                                size="sm"
                                onClick={() => deleteUser(user.id)}
                              >
                                Hapus
                              </Button>
                            </Stack>
                          </Td>
                        </Tr>
                      );
                    })}
                    {/* <Tr>
                                                <Td>Irfan Kurniawan</Td>
                                                <Td>None</Td>
                                                <Td>Admin</Td>
                                                <Td>
                                                    <Stack direction="row" spacing={3}>
                                                        <Button variant="solid" colorScheme="blue" size="sm">Edit</Button>
                                                        <Button variant="solid" colorScheme="red" size="sm">Hapus</Button>
                                                    </Stack>
                                                </Td>
                                            </Tr> */}
                  </Tbody>
                </Table>
              </TableContainer>
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
      {/* End main */}
    </MainAppLayout>
  );
}
