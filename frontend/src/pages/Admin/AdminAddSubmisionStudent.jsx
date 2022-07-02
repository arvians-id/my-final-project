import React, { useEffect, useRef, useState } from 'react';
import {
  Box,
  Flex,
  Stack,
  Text,
  Button,
  Input,
  createStandaloneToast,
  Select,
  FormLabel,
  Textarea,
} from '@chakra-ui/react';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_CREATE_MODULE_ARTICLE } from '../../api/moduleArticles';
import { API_GET_ALL_COURSE } from '../../api/course';
import { useNavigate } from 'react-router';
import { axiosWithToken } from '../../api/axiosWithToken';
import { BASE_URL } from '../../constant/api';

export default function AdminAddSubmisionStudent() {
  const [loadingSubmit, setLoadingSubmit] = useState(false);
  const [formAddSubmission, setFormAddSubmission] = useState({
    name: '',
    description: '',
    deadline: '',
  });
  const { toast } = createStandaloneToast();
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();
  const [listCourse, setListCourse] = useState([]);
  const navigate = useNavigate();

  const submit = async (e) => {
    e.preventDefault();
    setLoadingSubmit(true);
    let content = '';
    axiosWithToken()
      .post(
        `${BASE_URL}/api/courses/${selectedCodeCourse}/submissions`,
        formAddSubmission
      )
      .then((res) => {
        if (res.status === 200) {
          toast({
            status: 'success',
            title: 'Berhasil',
            description: 'Berhasil Menambahkan',
          });
          clearForm();
        } else {
          toast({
            status: 'error',
            title: 'Gagal',
            description: 'Gagal Menambahkan',
          });
        }
      })
      .catch((err) => {
        toast({
          status: 'error',
          title: 'Gagal',
          description: 'Gagal Menambahkan',
        });
      });
    setLoadingSubmit(false);
  };

  const handleChangeForm = (e) => {
    setFormAddSubmission({
      ...formAddSubmission,
      [e.target.name]: e.target.value,
    });
  };

  const getListCourse = async () => {
    const res = await API_GET_ALL_COURSE();
    if (res.status === 200) {
      const data = res.data.data ?? [];
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

  const clearForm = () => {
    setFormAddSubmission({
      name: '',
      description: '',
      deadline: '',
    });
  };

  const onChangeCourse = (e) => {
    setSelectedCodeCourse(e.target.value);
  };

  useEffect(() => {
    getListCourse();
  }, []);

  return (
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
              Tambah Materi
            </Box>
            {/* End Header */}
            {/* Content */}
            <Box>
              <Stack direction="column" spacing={3}>
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
                {selectedCodeCourse && (
                  <>
                    <FormLabel as="h2" fontSize="xl" fontWeight="semibold">
                      Nama Tugas
                    </FormLabel>
                    <Input
                      onChange={handleChangeForm}
                      placeholder="Masukkan Judul Materi Anda Di Sini"
                      name="name"
                      value={formAddSubmission.name}
                    />
                    <FormLabel as="h2" fontSize="xl" fontWeight="semibold">
                      Soal
                    </FormLabel>
                    <Textarea
                      onChange={handleChangeForm}
                      placeholder="Masukkan Judul Materi Anda Di Sini"
                      name="description"
                      value={formAddSubmission.description}
                    />
                    <FormLabel as="h2" fontSize="xl" fontWeight="semibold">
                      Batas Pengumpulan
                    </FormLabel>
                    <Input
                      type="date"
                      placeholder="Deadline"
                      name="deadline"
                      value={formAddSubmission.deadline}
                      onChange={handleChangeForm}
                    />

                    <Button
                      disabled={
                        !formAddSubmission.name ||
                        !formAddSubmission.description ||
                        !formAddSubmission.deadline ||
                        loadingSubmit
                      }
                      variant="solid"
                      colorScheme="green"
                      width="30%"
                      onClick={submit}
                      type="submit"
                    >
                      Tambahkan
                    </Button>
                  </>
                )}
              </Stack>
            </Box>
            {/* End Content */}
          </Stack>
        </Box>
      </Flex>
    </MainAppLayout>
  );
}
