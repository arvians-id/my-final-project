import React from 'react';
import {
    Box,
    Flex,
    Stack,
    HStack,
    Grid,
    MenuButton,
    MenuList,
    MenuItem,
    Menu,
    MenuDivider,
    Button,
    Spacer,
} from '@chakra-ui/react';
import {
    ChevronDownIcon,
    ArrowBackIcon,
    ArrowForwardIcon,
} from '@chakra-ui/icons';
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
            courseClass: 'X IPA/IPS',
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
        {
            courseTitle: 'Sosiologi',
            courseClass: 'XI IPS',
            courseTeacher: 'Susi Dina, S.Pd',
            courseDescription:
                'Mata Pelajaran yang akan membahas Stratifikasi Sosial ',
        },
        {
            courseTitle: 'Ekonomi',
            courseClass: 'XII IPS',
            courseTeacher: 'Nur Cahyani, S.Pd',
            courseDescription: 'Mata Pelajaran yang akan membahas Akuntansi',
        },
        {
            courseTitle: 'Biologi',
            courseClass: 'XII IPA',
            courseTeacher: 'Ningsi Indriani, S.Pd',
            courseDescription:
                'Mata Pelajaran yang akan membahas Virus dan Bakteri ',
        },
        {
            courseTitle: 'Fisika',
            courseClass: 'XII IPA',
            courseTeacher: 'Indra Darmono, S.Pd',
            courseDescription: 'Mata Pelajaran yang akan membahas Hukum Pascal',
        },
        {
            courseTitle: 'Kimia',
            courseClass: 'XII IPA',
            courseTeacher: 'Vira Monika, S.Pd',
            courseDescription: 'Mata Pelajaran yang akan membahas Biomolekul ',
        },
        {
            courseTitle: 'Agama',
            courseClass: 'XII IPA/IPS',
            courseTeacher: 'Abu Bakri, S.Pd',
            courseDescription: 'Mata Pelajaran yang akan membahas Biomolekul ',
        },
        {
            courseTitle: 'Kimia',
            courseClass: 'XII IPA/IPS',
            courseTeacher: 'Wati Safitri, S.Pd',
            courseDescription: 'Mata Pelajaran yang akan membahas Biomolekul ',
        },
    ];
    return (
        <>
            <Navbar />

            <Flex direction="row" justifyContent="center" alignItems="center">
                <Flex
                    width="20%"
                    minHeight="120vh"
                    bgColor="grey.100"
                    boxShadow="dark-lg"
                >
                    <Sidebar />
                </Flex>
                <Flex width="80%" minHeight="120vh" bg="white">
                    {/* Main */}
                    <Box m={5}>
                        {/* Header */}
                        <Stack spacing={6}>
                            <Box>
                                <Box
                                    as="h1"
                                    fontSize="2xl"
                                    fontWeight="semibold"
                                >
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
                            <Box pos="absolute" top="180" left="1450">
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
                                        Kelas <ChevronDownIcon />
                                    </MenuButton>
                                    <MenuList>
                                        <MenuItem>X SMA</MenuItem>
                                        <MenuItem>XI SMA</MenuItem>
                                        <MenuItem>XII SMA</MenuItem>
                                    </MenuList>
                                </Menu>
                            </Box>
                            <Box alignContent="flex-start">
                                <Grid
                                    spacing={8}
                                    templateColumns="repeat(4, 1fr)"
                                    gap={6}
                                >
                                    {/* {courseList.map((course, index) => {
                                        return (
                                            <CourseCard
                                                key={index}
                                                courseTitle={course.courseTitle}
                                                courseClass={course.courseClass}
                                                courseTeacher={
                                                    course.courseTeacher
                                                }
                                                courseDescription={
                                                    course.courseDescription
                                                }
                                            />
                                        );
                                    })} */}
                                </Grid>
                            </Box>
                            <Flex
                                minWidth="max-content"
                                alignItems="center"
                                gap="2"
                                mt={4}
                            >
                                <Button
                                    leftIcon={<ArrowBackIcon />}
                                    colorScheme="blue"
                                >
                                    Previous
                                </Button>
                                <Spacer />
                                <Button
                                    rightIcon={<ArrowForwardIcon />}
                                    colorScheme="blue"
                                >
                                    Next
                                </Button>
                            </Flex>
                        </Stack>
                        {/* End Header */}
                    </Box>
                    {/* End Main */}
                </Flex>
            </Flex>
        </>
    );
}
