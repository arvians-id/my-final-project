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
  useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Input,
  Spinner,
  createStandaloneToast,
  Textarea,
} from '@chakra-ui/react';
import MainAppLayout from '../../components/layout/MainAppLayout';
import {
  API_CREATE_COURSE,
  API_DELETE_COURSE,
  API_GET_ALL_COURSE,
  API_UPDATE_COURSE,
} from '../../api/course';
import ModalCustom from '../../components/ModalCustom';

export default function AdminCourseList() {
  const [loadingGetCourse, setLoadingGetCourse] = useState(false);
  const [loadingAddCourse, setLoadingAddCourse] = useState(false);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [listCourse, setListCourse] = useState([]);
  const { toast } = createStandaloneToast();
  const [formAddCourse, setFormAddCourse] = useState({
    name: '',
    class: '',
    tools: '', // longtext
    about: '', // longtext
    description: '', // longtext
  });

  const getAllCourse = async () => {
    setLoadingGetCourse(true);
    const res = await API_GET_ALL_COURSE();
    if (res.status === 200) {
      setListCourse(res.data.data ?? []);
    }
    setLoadingGetCourse(false);
  };

  const handleSubmitCourse = async () => {
    setLoadingAddCourse(true);
    const res = await API_CREATE_COURSE(formAddCourse);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Memambahkan Course',
      });
      getAllCourse();
      clearForm();
      onClose();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Menambahkan Course',
      });
    }
    setLoadingAddCourse(false);
  };

  const changeFormAddCourse = (e) => {
    setFormAddCourse({
      ...formAddCourse,
      [e.target.name]: e.target.value,
    });
  };

  const clearForm = () => {
    setFormAddCourse({
      name: '',
      class: '',
      tools: '', // longtext
      about: '', // longtext
      description: '', // longtext
    });
  };

  useEffect(() => {
    getAllCourse();
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
            <Box>
              <Box as="h1" fontSize="2xl" fontWeight="semibold">
                Manajemen Mata Pelajaran
              </Box>
              <Box as="span" fontSize="l" fontWeight="semibold" color="grey">
                Kelola Data Mata Pelajaran Anda
              </Box>
            </Box>
            {/* End Header */}
            {/* Content */}
            <Button
              variant="solid"
              colorScheme="green"
              width="30%"
              onClick={onOpen}
            >
              Tambah Mata Pelajaran
            </Button>
            <Box>
              <TableContainer>
                <Table variant="striped" colorScheme="blue">
                  <Thead>
                    <Tr>
                      <Th>No</Th>
                      <Th>Mata Pelajaran</Th>
                      <Th>Kelas</Th>
                      <Th>Aksi</Th>
                    </Tr>
                  </Thead>
                  <Tbody>
                    {loadingGetCourse ? (
                      <Spinner />
                    ) : listCourse.length === 0 ? (
                      <Text>Tidak ada course</Text>
                    ) : (
                      listCourse.map((course, index) => {
                        return (
                          <CourseListItem
                            key={index}
                            getAllCourse={getAllCourse}
                            index={index}
                            course={course}
                          />
                        );
                      })
                    )}
                  </Tbody>
                </Table>
              </TableContainer>
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
      {/* End main */}
      <ModalFormCourse
        name={formAddCourse.name}
        nameClass={formAddCourse.class}
        tools={formAddCourse.tools}
        about={formAddCourse.about}
        description={formAddCourse.description}
        isOpen={isOpen}
        onClose={onClose}
        onSubmit={handleSubmitCourse}
        onChange={changeFormAddCourse}
        loading={loadingAddCourse}
      />
    </MainAppLayout>
  );
}

const CourseListItem = ({
  getAllCourse,
  index,
  course,
  // name,
  // className,
  // code_course,
}) => {
  const [loadingDeleteCourse, setLoadingDeleteCourse] = useState(false);
  const [loadingUpdateCourse, setLoadingUpdateCourse] = useState(false);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const dialogDelete = useDisclosure();
  const { toast } = createStandaloneToast();

  const [formUpdateCourse, setFormUpdateCourse] = useState({
    name: course.name,
    class: course.class,
    tools: course.tools, // longtext
    about: course.about, // longtext
    description: course.description, // longtext
  });

  const changeFormUpdateCourse = (e) => {
    setFormUpdateCourse({
      ...formUpdateCourse,
      [e.target.name]: e.target.value,
    });
  };

  const clearForm = () => {
    setFormUpdateCourse({
      name: '',
      class: '',
      tools: '', // longtext
      about: '', // longtext
      description: '', // longtext
    });
  };

  const handleDeleteCourse = async () => {
    setLoadingDeleteCourse(true);
    const res = await API_DELETE_COURSE(course.code_course);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Hapus Course',
      });
      dialogDelete.onClose();
      getAllCourse();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Hapus Course',
      });
    }
    setLoadingDeleteCourse(false);
  };

  const onDelete = (code) => {
    dialogDelete.onOpen();
  };

  const handleUpdateCourse = async () => {
    setLoadingUpdateCourse(true);
    const res = await API_UPDATE_COURSE(course.code_course, formUpdateCourse);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Memambahkan Course',
      });
      getAllCourse();
      clearForm();
      onClose();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Menambahkan Course',
      });
    }
    setLoadingUpdateCourse(false);
  };

  return (
    <>
      <Tr key={index}>
        <Td>{index + 1}</Td>
        <Td>{course.name}</Td>
        <Td>{course.class}</Td>
        <Td>
          <Stack direction="row" spacing={3}>
            <Button variant="solid" colorScheme="teal" size="sm">
              Detail
            </Button>
            <Button
              variant="solid"
              colorScheme="blue"
              size="sm"
              onClick={onOpen}
            >
              Edit
            </Button>
            <Button
              variant="solid"
              colorScheme="red"
              size="sm"
              onClick={onDelete}
              disabled={loadingDeleteCourse}
            >
              Hapus
            </Button>
          </Stack>
        </Td>
      </Tr>

      {/* Modal */}
      <ModalCustom
        isOpen={dialogDelete.isOpen}
        onClose={dialogDelete.onClose}
        modalBody={<Text>Apakah Yakin Untuk Menghapus ? </Text>}
        modalFooter={
          <>
            <Button colorScheme="green" mr={3} onClick={dialogDelete.onClose}>
              Tidak
            </Button>
            <Button colorScheme="red" onClick={handleDeleteCourse}>
              Oke
            </Button>
          </>
        }
      />
      {/* End Modal */}
      <ModalFormCourse
        name={formUpdateCourse.name}
        nameClass={formUpdateCourse.class}
        tools={formUpdateCourse.tools}
        about={formUpdateCourse.about}
        description={formUpdateCourse.description}
        isOpen={isOpen}
        onClose={onClose}
        onSubmit={handleUpdateCourse}
        onChange={changeFormUpdateCourse}
        loading={loadingUpdateCourse}
      />
    </>
  );
};

const ModalFormCourse = ({
  isOpen,
  onClose,
  onSubmit,
  onChange,
  name,
  nameClass,
  tools,
  about,
  description,
  loading,
}) => {
  return (
    <Modal isOpen={isOpen} closeOnOverlayClick={false} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Tambah Data Mata Pelajaran</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <Stack direction="column">
            <Text as="h3" fontSize="md" fontWeight="semibold">
              Nama Mata Pelajaran
            </Text>
            <Input
              value={name}
              name="name"
              onChange={onChange}
              placeholder="Masukkan Nama Mata Pelajaran"
            />
            <Text as="h3" fontSize="md" fontWeight="semibold">
              Kelas
            </Text>
            <Input
              value={nameClass}
              name="class"
              onChange={onChange}
              placeholder="Masukkan Kelas"
            />
            <Text as="h3" fontSize="md" fontWeight="semibold">
              Tools
            </Text>
            <Input
              value={tools}
              name="tools"
              onChange={onChange}
              placeholder="Masukkan Tools"
            />
            <Text as="h3" fontSize="md" fontWeight="semibold">
              About
            </Text>
            <Textarea
              value={about}
              name="about"
              onChange={onChange}
              placeholder="Masukkan Tentang Course"
            />
            <Text as="h3" fontSize="md" fontWeight="semibold">
              Description
            </Text>
            <Textarea
              value={description}
              name="description"
              onChange={onChange}
              placeholder="Masukkan Deskripsi Course"
            />
          </Stack>
        </ModalBody>
        <ModalFooter>
          <Button
            disabled={!name || !nameClass || loading}
            colorScheme="blue"
            isLoading={loading}
            mr={3}
            onClick={onSubmit}
          >
            Tambah Mata Pelajaran
          </Button>
          <Button variant="ghost" onClick={onClose}>
            Cancel
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
