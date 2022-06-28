import React, { useState } from 'react';
import {
  Box,
  Flex,
  Stack,
  HStack,
  Text,
  OrderedList,
  ListItem,
  UnorderedList,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import DetailCard from '../components/DetailCard';
import ListModule from '../components/ListModule';

export default function CourseDetail() {
  const [iterate, setIterate] = useState(1);
  let moduleDetail = [
    {
      id: 1,
      name: 'Pemrograman Web',
      class: 'XI RPL',
      description:
        'Pemrograman web terbentuk atas 2 kata yaitu pemrograman dan web dimana pemrograman sendiri adalah Proses atau Cara dalam menjalankan sebuah urutan intruksi atau perintah yang diberikan kepada komputer untuk membuat fungsi atau tugas tertentu. dan Web adalah Sistem untuk mengakses, memanipulasi, dan mengunduh dokumen yang terdapat pada komputer yang di hubungkan melalui internet atau jaringan.Jadi Pemrograman Web adalah Proses atau Cara untuk menjalankan intruksi pada sebuah komputer yang terhubung ke internet untuk membuat fungsi atau tugas tertentu.',
    },
  ];
  let listModule = [
    {
      id: 1,
      name: 'Pemrograman Web 1',
      list: 'HTML',
    },
    {
      id: 2,
      name: 'Pemrograman Web 2',
      list: 'CSS',
    },
    {
      id: 3,
      name: 'Pemrograman Web 3',
      list: 'Layout Responsif Menggunakan Flexbox',
    },
  ];

  const handleClick = () => {
    setIterate(iterate + 1);
  };

  return (
    <>
      {/* Navbar */}
      <Navbar />
      <Flex
        direction="row"
        justifyContent="flex-start"
        alignItems="flex-start"
        top="30"
      >
        {/* Sidebar */}
        <Flex
          width="20%"
          minHeight="100vh"
          bgColor="grey.100"
          boxShadow="md"
          position="fixed"
          left="0"
          top="20"
          overflowY="auto"
        >
          <Sidebar />
        </Flex>
        {/* End Sidebar */}
        <Flex
          width="80%"
          minHeight="90vh"
          bg="white"
          position="sticky"
          left="80"
          marginTop={20}
        >
          <Box m={5}>
            <Stack spacing={6}>
              <Box alignContent="flex-start">
                <HStack spacing={20}>
                  <Box>
                    {moduleDetail.map((module, index) => {
                      return (
                        <DetailCard
                          key={index}
                          name={module.name}
                          className={module.class}
                          description={module.description}
                        />
                      );
                    })}
                  </Box>
                  <Box
                    p={6}
                    ml={4}
                    shadow="md"
                    borderWidth="1px"
                    w={550}
                    h={350}
                    borderRadius={10}
                  >
                    <Stack mt="4" spacing={3}>
                      <Text fontWeight="semibold" fontSize="xl" mt={4}>
                        Modul :
                      </Text>
                      {listModule.map((submission, index) => {
                        return (
                          <ListModule
                            key={index}
                            name={submission.name}
                            list={submission.list}
                            onClick={handleClick}
                          />
                        );
                      })}
                    </Stack>
                  </Box>
                </HStack>
              </Box>
              <Box
                p={6}
                m={4}
                shadow="md"
                borderWidth="1px"
                w={950}
                h={750}
                borderRadius={10}
              >
                <Text fontWeight="semibold" fontSize="xl" mt={4}>
                  Tentang Pelajaran :
                </Text>
                <Text mt={4} align="justify">
                  Berdasarkan struktur katanya, maka pemrograman web terdiri
                  dari dua kata, yaitu pemrograman yang artinya adalah
                  sekumpulan perintah yang diciptakan oleh manusia agar bisa
                  membantu manusia lainnya untuk menghasilkan program. Sementara
                  itu web diartikan sebagai sumber informasi yang dapat diakses
                  hanya dengan menggunakan jaringan komputer yang terhubung
                  dengan internet. Bentuk informasi yang berasal dari web itu
                  bermacam-macam, mulai dari teks, gambar, audio, video hingga
                  animasi. Jadi, bisa disimpulkan bahwa pemrograman web adalah
                  instruksi untuk dapat menghasilkan program atau situs web yang
                  bisa ditampilkan dengan menggunakan browser melalui jaringan
                  internet. Teknologi web generasi ketiga ini sebenarnya masih
                  dalam perdebatan, sebab pengertian dari teknologi web ini
                  masih beragam, ada yang berpendapat bahwa teknologi web 3.0
                  ini adalah layanan akses broadband secara mobile hingga
                  layanan web yang isinya adalah perangkat lunak dengan sifat
                  on-demand. Teknologi web 3.0 ini juga ada yang menyebutnya
                  dengan istilah semantic web ini juga dikenal sebagai teknologi
                  web yang bukan hanya memiliki isi web yang dapat dimengerti
                  manusia namun juga dapat diinterpretasikan oleh software
                  sehingga proses pengintegrasian informasi akan terasa lebih
                  mudah. Ada sejumlah standar yang digunakan untuk membangun
                  semantic web, diantaranya seperti XML, XML Schema, RDF, OWL
                  dan SPAROL.
                </Text>
                <Text fontWeight="semibold" fontSize="xl" mt={4}>
                  Yang akan Dipelajari :
                </Text>
                <OrderedList>
                  <ListItem>
                    Di akhir pelatihan, peserta dapat membuat sebuah website
                    sederhana menggunakan kode pemrograman yang sesuai standar
                    global.
                  </ListItem>
                  <ListItem>
                    Membangun website menggunakan kode HTML, CSS, dan JavaScript
                    sederhana.
                  </ListItem>
                  <ListItem>
                    Menerapkan struktur website yang baik menggunakan standar
                    semantic HTML.{' '}
                  </ListItem>
                  <ListItem>
                    Mendemonstrasikan penyusunan layout website menggunakan
                    teknik float atau flexbox.
                  </ListItem>
                </OrderedList>
                <Text fontWeight="semibold" fontSize="xl" mt={4}>
                  Target Pembelajaran :
                </Text>
                <UnorderedList>
                  <ListItem>
                    Kelas ditujukan bagi pemula yang ingin memulai karirnya di
                    bidang web development (pembuatan web) dan membutuhkan dasar
                    atau fondasi yang kuat sebelum belajar lebih dalam di bidang
                    web
                  </ListItem>
                  <ListItem>
                    Kelas dapat diikuti oleh siswa yang melek IT sehingga wajib
                    memiliki dan dapat mengoperasikan komputer dengan baik.
                  </ListItem>
                  <ListItem>
                    Kelas ini didesain untuk pemula sehingga tidak ada prasyarat
                    dalam pemahaman pemrograman sebelumnya.{' '}
                  </ListItem>
                  <ListItem>
                    Siswa harus bisa belajar mandiri, berkomitmen, benar-benar
                    punya rasa ingin tahu, dan tertarik pada subjek materi,{' '}
                  </ListItem>
                </UnorderedList>
              </Box>
            </Stack>
          </Box>
        </Flex>
      </Flex>
    </>
  );
}
