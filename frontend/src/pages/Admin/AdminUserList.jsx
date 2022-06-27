import React from 'react'
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
} from '@chakra-ui/react'
import Navbar from '../components/Navbar'
import Sidebar from '../components/Sidebar'

export default function AdminUserList() {
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
        if(type === 0) {
            return "None"
        }
        else if(type === 1) {
            return "Tuna Netra"
        }
        else {
            return "Tuna Rungu"
        }
    }

    let role_type = (role) => {
        if(role === 1) {
            return "Guru"
        }
        else {
            return "Siswa"
        }
    }
    return (
        <>
            {/* Navbar */}
            <Navbar />
            <Flex direction="row" justifyContent="flex-start" alignItems="flex-start" top="30">
                {/* Sidebar */}
                <Flex width="20%" minHeight="100vh" bgColor="grey.100" boxShadow='md' position="fixed" left="0" top="20" overflowY="auto">
                    <Sidebar />
                </Flex>
                {/* End Sidebar */}
                {/* Main */}
                <Flex width="80%" minHeight="90vh" bg="white" position="sticky" left="80" marginTop={20}>
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
                                                                    <Button variant="solid" colorScheme="blue" size="sm">Ganti Role</Button>
                                                                    <Button variant="solid" colorScheme="red" size="sm">Hapus</Button>
                                                                </Stack>
                                                            </Td>
                                                        </Tr>
                                                    )
                                                })
                                            }
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
            </Flex>
        </ >
    )
}

