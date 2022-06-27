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
import { MdStackedBarChart } from 'react-icons/md'

export default function AdminCourseDetail() {
    let num = 1
    let courseDetail = [
        {
            id: 1,
            name: "Bahasa Indonesia",
            class: "XI Bahasa",
            description: "Pada Kelas Ini akan Mempelajari tentang Menyusun Prosedur, Teks Eksplanasi, Mengelola Informasi dan lainnya",
        }
    ]

    let moduleList = [
        {
            id: 1,
            title: "Pengenalan Teks Prosedur",
            content: "A. Mengonstruksi Informasi dalam Teks Prosedur \n Menunjukkan Pernyataan Umum dalam Suatu Kegiatan \n  Seseorang melakukan suatu kegiatan tentu saja harus memperhatikan langkah-langkah mengerjakannya. Apabila kita akan melakukan pekerjaan, maka harus memahami langkah-langkahnya agar hasil kegiatan tersebut berhasil dengan baik. Ciri teks prosedur yaitu terdapat bagian pernyataan umum dan tahapan-tahapan melakukan kegiatan"
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
                            <Box as="h1" fontSize="3xl" fontWeight="semibold">
                                Detail Materi
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
                            <Button variant="solid" colorScheme="green" width="30%">Tambah Siswa</Button>
                            <Button variant="solid" colorScheme="green" width="30%">Tambah Materi</Button>
                            </Stack>
                            <Box>
                                <TableContainer>
                                    <Table variant='striped' colorScheme='blue'>
                                        <Thead>
                                            <Tr>
                                                <Th>No</Th>
                                                <Th>Judul Materi</Th>
                                                <Th>Isi Materi</Th>
                                                <Th>Aksi</Th>
                                            </Tr>
                                        </Thead>
                                        {
                                            <Tbody>
                                                {
                                                    moduleList.map((module) => {
                                                        return (
                                                            <Tr>
                                                                <Td>{num++}</Td>
                                                                <Td>{module.title}</Td>
                                                                <Td>{`${module.content.substring(0, 30)}`}</Td>
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
                                        }
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

