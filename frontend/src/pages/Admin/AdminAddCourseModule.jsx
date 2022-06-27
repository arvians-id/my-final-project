import React, { useRef } from 'react';
import {
    Box,
    Flex,
    Stack,
    Text,
    Button,
    Input

} from '@chakra-ui/react'
import Navbar from '../components/Navbar'
import Sidebar from '../components/Sidebar'
import { MdStackedBarChart } from 'react-icons/md'
import { Editor } from '@tinymce/tinymce-react';

export default function AdminAddCourseModule() {

    const editorRef = useRef(null);
    const log = () => {
        if (editorRef.current) {
            console.log(editorRef.current.getContent());
        }
    };


    return (
        <>
            {/* Navbar */}
            <Navbar />
            <Flex direction="row" justifyContent="flex-start" alignItems="flex-start" top="30">
                {/* Sidebar */}
                <Flex width="20%" minHeight="100vh" bgColor="grey.100" boxShadow='md' position="fixed" left="0" top="20" overflowY="auto">
                    <Sidebar />
                </Flex>
                {/* End Sidebar */}
                {/* Main */}
                <Flex width="80%" minHeight="90vh" bg="white" position="sticky" left="80" marginTop={20}>
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
                                    <Text as="h2" fontSize="xl" fontWeight='semibold'>Judul Materi</Text>
                                    <Input placeholder='Masukkan Judul Materi Anda Di Sini' />
                                    <Text as="h2" fontSize="xl" fontWeight='semibold'>Isi Materi</Text>
                                    <Editor
                                        apiKey='hjoe212eg1j17e47oon829pkkrldrjd0pgxy67rc98fpflgd'
                                        onInit={(evt, editor) => editorRef.current = editor}
                                        initialValue="<p>Silahkan Tambahkan Materi Anda Di Sini.</p>"
                                        init={{
                                            height: 500,
                                            menubar: false,
                                            plugins: [
                                                'advlist autolink lists link image charmap print preview anchor',
                                                'searchreplace visualblocks code fullscreen',
                                                'insertdatetime media table paste code help wordcount'
                                            ],
                                            toolbar: 'undo redo | styles | formatselect | ' +
                                                'bold italic backcolor | alignleft aligncenter ' +
                                                'alignright alignjustify | bullist numlist outdent indent | ' +
                                                'removeformat | help',
                                            content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }'
                                        }}
                                    />
                                    <Button variant="solid" colorScheme="green" width="30%">Tambahkan Materi</Button>
                                </Stack>
                            </Box>
                            {/* End Content */}
                        </Stack>
                    </Box>
                </Flex>
                {/* End main */}
            </Flex>
        </ >
    )
}

