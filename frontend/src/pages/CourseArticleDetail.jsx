import React, { useEffect, useState } from 'react';
import {
  Box,
  Flex,
  Stack,
  HStack,
  Text,
  OrderedList,
  ListItem,
  UnorderedList,
  Spinner,
} from '@chakra-ui/react';
import DetailCard from '../components/DetailCard';
import ListModule from '../components/ListModule';
import MainAppLayout from '../components/layout/MainAppLayout';
import { useParams } from 'react-router';
import { API_GET_COURSE_BY_CODE } from '../api/course';
import {
  API_GET_ALL_ARTICLE_BY_COURSE_CODE,
  API_GET_ARTICLE_DETAIL,
} from '../api/moduleArticles';

export default function CourseArticleDetail() {
  const [loadingGetDetailCourse, setLoadingGetDetailCourse] = useState(false);
  const [detailCourse, setDetailCourse] = useState();
  const [listArticle, setListArticle] = useState([]);
  const [articleDetail, setArticleDetail] = useState();

  const { courseCode, articleId } = useParams();

  const getCourseDetail = async () => {
    setLoadingGetDetailCourse(true);
    const res = await API_GET_COURSE_BY_CODE(courseCode);
    if (res.status === 200) {
      setDetailCourse(res.data.data);
    }
    setLoadingGetDetailCourse(false);
  };

  const getListArticle = async () => {
    const res = await API_GET_ALL_ARTICLE_BY_COURSE_CODE(courseCode);
    if (res.status === 200) {
      setListArticle(res.data.data ?? []);
    }
  };

  const getDetailArticle = async () => {
    const res = await API_GET_ARTICLE_DETAIL(courseCode, articleId);
    if (res.status === 200) {
      setArticleDetail(res.data.data);
    }
  };

  useEffect(() => {
    getCourseDetail();
    getListArticle();
    getDetailArticle();
  }, []);

  useEffect(() => {
    getDetailArticle();
  }, [articleId]);

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

  return (
    <MainAppLayout>
      <Flex
        width="full"
        minHeight="90vh"
        bg="white"
        position="sticky"
        left="80"
        marginTop={20}
      >
        {loadingGetDetailCourse ? (
          <Spinner />
        ) : detailCourse ? (
          <Flex m={5} gap="10px">
            <Stack w="75%" spacing={6}>
              <Box w="full" alignContent="flex-start">
                <HStack spacing={20} w="full">
                  <Box position="relative" w="full">
                    <DetailCard
                      name={detailCourse.name}
                      className={detailCourse.class}
                      description={detailCourse.description}
                    />
                    {detailCourse.is_active === false && (
                      <Box
                        position="absolute"
                        backgroundColor="gray.100"
                        padding={4}
                        opacity={0.3}
                        top="50%"
                        left="10%"
                        transform="rotate(-10deg)"
                      >
                        <Text
                          fontSize="50px"
                          fontWeight="extrabold"
                          color="red"
                        >
                          Kelas Sedang Tidak Aktif
                        </Text>
                      </Box>
                    )}
                  </Box>
                </HStack>
              </Box>
              {articleDetail ? (
                <Box
                  p={6}
                  m={4}
                  shadow="md"
                  borderWidth="1px"
                  // w={950}
                  // h={750}
                  borderRadius={10}
                  w="full"
                >
                  <Text fontWeight="semibold" fontSize="xl" mt={4}>
                    Judul : {articleDetail.name}
                  </Text>

                  <div
                    dangerouslySetInnerHTML={{ __html: articleDetail.content }}
                  />
                </Box>
              ) : (
                <Text>Gagal Mendapatkan Detail Article Course</Text>
              )}

              {detailCourse.is_active === false && (
                <Box
                  position="absolute"
                  backgroundColor="gray.100"
                  padding={4}
                  opacity={0.3}
                  top="50%"
                  left="10%"
                  transform="rotate(-10deg)"
                >
                  <Text fontSize="50px" fontWeight="extrabold" color="red">
                    Kelas Sedang Tidak Aktif
                  </Text>
                </Box>
              )}
            </Stack>
            <Box
              p={6}
              ml={4}
              shadow="md"
              borderWidth="1px"
              h="fit-content"
              borderRadius={10}
              w="25%"
            >
              <Stack mt="4" spacing={3}>
                <Text fontWeight="semibold" fontSize="xl" mt={4}>
                  Modul :
                </Text>
                {listArticle.map((article, index) => {
                  return (
                    <ListModule
                      key={index}
                      courseCode={courseCode}
                      article={article}
                    />
                  );
                })}
              </Stack>
            </Box>
          </Flex>
        ) : (
          <Text>Gagal Mendapatkan Course</Text>
        )}
      </Flex>
    </MainAppLayout>
  );
}
