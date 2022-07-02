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
  createStandaloneToast,
} from '@chakra-ui/react';
import { MdStackedBarChart } from 'react-icons/md';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_GET_ALL_COURSE } from '../../api/course';
import { axiosWithToken } from '../../api/axiosWithToken';
import { BASE_URL } from '../../constant/api';

import fileDownload from 'js-file-download';
export default function AdminCourseSubmissionList() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [listCourse, setListCourse] = useState([]);
  const [selectedCourse, setSelectedCourse] = useState();
  const [listCourseFull, setListCourseFull] = useState([]);
  const [listSubmission, setListSubmission] = useState([]);
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();
  const [selectedIdSubmission, setSelectedIdSubmission] = useState(0);
  const [listSiswaAssignment, setListSiswaAssignment] = useState([]);
  const [selectedSiswaToGrade, setSelectedSiswaToGrade] = useState();
  const [valueGrade, setValueGrade] = useState(0);
  const { toast } = createStandaloneToast();
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
    if (score) {
      return score;
    } else {
      return 'Belum Dinilai';
    }
  };

  const onChangeValueGrade = (e) => {
    setValueGrade(e.target.value);
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
    getListSubmission(e.target.value);
  };

  const onChangeSubmission = (e) => {
    setSelectedIdSubmission(Number(e.target.value));
    const selected = listSubmission.filter(
      (course) => course.code_course === Number(e.target.value)
    )[0];
    getListSiswaAssignSubmission(Number(e.target.value));
  };

  const getListSubmission = async (courseCode) => {
    axiosWithToken()
      .get(`${BASE_URL}/api/courses/${courseCode}/submissions`)
      .then((res) => {
        if (res.status === 200) {
          setListSubmission(res.data.data);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const getListSiswaAssignSubmission = async (submissionId) => {
    axiosWithToken()
      .get(
        `${BASE_URL}/api/courses/${selectedCodeCourse}/submissions/${submissionId}/get`
      )
      .then((res) => {
        if (res.status === 200) {
          setListSiswaAssignment(res.data.data);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const downloadSiswaAssign = async (filename, userSubmissionId) => {
    axiosWithToken()
      .post(
        `${BASE_URL}/api/courses/${selectedCodeCourse}/submissions/${selectedIdSubmission}/user-submit/${userSubmissionId}/download`,
        {
          responseType: 'blob',
        }
      )
      .then((res) => {
        if (res.status === 200) {
          fileDownload(res.data, `${filename}.pdf`);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const onOpenDialog = (idUserSubmission) => {
    const selectedSiswa = listSiswaAssignment.filter(
      (siswa) => siswa.id_user_submission === Number(idUserSubmission)
    )[0];
    setSelectedSiswaToGrade(selectedSiswa);
    onOpen();
  };

  const onCloseDialog = () => {
    setSelectedSiswaToGrade(undefined);
    onClose();
  };

  const handleGradeSiswa = async () => {
    axiosWithToken()
      .patch(
        `${BASE_URL}/api/courses/${selectedCodeCourse}/submissions/${selectedIdSubmission}/user-submit/${selectedSiswaToGrade.id_user_submission}`,
        {
          grade: Number(valueGrade),
        }
      )
      .then((res) => {
        if (res.status === 200) {
          toast({
            status: 'success',
            title: 'Berhasil',
            description: 'Berhasil memberikan nilai siswa',
          });
          getListSiswaAssignSubmission(selectedIdSubmission);
          setValueGrade(0);
          onCloseDialog();
        } else {
          toast({
            status: 'error',
            title: 'Gagal',
            description: 'Gagal memberikan nilai siswa',
          });
        }
      })
      .catch((err) => {
        console.log(err);
      });
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
                <Stack my={4}>
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

              {selectedCourse && (
                <Box my={4}>
                  <Text as="h2" fontSize="xl" fontWeight="semibold">
                    Pilih Submission
                  </Text>
                  <Select
                    id="submission"
                    name="submission"
                    placeholder="Select Submission"
                    value={selectedIdSubmission}
                    onChange={onChangeSubmission}
                    required
                  >
                    {listSubmission.map((course, index) => (
                      <option value={course.id} key={index}>
                        {course.name}
                      </option>
                    ))}
                  </Select>
                </Box>
              )}
            </Box>
            {/* End Header */}
            {/* Content */}
            {selectedIdSubmission && (
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
                    {listSiswaAssignment.length > 0 ? (
                      <Tbody>
                        {listSiswaAssignment.map((assign, index) => {
                          return (
                            <Tr key={index}>
                              <Td>{index + 1}</Td>
                              <Td>{assign.user_name}</Td>
                              <Td>{new Date().toString()}</Td>
                              <Td>{scoreStatus(assign.grade)}</Td>
                              <Td>
                                <Stack direction="row" spacing={3}>
                                  <Button
                                    variant="solid"
                                    colorScheme="green"
                                    size="sm"
                                    onClick={() =>
                                      onOpenDialog(assign.id_user_submission)
                                    }
                                  >
                                    Nilai
                                  </Button>
                                  <Button
                                    variant="solid"
                                    colorScheme="blue"
                                    size="sm"
                                    onClick={() =>
                                      downloadSiswaAssign(
                                        `${assign.user_name}-${assign.module_submission_name}`,
                                        assign.id_user_submission
                                      )
                                    }
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
                    ) : (
                      <Text>Tidak ada siswa</Text>
                    )}
                  </Table>
                </TableContainer>
              </Box>
            )}{' '}
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
      {/* End main */}
      {/* Modal */}
      <Modal isOpen={isOpen} onClose={onCloseDialog}>
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
                value={selectedSiswaToGrade?.user_name}
                disabled
              />
              <Text as="h3" fontSize="md" fontWeight="semibold">
                Nilai
              </Text>
              <Input
                type="number"
                value={valueGrade}
                onChange={onChangeValueGrade}
                placeholder="Masukkan nilai"
              />
            </Stack>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={handleGradeSiswa}>
              Berikan Nilai
            </Button>
            <Button variant="ghost" onClick={onCloseDialog}>
              Cancel
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
      {/* End Modal */}
    </MainAppLayout>
  );
}
