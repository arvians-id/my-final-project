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
  InputLeftElement,
  InputGroup,
  Input,
  Avatar,
  Select,
  createStandaloneToast,
} from '@chakra-ui/react';
import { MdStackedBarChart } from 'react-icons/md';
import { SearchIcon } from '@chakra-ui/icons';
import MainAppLayout from '../../components/layout/MainAppLayout';
import {
  API_ADD_USER_IN_COURSE,
  API_GET_ALL_COURSE,
  API_GET_LIST_USER_IN_COURSE,
  API_REMOVE_USER_IN_COURSE,
} from '../../api/course';
import { API_GET_LIST_USER } from '../../api/user';

export default function AdminAddCourseStudent() {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [listCourse, setListCourse] = useState([]);
  const [selectedCourse, setSelectedCourse] = useState();
  const [listCourseFull, setListCourseFull] = useState([]);
  const [selectedIdCourse, setSelectedIdCourse] = useState();
  const [listUser, setListUser] = useState([]);
  const [listAllUser, setListAllUser] = useState([]);
  const [selectedUserToAdd, setSelectedUserToAdd] = useState();
  const [selectedUserIdToAdd, setSelectedUserIdToAdd] = useState();

  const { toast } = createStandaloneToast();
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

  let studentList = [
    {
      id: 1,
      name: 'Irfan Kurniawan',
    },
    {
      id: 2,
      name: 'Amellisa Anggun Oktadilla',
    },
  ];

  const getListCourse = async () => {
    const res = await API_GET_ALL_COURSE();
    if (res.status === 200) {
      const data = res.data.data ?? [];
      setListCourseFull(data);
      let result = [];
      for (const course of data) {
        result.push({
          label: `${course.name} - ${course.class}`,
          value: course.id,
        });
      }
      setListCourse(result);
    }
  };

  const onChangeCourse = (e) => {
    setSelectedIdCourse(e.target.value);
    const selected = listCourseFull.filter(
      (course) => course.id === Number(e.target.value)
    )[0];
    setSelectedCourse(selected);
    getListUser(selected.code_course);
  };

  const onChangeSelectUser = (e) => {
    setSelectedUserIdToAdd(e.target.value);
    const selected = selectedUserToAdd.filter(
      (user) => user.id === e.target.value
    )[0];
    setSelectedUserToAdd(selected);
  };

  const getListUser = async (courseCode) => {
    const res = await API_GET_LIST_USER_IN_COURSE(courseCode);
    if (res.status === 200) {
      setListUser(res.data.data ?? []);
    }
  };

  const handleRemoveUser = async (userId) => {
    const res = await API_REMOVE_USER_IN_COURSE(userId, selectedIdCourse);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Mengeluarkan Siswa Dalam Course',
      });
      const selected = listCourseFull.filter(
        (course) => course.id === Number(selectedIdCourse)
      )[0];
      getListUser(selected.code_course);
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Mengeluarkan Siswa Dalam Course',
      });
    }
  };

  const getAllUsers = async () => {
    const res = await API_GET_LIST_USER();
    if (res.status === 200) {
      setListAllUser(res.data.data ?? []);
    }
  };

  const onAddSiswa = async () => {
    const res = await API_ADD_USER_IN_COURSE({
      user_id: Number(selectedUserIdToAdd),
      course_id: Number(selectedIdCourse),
    });
    if (res.status === 200 || res.status === 201) {
      const selected = listCourseFull.filter(
        (course) => course.id === Number(selectedIdCourse)
      )[0];
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Tambahkan Siswa Kedalam Course',
      });
      getListUser(selected.code_course);
      onClose();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Tambahkan Siswa Kedalam Course',
      });
    }
  };

  const filterListAllUser = () => {
    const listIdUserExist = listUser.map((user) => user.id_user);
    return listAllUser.filter((user) => !listIdUserExist.includes(user.id));
  };

  useEffect(() => {
    getListCourse();
    getAllUsers();
  }, []);

  return (
    <>
      <MainAppLayout>
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
                Daftar Siswa
              </Box>
              <Box>
                <Text as="h2" fontSize="xl" fontWeight="semibold">
                  Pilih Course
                </Text>
                <Select
                  id="course"
                  name="course"
                  placeholder="Select Course"
                  value={selectedIdCourse}
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
                    variant="solid"
                    colorScheme="green"
                    width="30%"
                    onClick={onOpen}
                  >
                    Tambah Siswa
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
                          <Th>Nama Siswa</Th>
                          <Th>Aksi</Th>
                        </Tr>
                      </Thead>
                      {listUser.length > 0 ? (
                        <Tbody>
                          {listUser.map((student, index) => {
                            return (
                              <Tr key={index}>
                                <Td>{num++}</Td>
                                <Td>{student.user_username}</Td>
                                <Td>
                                  <Stack direction="row" spacing={3}>
                                    <Button
                                      variant="solid"
                                      colorScheme="red"
                                      size="sm"
                                      onClick={() =>
                                        handleRemoveUser(student.id_user)
                                      }
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
                        <Text>Belum Ada Siswa Dalam Course</Text>
                      )}
                    </Table>
                  </TableContainer>
                </Box>
              )}
              {/* End Content */}
            </Stack>
          </Box>
        </Flex>
        {/* End main */}

        {/* Modal */}
        <Modal isOpen={isOpen} onClose={onClose}>
          <ModalOverlay />
          <ModalContent>
            <ModalHeader>Tambah Siswa</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              <Select
                id="user"
                name="user"
                placeholder="Pilih Siswa"
                value={selectedUserIdToAdd}
                onChange={onChangeSelectUser}
                required
              >
                {filterListAllUser().map((user, index) => (
                  <option value={user.id} key={index}>
                    {user.username}
                  </option>
                ))}
              </Select>
              {/* <Box>
                <InputGroup>
                  <InputLeftElement
                    pointerEvents="none"
                    children={<SearchIcon color="gray.300" />}
                  />
                  <Input type="tel" placeholder="Cari" />
                </InputGroup>
              </Box>
              <Box
                my={4}
                p={2}
                border="2px"
                borderColor="gray.200"
                borderRadius={5}
              >
                listAllUser
                <Stack direction="row" alignItems="center">
                  <Avatar name="Irfan Kurniawan" width={10} height={10} />
                  <Text>Irfan Kurniawan</Text>
                </Stack>
              </Box> */}
            </ModalBody>

            <ModalFooter>
              <Button colorScheme="blue" mr={3} onClick={onAddSiswa}>
                Tambah Data
              </Button>
              <Button onClick={onClose}>Cancel</Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
        {/* End Modal */}
      </MainAppLayout>
    </>
  );
}
