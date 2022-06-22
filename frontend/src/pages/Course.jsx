<<<<<<< HEAD
import React from 'react';
import {
    Button,
    InputGroup,
    InputLeftElement,
    Input,
    FormLabel,
    MenuButton,
    MenuList,
    Menu,
    MenuItem,
    Spacer,
    VStack,
    Box,
    Flex,
    HStack,
    SimpleGrid,
} from '@chakra-ui/react';
import { SearchIcon, ChevronDownIcon } from '@chakra-ui/icons';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
export default function Course() {
    return (
        <div>
            <Navbar />

            <Flex direction="row" justifyContent="center" alignItems="center">
                <Flex
                    width="20%"
                    minHeight="200vh"
                    bgColor="grey.100"
                    boxShadow="dark-lg"
                >
                    <Sidebar />
                </Flex>
                <Box width="80%" minHeight="200vh" bg="white">
                    <Box m={5}>
                        <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
                            Course
                        </Box>
                    </Box>
                    <Spacer />
                    <Box mt="80px" ml="1150px">
                        <Menu>
                            <MenuButton
                                as={Button}
                                rightIcon={<ChevronDownIcon />}
                            >
                                Kelas
                            </MenuButton>
                            <MenuList>
                                <MenuItem>1 SMA</MenuItem>
                                <MenuItem>2 SMA</MenuItem>
                                <MenuItem>3 SMA</MenuItem>
                            </MenuList>
                        </Menu>
                    </Box>
                    <Box>
                        <SimpleGrid
                            bg="white"
                            columns={{ sm: 4, md: 4 }}
                            spacing="8"
                            p="10"
                            textAlign="center"
                            rounded="lg"
                            color="gray.400"
                        >
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="md"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="md"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                            <Box
                                boxSize="250px"
                                color="black"
                                as="button"
                                boxShadow="md"
                                p="6"
                                rounded="lg"
                                bg="gray.100"
                            >
                                Bahasa Indonesia Kelas 3 SMA
                            </Box>
                        </SimpleGrid>
                    </Box>
                </Box>
            </Flex>
        </div>
    );
=======
import {
  ArrowBackIcon,
  ArrowForwardIcon,
  ChevronDownIcon,
} from '@chakra-ui/icons';
import {
  Box,
  Button,
  Flex,
  Grid,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Spacer,
  Stack,
} from '@chakra-ui/react';
import React from 'react';
import CourseCard from '../components/CourseCard';
import MainAppLayout from '../components/layout/MainAppLayout';

export default function Course() {
  let moduleList = [
    {
      id: 1,
      name: 'Pemrograman Web',
      class: 'XI RPL',
      description: 'tentang Web Programming',
      percent: 70,
    },
    {
      id: 2,
      name: 'Bahasa Indonesia',
      class: 'X TKJ',
      description: 'Bahasa Indonesia Pelajaran Mengenai Bahasa Indonesia',
      percent: 85,
    },
    {
      id: 3,
      name: 'Matematika',
      class: 'XII TKJ',
      description: 'Mata Pelajaran yang akan membahas Konversi Biner, Aljabar ',
      percent: 60,
    },
    {
      id: 4,
      name: 'Geografi',
      class: 'XI IPS',
      description: 'Mempelajari tentang Struktur Bumi',
      percent: 60,
    },
    {
      id: 5,
      name: 'Matematika',
      class: 'XII IPA',
      description: 'Mata Pelajaran yang akan membahas Konversi Biner, Aljabar ',
      percent: 60,
    },
    {
      id: 6,
      name: 'Sosiologi',
      class: 'XI IPS',
      description: 'Mata Pelajaran yang akan membahas Stratifikasi Sosial ',
      percent: 60,
    },
    {
      id: 7,
      name: 'Ekonomi',
      class: 'XII IPS',
      description: 'Mata Pelajaran yang akan membahas Akuntansi',
      percent: 60,
    },
    {
      id: 8,
      name: 'Biologi',
      class: 'XII IPA',
      description: 'Mata Pelajaran yang akan membahas Virus dan Bakteri ',
      percent: 60,
    },
    {
      id: 9,
      name: 'Fisika',
      class: 'XII IPA',
      description: 'Mata Pelajaran yang akan membahas Hukum Pascal',
      percent: 60,
    },
    {
      id: 10,
      name: 'Kimia',
      class: 'XII IPA',
      description: 'Mata Pelajaran yang akan membahas Biomolekul ',
      percent: 60,
    },
    {
      id: 11,
      name: 'Agama',
      class: 'XII IPA/IPS',
      description: 'Mata Pelajaran yang akan membahas Biomolekul ',
      percent: 60,
    },
  ];

  return (
    <MainAppLayout>
      <Box m={5}>
        <Stack spacing={6}>
          {/* Header */}
          <Box>
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Selamat Datang Kembali
            </Box>
            <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
              Lanjutkan Pembelajaran Anda
            </Box>
          </Box>
          {/* End Header */}
          {/* Content */}
          <Flex justifyContent="baseline">
            <Box as="h1" fontSize="2xl" fontWeight="semibold">
              Pelajaran Anda
            </Box>
            <Spacer />
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
          </Flex>
          <Box alignContent="flex-start">
            <Grid spacing={8} templateColumns="repeat(3, 1fr)" gap={6}>
              {moduleList.map((module, index) => {
                return (
                  <CourseCard
                    key={index}
                    name={module.name}
                    className={module.class}
                    description={module.description}
                    percent={module.percent}
                  />
                );
              })}
            </Grid>
          </Box>
          {/* End Content */}
        </Stack>
        <Flex minWidth="max-content" alignItems="center" gap="2" mt={4}>
          <Button leftIcon={<ArrowBackIcon />} colorScheme="blue">
            Previous
          </Button>
          <Spacer />
          <Button rightIcon={<ArrowForwardIcon />} colorScheme="blue">
            Next
          </Button>
        </Flex>
      </Box>
    </MainAppLayout>
  );
>>>>>>> 28ee5ed6f3b932b186ee81144b50e15402a23589
}
