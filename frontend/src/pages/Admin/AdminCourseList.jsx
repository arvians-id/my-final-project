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
    useDisclosure,
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    ModalCloseButton,
    Input,
} from '@chakra-ui/react'
import Navbar from '../components/Navbar'
import Sidebar from '../components/Sidebar'

export default function AdminCourseList() {
    const { isOpen, onOpen, onClose } = useDisclosure()
    let num = 1
    let courseList = [
        {
            id: 1,
            name: "Bahasa Indonesia",
            class: "XI Bahasa",
        },
        {
            id: 2,
            name: "Matematika",
            class: "X IPA",
        },
    ]

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
                                    Manajemen Mata Pelajaran
                                </Box>
                                <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                                    Kelola Data Mata Pelajaran Anda
                                </Box>
                            </Box>
                            {/* End Header */}
                            {/* Content */}
                            <Button variant="solid" colorScheme="green" width="30%" onClick={onOpen}>Tambah Mata Pelajaran</Button>
                            <Box>
                                <TableContainer>
                                    <Table variant='striped' colorScheme='blue'>
                                        <Thead>
                                            <Tr>
                                                <Th>No</Th>
                                                <Th>Mata Pelajaran</Th>
                                                <Th>Kelas</Th>
                                                <Th>Aksi</Th>
                                            </Tr>
                                        </Thead>
                                        <Tbody>
                                            {
                                                courseList.map((course) => {
                                                    return (
                                                        <Tr>
                                                            <Td>{num++}</Td>
                                                            <Td>{course.name}</Td>
                                                            <Td>{course.class}</Td>
                                                            <Td>
                                                                <Stack direction="row" spacing={3}>
                                                                    <Button variant="solid" colorScheme="teal" size="sm">Detail</Button>
                                                                    <Button variant="solid" colorScheme="blue" size="sm" onClick={onOpen}>Edit</Button>
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
                </Flex>
                {/* End main */}
                {/* Modal */}
                <Modal isOpen={isOpen} onClose={onClose}>
                    <ModalOverlay />
                    <ModalContent>
                        <ModalHeader>Tambah Data Mata Pelajaran</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            <Stack direction="column">
                                <Text as="h3" fontSize="md" fontWeight="semibold">
                                    Nama Mata Pelajaran
                                </Text>
                                <Input placeholder='Masukkan Nama Mata Pelajaran' />
                                <Text as="h3" fontSize="md" fontWeight="semibold">
                                    Kelas
                                </Text>
                                <Input placeholder='Masukkan Kelas' />
                            </Stack>
                        </ModalBody>

                        <ModalFooter>
                            <Button colorScheme='blue' mr={3} onClick={onClose}>
                                Tambah Mata Pelajaran
                            </Button>
                            <Button variant='ghost' onClick={onClose}>Cancel</Button>
                        </ModalFooter>
                    </ModalContent>
                </Modal>
                {/* End Modal */}
            </Flex>
        </ >
    )
}

