import React from 'react';
import { Box, Flex, Stack, HStack, Text, Spacer, Button, Grid, MenuButton, MenuList, MenuItem, Menu } from '@chakra-ui/react';
import { ChevronDownIcon, ArrowBackIcon, ArrowForwardIcon } from '@chakra-ui/icons';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import CourseCard from '../components/CourseCard';

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
    <>
      {/* Navbar */}
      <Navbar />
      <Flex direction="row" justifyContent="flex-start" alignItems="flex-start" top="30">
        {/* Sidebar */}
        <Flex width="20%" minHeight="100vh" bgColor="grey.100" boxShadow="md" position="fixed" left="0" top="20" overflowY="auto">
          <Sidebar />
        </Flex>
        {/* End Sidebar */}
        {/* Main */}
        <Flex width="80%" minHeight="90vh" bg="white" position="sticky" left="80" marginTop={20}>
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
                  <MenuButton px={4} py={2} transition="all 0.2s" borderRadius="md" borderWidth="1px" _hover={{ bg: 'gray.400' }} _expanded={{ bg: 'blue.400' }} _focus={{ boxShadow: 'outline' }}>
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
                    return <CourseCard key={index} name={module.name} className={module.class} description={module.description} percent={module.percent} />;
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
        </Flex>
        {/* End main */}
      </Flex>
    </>
  );
}
