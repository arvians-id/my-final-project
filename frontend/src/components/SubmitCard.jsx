import React, { useState } from 'react';
import {
  Box,
  Text,
  FormControl,
  Stack,
  FormLabel,
  Flex,
  Icon,
  Button,
  chakra,
  VisuallyHidden,
  createStandaloneToast,
} from '@chakra-ui/react';
import { API_USER_SUBMIT_SUBMISSION } from '../api/submission';
import formatBytes from '../utils/formatBytes';
import { BiReset } from 'react-icons/bi';

export default function SubmitCard({ courseCode, submissionId, onSuccess }) {
  const [file, setFile] = useState();
  const [loading, setLoading] = useState(false);
  const { toast } = createStandaloneToast();

  const onChange = (e) => {
    setFile(e.target.files[0]);
  };

  const onSubmit = async () => {
    setLoading(true);
    const res = await API_USER_SUBMIT_SUBMISSION(
      courseCode,
      submissionId,
      file
    );
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Upload Tugas',
      });
      onSuccess();
      clearFile();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Upload Tugas',
      });
    }
    setLoading(false);
  };

  const clearFile = () => {
    setFile(undefined);
  };

  return (
    <>
      <Stack
        px={4}
        py={5}
        bg="blackAlpha.50"
        _dark={{
          bg: '#141517',
        }}
        spacing={6}
        p={{
          sm: 6,
        }}
        ml={5}
        borderRadius="10"
      >
        {file ? (
          <Flex>
            <Text w="full">
              {file.name} ({formatBytes(file.size)})
            </Text>
            <BiReset onClick={clearFile} />
          </Flex>
        ) : (
          <FormControl>
            <FormLabel
              fontSize="sm"
              fontWeight="md"
              color="gray.700"
              _dark={{
                color: 'gray.50',
              }}
            >
              Uploda File
            </FormLabel>
            <Flex
              mt={1}
              justify="center"
              px={6}
              pt={5}
              pb={6}
              borderWidth={2}
              _dark={{
                color: 'gray.500',
              }}
              borderStyle="dashed"
              rounded="md"
            >
              <Stack spacing={1} textAlign="center">
                <Icon
                  mx="auto"
                  boxSize={12}
                  color="gray.400"
                  _dark={{
                    color: 'gray.500',
                  }}
                  stroke="currentColor"
                  fill="none"
                  viewBox="0 0 48 48"
                  aria-hidden="true"
                >
                  <path
                    d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  />
                </Icon>
                <Flex
                  fontSize="sm"
                  color="gray.600"
                  _dark={{
                    color: 'gray.400',
                  }}
                  alignItems="baseline"
                >
                  <chakra.label
                    htmlFor="file-upload"
                    cursor="pointer"
                    rounded="md"
                    fontSize="md"
                    color="brand.600"
                    _dark={{
                      color: 'brand.200',
                    }}
                    pos="relative"
                    _hover={{
                      color: 'brand.400',
                      _dark: {
                        color: 'brand.300',
                      },
                    }}
                  >
                    <span
                      style={{
                        padding: 8,
                        backgroundColor: 'blue',
                        color: 'white',
                        fontSize: '20px',
                        fonteWight: '800',
                        bordeRadius: '16px',
                      }}
                    >
                      Select file
                    </span>
                    <VisuallyHidden>
                      <input
                        onChange={onChange}
                        id="file-upload"
                        name="file-upload"
                        type="file"
                      />
                    </VisuallyHidden>
                  </chakra.label>
                  {/* <Text pl={1}>or drag and drop</Text> */}
                </Flex>
                <Text
                  fontSize="xs"
                  color="gray.500"
                  _dark={{
                    color: 'gray.50',
                  }}
                >
                  File up to 10MB
                </Text>
              </Stack>
            </Flex>
          </FormControl>
        )}
      </Stack>
      <Box textAlign="right" mr={7}>
        <Button
          type="submit"
          colorScheme="blue"
          _focus={{
            shadow: '',
          }}
          fontWeight="md"
          onClick={onSubmit}
          disabled={!file}
        >
          Kirim Tugas
        </Button>
      </Box>
    </>
  );
}
