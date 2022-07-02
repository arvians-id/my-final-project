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
} from '@chakra-ui/react';
import { MdStackedBarChart } from 'react-icons/md';
import { Editor } from '@tinymce/tinymce-react';
import MainAppLayout from '../../components/layout/MainAppLayout';
import { API_CREATE_MODULE_ARTICLE } from '../../api/moduleArticles';
import { API_GET_ALL_COURSE } from '../../api/course';
import { useNavigate, useParams } from 'react-router';
import { useSearchParams } from 'react-router-dom';
import { axiosWithToken } from '../../api/axiosWithToken';
import { BASE_URL } from '../../constant/api';

export default function AdminAddCourseModule() {
  const editorRef = useRef(null);
  const [loadingSubmit, setLoadingSubmit] = useState(false);
  const [formModuleAricle, setFormModuleArticle] = useState({
    name: '',
    content: '',
    estimate: 0,
  });
  const { toast } = createStandaloneToast();
  const [selectedCodeCourse, setSelectedCodeCourse] = useState();
  const [listCourse, setListCourse] = useState([]);
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [actionType, setActionType] = useState('');
  const [initialValue, setInitialValue] = useState(
    '<p>Silahkan Tambahkan Materi Anda Di Sini.</p>'
  );

  const log = () => {
    if (editorRef.current) {
      // console.log(editorRef.current.getContent());
    }
  };

  const submit = async (e) => {
    e.preventDefault();
    setLoadingSubmit(true);
    if (actionType !== 'edit') {
      // add
      let content = '';
      if (editorRef.current) {
        content = editorRef.current.getContent();
      }
      const data = {
        name: formModuleAricle.name,
        content: content,
        estimate: formModuleAricle.estimate,
      };
      const res = await API_CREATE_MODULE_ARTICLE(selectedCodeCourse, data);
      if (res.status === 200) {
        toast({
          status: 'success',
          title: 'Berhasil',
          description: 'Berhasil Menambahkan Materi',
        });
        navigate('/admin-course-detail');
      } else {
        toast({
          status: 'error',
          title: 'Gagal',
          description: 'Gagal  Menambahkan ',
        });
      }
    } else {
      // edit
      let content = '';
      if (editorRef.current) {
        content = editorRef.current.getContent();
      }
      const data = {
        name: formModuleAricle.name,
        content: content,
        estimate: formModuleAricle.estimate,
      };
      const res = await API_CREATE_MODULE_ARTICLE(selectedCodeCourse, data);
      if (res.status === 200) {
        toast({
          status: 'success',
          title: 'Berhasil',
          description: 'Berhasil Ubah Materi',
        });
        navigate('/admin-course-detail');
      } else {
        toast({
          status: 'error',
          title: 'Gagal',
          description: 'Gagal Ubah',
        });
      }
    }
    setLoadingSubmit(false);
  };

  const handleChangeTitle = (e) => {
    setFormModuleArticle({
      ...formModuleAricle,
      name: e.target.value,
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
    setFormModuleArticle({
      name: '',
      content: '',
      estimate: 0,
    });
  };

  const onChangeCourse = (e) => {
    setSelectedCodeCourse(e.target.value);
  };

  const getDetailModuleArticle = async (courseCode, articleId) => {
    axiosWithToken()
      .get(`${BASE_URL}/api/courses/${courseCode}/articles/${articleId}`)
      .then((res) => {
        if (res.status === 200) {
          setFormModuleArticle({
            name: res.data.data.name,
            content: res.data.data.content,
            estimate: res.data.data.estimate,
          });
          setInitialValue(res.data.data.content);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    getListCourse();
  }, []);

  useEffect(() => {
    const codeCourse = searchParams.get('codeCourse');
    const articleId = searchParams.get('moduleId');
    if (codeCourse && articleId && searchParams.get('action') === 'edit') {
      setSelectedCodeCourse(searchParams.get('codeCourse'));
      setActionType('edit');
      getDetailModuleArticle(codeCourse, articleId);
    }
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
                    <Text as="h2" fontSize="xl" fontWeight="semibold">
                      Judul Materi
                    </Text>
                    <Input
                      onChange={handleChangeTitle}
                      placeholder="Masukkan Judul Materi Anda Di Sini"
                      value={formModuleAricle.name}
                      name="name"
                    />
                    <Text as="h2" fontSize="xl" fontWeight="semibold">
                      Isi Materi
                    </Text>
                    <Editor
                      apiKey="hjoe212eg1j17e47oon829pkkrldrjd0pgxy67rc98fpflgd"
                      onInit={(evt, editor) => (editorRef.current = editor)}
                      initialValue={initialValue}
                      init={{
                        height: 500,
                        menubar: false,
                        plugins: [
                          'advlist autolink lists link image charmap print preview anchor',
                          'searchreplace visualblocks code fullscreen',
                          'insertdatetime media table paste code help wordcount',
                        ],
                        toolbar:
                          'undo redo | styles | formatselect | ' +
                          'bold italic backcolor | alignleft aligncenter ' +
                          'alignright alignjustify | bullist numlist outdent indent | ' +
                          'removeformat | help',
                        content_style:
                          'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }',
                      }}
                    />
                    <Button
                      disabled={!formModuleAricle.name || loadingSubmit}
                      variant="solid"
                      colorScheme="green"
                      width="30%"
                      onClick={submit}
                      type="submit"
                    >
                      {actionType === 'edit'
                        ? 'Ubah Materi'
                        : 'Tambahkan Materi'}
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
