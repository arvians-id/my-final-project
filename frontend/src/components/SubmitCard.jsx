import React from 'react';
import {
  Box,
  SimpleGrid,
  GridItem,
  Text,
  FormControl,
  Stack,
  FormLabel,
  InputGroup,
  InputLeftAddon,
  Input,
  Flex,
  Textarea,
  FormHelperText,
  Icon,
  Button,
  chakra,
  VisuallyHidden,
} from '@chakra-ui/react';

export default function SubmitCard() {
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
        <SimpleGrid columns={3} spacing={6}>
          <FormControl as={GridItem} colSpan={[3, 2]}>
            <FormLabel
              fontSize="sm"
              fontWeight="md"
              color="gray.700"
              _dark={{
                color: 'gray.50',
              }}
            >
              Link
            </FormLabel>
            <InputGroup size="sm">
              <InputLeftAddon
                bg="gray.50"
                _dark={{
                  bg: 'gray.800',
                }}
                color="gray.500"
                rounded="md"
              >
                http://
              </InputLeftAddon>
              <Input type="tel" focusBorderColor="brand.400" rounded="md" />
            </InputGroup>
          </FormControl>
        </SimpleGrid>

        <div>
          <FormControl id="email" mt={1}>
            <FormLabel
              fontSize="sm"
              fontWeight="md"
              color="gray.700"
              _dark={{
                color: 'gray.50',
              }}
            >
              About
            </FormLabel>
            <Textarea
              mt={1}
              rows={3}
              shadow="sm"
              focusBorderColor="brand.400"
              fontSize={{
                sm: 'sm',
              }}
            />
            <FormHelperText>Jelaskan Pengumpulan</FormHelperText>
          </FormControl>
        </div>
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
                  <span>Upload file</span>
                  <VisuallyHidden>
                    <input id="file-upload" name="file-upload" type="file" />
                  </VisuallyHidden>
                </chakra.label>
                <Text pl={1}>or drag and drop</Text>
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
      </Stack>
      <Box textAlign="right" mr={7}>
        <Button
          type="submit"
          colorScheme="blue"
          _focus={{
            shadow: '',
          }}
          fontWeight="md"
        >
          Save
        </Button>
      </Box>
    </>
  );
}
