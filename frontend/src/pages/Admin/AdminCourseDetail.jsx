import React, { useEffect, useState } from 'react';
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
  Select,
  createStandaloneToast,
} from '@chakra-ui/react';
import { MdStackedBarChart } from 'react-icons/md';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_GET_ALL_COURSE } from '../../api/course';
import { Link, useNavigate } from 'react-router-dom';
import {
  API_DELETE_MODULE_ARTICLES,
  API_GET_ALL_ARTICLE_BY_COURSE_CODE,
} from '../../api/moduleArticles';

export default function AdminCourseDetail() {
  let num = 1;
  let courseDetail = [
    {
      id: 1,
      name: 'Bahasa Indonesia',
      class: 'XI Bahasa',
      description:
        'Pada Kelas Ini akan Mempelajari tentang Menyusun Prosedur, Teks Eksplanasi, Mengelola Informasi dan lainnya',
    },
  ];

  let moduleList = [
    {
      id: 1,
      title: 'Pengenalan Teks Prosedur',
      content:
        'A. Mengonstruksi Informasi dalam Teks Prosedur \n Menunjukkan Pernyataan Umum dalam Suatu Kegiatan \n  Seseorang melakukan suatu kegiatan tentu saja harus memperhatikan langkah-langkah mengerjakannya. Apabila kita akan melakukan pekerjaan, maka harus memahami langkah-langkahnya agar hasil kegiatan tersebut berhasil dengan baik. Ciri teks prosedur yaitu terdapat bagian pernyataan umum dan tahapan-tahapan melakukan kegiatan',
    },
  ];

  const [listCourse, setListCourse] = useState([]);
  const [selectedCourse, setSelectedCourse] = useState();
  const [listCourseFull, setListCourseFull] = useState([]);
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();
  const [listArticle, setListArticle] = useState([]);
  const navigate = useNavigate();
  const { toast } = createStandaloneToast();
  const getListCourse = async () => {
    const res = await API_GET_ALL_COURSE();
    if (res.status === 200) {
      console.log('res.data.data', res.data.data);
      const data = res.data.data ?? [];
      setListCourseFull(data);
      let result = [];
      for (const course of data) {
        result.push({
          label: `${course.name} - ${course.class}`,
          value: course.code_course,
        });
      }
      setListCourse(result);
    }
  };

  const getListArticle = async (courseCode) => {
    const res = await API_GET_ALL_ARTICLE_BY_COURSE_CODE(courseCode);
    if (res.status === 200) {
      setListArticle(res.data.data ?? []);
    }
  };

  const onChangeCourse = (e) => {
    setSelectedCodeCourse(e.target.value);
    const selected = listCourseFull.filter(
      (course) => course.code_course === e.target.value
    )[0];
    setSelectedCourse(selected);
    getListArticle(e.target.value);
  };

  const handleDeleteArticle = async (id) => {
    const res = await API_DELETE_MODULE_ARTICLES(selectedCodeCourse, id);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Menghapus Module',
      });
      getListArticle(selectedCodeCourse);
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Menghapus Module',
      });
    }
  };

  useEffect(() => {
    getListCourse();
  }, []);

  return (
    <MainAppLayout>
      {/* Main */}
      <Flex
        width="80%"
        minHeight="90vh"
        bg="white"
        position="sticky"
        left="80"
        marginTop={20}
      >
        <Box m={5} width="full">
          <Stack spacing={6}>
            {/* Header */}
            <Box as="h1" fontSize="3xl" fontWeight="semibold">
              Detail Materi
            </Box>
            <Box>
              <Text as="h2" fontSize="xl" fontWeight="semibold">
                Pilih Course
              </Text>
              <Select
                id="course"
                name="course"
                placeholder="Select Course"
                value={selectedCodeCourse}
                onChange={onChangeCourse}
                required
              >
                {listCourse.map((course, index) => (
                  <option value={course.value} key={index}>
                    {course.label}
                  </option>
                ))}
              </Select>
              {selectedCourse && (
                <Stack>
                  <Box as="h1" fontSize="2xl" fontWeight="semibold">
                    {selectedCourse.name}
                  </Box>
                  <Box
                    as="span"
                    fontSize="l"
                    fontWeight="semibold"
                    color="grey"
                  >
                    {selectedCourse.class}
                  </Box>
                  <Box
                    as="span"
                    fontSize="l"
                    fontWeight="semibold"
                    color="grey"
                  >
                    {selectedCourse.description}
                  </Box>
                </Stack>
              )}
            </Box>
            {/* End Header */}
            {/* Content */}
            {selectedCourse && (
              <Stack direction="row">
                <Button
                  onClick={() => navigate('/add-student-to-course')}
                  variant="solid"
                  colorScheme="green"
                  width="30%"
                >
                  Tambah Siswa
                </Button>
                <Button
                  onClick={() => navigate('/add-course')}
                  variant="solid"
                  colorScheme="green"
                  width="30%"
                >
                  Tambah Materi
                </Button>
              </Stack>
            )}
            {selectedCourse && (
              <Box>
                <TableContainer>
                  <Table variant="striped" colorScheme="blue">
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
                        {listArticle.length > 0 ? (
                          listArticle.map((module, index) => {
                            return (
                              <Tr key={index}>
                                <Td>{index + 1}</Td>
                                <Td>{module.name}</Td>
                                <Td>{`${module.content.substring(0, 30)}`}</Td>
                                <Td>
                                  <Stack direction="row" spacing={3}>
                                    <Button
                                      variant="solid"
                                      colorScheme="blue"
                                      size="sm"
                                    >
                                      Edit
                                    </Button>
                                    <Button
                                      variant="solid"
                                      colorScheme="red"
                                      size="sm"
                                      onClick={() =>
                                        handleDeleteArticle(module.id)
                                      }
                                    >
                                      Hapus
                                    </Button>
                                  </Stack>
                                </Td>
                              </Tr>
                            );
                          })
                        ) : (
                          <Box>
                            <Text>Module Belum Ada</Text>
                            <Link to="/add-course">
                              <Button colorScheme="green">Buat Module</Button>
                            </Link>
                          </Box>
                        )}
                      </Tbody>
                    }
                  </Table>
                </TableContainer>
              </Box>
            )}
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
      {/* End main */}
    </MainAppLayout>
  );
}
