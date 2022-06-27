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
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    ModalCloseButton,
    useDisclosure,
    InputLeftElement,
    InputGroup,
    Input,
    Avatar
} from '@chakra-ui/react'
import Navbar from '../components/Navbar'
import Sidebar from '../components/Sidebar'
import { MdStackedBarChart } from 'react-icons/md'
import { SearchIcon } from '@chakra-ui/icons'

export default function AdminAddCourseStudent() {
    const { isOpen, onOpen, onClose } = useDisclosure()

    let num = 1
    let courseDetail = [
        {
            id: 1,
            name: "Bahasa Indonesia",
            class: "XI Bahasa",
            description: "Pada Kelas Ini akan Mempelajari tentang Menyusun Prosedur, Teks Eksplanasi, Mengelola Informasi dan lainnya",
        }
    ]

    let studentList = [
        {
            id: 1,
            name: "Irfan Kurniawan"
        },
        {
            id: 2,
            name: "Amellisa Anggun Oktadilla"
        }
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
                            <Box as="h1" fontSize="3xl" fontWeight="semibold">
                                Daftar Siswa
                            </Box>
                            <Box>
                                {
                                    courseDetail.map((detail) => {
                                        return (
                                            <Stack>
                                                <Box as="h1" fontSize="2xl" fontWeight="semibold">
                                                    {detail.name}
                                                </Box>
                                                <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                                                    {detail.class}
                                                </Box>
                                                <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                                                    {detail.description}
                                                </Box>
                                            </Stack>

                                        )
                                    })
                                }
                            </Box>
                            {/* End Header */}
                            {/* Content */}
                            <Stack direction="row">
                                <Button variant="solid" colorScheme="green" width="30%" onClick={onOpen}>Tambah Siswa</Button>
                            </Stack>
                            <Box>
                                <TableContainer>
                                    <Table variant='striped' colorScheme='blue'>
                                        <Thead>
                                            <Tr>
                                                <Th>No</Th>
                                                <Th>Nama Siswa</Th>
                                                <Th>Aksi</Th>
                                            </Tr>
                                        </Thead>
                                        {
                                            <Tbody>
                                                {
                                                    studentList.map((student, index) => {
                                                        return (
                                                            <Tr key={index}>
                                                                <Td>{num++}</Td>
                                                                <Td>{student.name}</Td>
                                                                <Td>
                                                                    <Stack direction="row" spacing={3}>
                                                                        <Button variant="solid" colorScheme="red" size="sm">Hapus</Button>
                                                                    </Stack>
                                                                </Td>
                                                            </Tr>
                                                        )
                                                    })
                                                }

                                            </Tbody>
                                        }
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
                        <ModalHeader>Tambah Siswa</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            <Box>
                                <InputGroup>
                                    <InputLeftElement
                                        pointerEvents='none'
                                        children={<SearchIcon color='gray.300' />}
                                    />
                                    <Input type='tel' placeholder='Cari' />
                                </InputGroup>
                            </Box>
                            <Box my={4} p={2} border="2px" borderColor='gray.200' borderRadius={5}>
                                <Stack direction="row" alignItems="center">
                                    <Avatar name="Irfan Kurniawan" width={10} height={10} />
                                    <Text>Irfan Kurniawan</Text>
                                </Stack>
                            </Box>
                        </ModalBody>

                        <ModalFooter>
                            <Button colorScheme='blue' mr={3} onClick={onClose}>
                                Tambah Data
                            </Button>
                            <Button onClick={onClose}>Cancel</Button>
                        </ModalFooter>
                    </ModalContent>
                </Modal>
                {/* End Modal */}
            </Flex>
        </ >
    )
}

