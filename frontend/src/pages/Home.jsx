import React from 'react';
import {
    Box,
    Flex,
    Stack,
    HStack,
    MenuButton,
    MenuList,
    MenuItem,
    Menu,
    MenuDivider,
} from '@chakra-ui/react';
import { ChevronDownIcon } from '@chakra-ui/icons';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import CourseCard from '../components/CourseCard';
export default function Home() {
    let courseList = [
        {
            courseTitle: 'Geografi',
            courseClass: 'XI IPS',
            courseTeacher: 'Rahmat Pratama, S.Pd',
            courseDescription: 'Mempelajari tentang Struktur Bumi',
        },
        {
            courseTitle: 'Bahasa Indonesia',
            courseClass: 'X Bahasa',
            courseTeacher: 'Isna Rahmawati, S.Pd',
            courseDescription:
                'Bahasa Indonesia Pelajaran Mengenai Bahasa Indonesia',
        },
        {
            courseTitle: 'Matematika',
            courseClass: 'XII IPA',
            courseTeacher: 'Suci Rahma, S.Pd',
            courseDescription:
                'Mata Pelajaran yang akan membahas Konversi Biner, Aljabar ',
        },
    ];
    return (
        <>
            <Navbar />
            <Flex width="80%" minHeight="90vh" bg="white">
                {/* Main */}
                <Box m={5}>
                    {/* Header */}
                    <Stack spacing={6}>
                        <Box>
                            <Box as="h1" fontSize="2xl" fontWeight="semibold">
                                Selamat Datang Kembali
                            </Box>
                            <Box
                                as="span"
                                fontSize="l"
                                fontWeight="semibold"
                                color="grey"
                            >
                                Lanjutkan Pembelajaran Anda
                            </Box>
                        </Box>
                        <Box as="h1" fontSize="2xl" fontWeight="semibold">
                            Pelajaran Anda
                        </Box>
                        <Box>
                            <Menu>
                                <MenuButton
                                    px={4}
                                    py={2}
                                    transition="all 0.2s"
                                    borderRadius="md"
                                    borderWidth="1px"
                                    _hover={{ bg: 'gray.400' }}
                                    _expanded={{ bg: 'blue.400' }}
                                    _focus={{ boxShadow: 'outline' }}
                                >
                                    File <ChevronDownIcon />
                                </MenuButton>
                                <MenuList>
                                    <MenuItem>New File</MenuItem>
                                    <MenuItem>New Window</MenuItem>
                                    <MenuDivider />
                                    <MenuItem>Open...</MenuItem>
                                    <MenuItem>Save File</MenuItem>
                                </MenuList>
                            </Menu>
                        </Box>
                        <Box alignContent="flex-start">
                            <HStack spacing={8}>
                                {courseList.map((course, index) => {
                                    return (
                                        <CourseCard
                                            key={index}
                                            courseTitle={course.courseTitle}
                                            courseClass={course.courseClass}
                                            courseTeacher={course.courseTeacher}
                                            courseDescription={
                                                course.courseDescription
                                            }
                                        />
                                    );
                                })}
                            </HStack>
                        </Box>
                    </Stack>
                    {/* End Header */}
                </Box>
                {/* End Main */}
            </Flex>
        </>
    );
}
