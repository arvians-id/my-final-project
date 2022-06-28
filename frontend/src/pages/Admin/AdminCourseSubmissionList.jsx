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
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Input,
  Select,
} from '@chakra-ui/react';
import { MdStackedBarChart } from 'react-icons/md';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_GET_ALL_COURSE } from '../../api/course';

export default function AdminCourseSubmissionList() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [listCourse, setListCourse] = useState([]);
  const [selectedCourse, setSelectedCourse] = useState();
  const [listCourseFull, setListCourseFull] = useState([]);
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();

  let num = 1;
  let courseModule = [
    {
      id: 1,
      name: 'Bahasa Indonesia',
      class: 'XI Bahasa',
      submissionTitle: 'Tugas Bahasa Indonesia 1',
    },
  ];

  let studentSubmission = [
    {
      id: 1,
      name: 'Irfan Kurniawan',
      file: 'tugas-1-irfan',
      time: '2022-06-22T15:23:51.141Z',
      score: 80,
    },
    {
      id: 2,
      name: 'Rahmalina',
      file: 'tugas-1-rahmalina',
      time: '2022-06-22T15:23:51.141Z',
      score: null,
    },
  ];

  let scoreStatus = (score) => {
    if (score !== null) {
      return score;
    } else {
      return 'Belum Dinilai';
    }
  };

  const getListCourse = async () => {
    const res = await API_GET_ALL_COURSE();
    if (res.status === 200) {
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

  const onChangeCourse = (e) => {
    setSelectedCodeCourse(e.target.value);
    const selected = listCourseFull.filter(
      (course) => course.code_course === e.target.value
    )[0];
    setSelectedCourse(selected);
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
              Tugas Siswa
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
            <Box>
              <TableContainer>
                <Table variant="striped" colorScheme="blue">
                  <Thead>
                    <Tr>
                      <Th>No</Th>
                      <Th>Nama Siswa</Th>
                      <Th>Waktu Pengumpulan</Th>
                      <Th>Nilai</Th>
                      <Th>Aksi</Th>
                    </Tr>
                  </Thead>
                  {
                    <Tbody>
                      {studentSubmission.map((submission) => {
                        return (
                          <Tr>
                            <Td>{num++}</Td>
                            <Td>{submission.name}</Td>
                            <Td>{submission.time}</Td>
                            <Td>{scoreStatus(submission.score)}</Td>
                            <Td>
                              <Stack direction="row" spacing={3}>
                                <Button
                                  variant="solid"
                                  colorScheme="green"
                                  size="sm"
                                  onClick={onOpen}
                                >
                                  Nilai
                                </Button>
                                <Button
                                  variant="solid"
                                  colorScheme="blue"
                                  size="sm"
                                >
                                  Dowload
                                </Button>
                                <Button
                                  variant="solid"
                                  colorScheme="red"
                                  size="sm"
                                >
                                  Hapus
                                </Button>
                              </Stack>
                            </Td>
                          </Tr>
                        );
                      })}
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
          <ModalHeader>Penilaian Tugas Siswa</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Stack direction="column">
              <Text as="h3" fontSize="md" fontWeight="semibold">
                Nama Siswa
              </Text>
              <Input
                placeholder="Masukkan Nama Mata Pelajaran"
                value="Irfan Kurniawan"
                disabled
              />
              <Text as="h3" fontSize="md" fontWeight="semibold">
                Nilai
              </Text>
              <Input type="number" placeholder="Masukkan nilai" />
            </Stack>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={onClose}>
              Berikan Nilai
            </Button>
            <Button variant="ghost" onClick={onClose}>
              Cancel
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
      {/* End Modal */}
    </MainAppLayout>
  );
}
