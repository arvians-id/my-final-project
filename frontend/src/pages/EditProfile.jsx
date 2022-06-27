import { React, useEffect, useState } from 'react';
import {
  Box,
  createStandaloneToast,
  Flex,
  Image,
  Spinner,
  Text,
  Textarea,
} from '@chakra-ui/react';
import MainAppLayout from '../components/layout/MainAppLayout';
import {
  API_GET_USER_DETAIL_BY_ID,
  API_UPDATE_USER_PROFILE_DETAIL,
} from '../api/user';
import useStore from '../provider/zustand/store';
import { checkIsValidUsername } from '../utils/user';
import {
  VStack,
  FormLabel,
  Input,
  Button,
  Select,
  InputGroup,
  InputLeftAddon,
} from '@chakra-ui/react';
import { adapterUserToFE } from '../utils/adapterToFE';

export default function EditProfile() {
  const [userDetail, setUserDetail] = useState();
  const [loadingGetUserDetail, setLoadingGetUserDetail] = useState(false);
  const [loadingSubmit, setLoadingSubmit] = useState(false);
  const user = useStore((state) => state.user);
  const { toast } = createStandaloneToast();
  const [formProfile, setFormProfile] = useState({
    name: '',
    username: '',
    // email: '',
    // password: '',
    // confirmPassword: '',
    gender: 0,
    type_of_disability: 0,
    role: 2,
    phone: '',
    address: '',
    birthdate: '',
    description: '',
    image: '',
  });
  const setUser = useStore((state) => state.setUser);

  const getUserDetail = async () => {
    setLoadingGetUserDetail(true);
    const res = await API_GET_USER_DETAIL_BY_ID(user.id);
    if (res.status === 200) {
      console.log('res.data.data', res.data.data);
      setUserDetail(res.data.data);
      setFormProfile({
        name: res.data.data.name,
        username: res.data.data.username,
        description: res.data.data.description,
        image: res.data.data.image,
        // email: res.data.data.email,
        // password: '',
        // confirmPassword: '',
        gender: res.data.data.gender,
        type_of_disability: res.data.data.type_of_disability,
        role: res.data.data.role,
        phone: res.data.data.phone,
        address: res.data.data.address,
        birthdate: res.data.data.birthdate,
      });
    }
    setLoadingGetUserDetail(false);
  };

  const onChangeForm = (e) => {
    setFormProfile({
      ...formProfile,
      [e.target.name]: e.target.value,
    });
  };

  const onChangeGender = (e) => {
    setFormProfile({
      ...formProfile,
      gender: Number(e.target.value),
    });
  };

  const onChangeDisabilitas = (e) => {
    setFormProfile({
      ...formProfile,
      type_of_disability: Number(e.target.value),
    });
  };

  const checkIsValidRegister = () => {
    if (
      !formProfile.name ||
      // !formProfile.email ||
      // !formProfile.password ||
      !formProfile.birthdate ||
      // !formProfile.password ||
      !formProfile.phone ||
      formProfile.gender === 0 ||
      formProfile.type_of_disability === 0 ||
      // formProfile.password !== formProfile.confirmPassword ||
      // !checkIsValidPassword() ||
      !checkIsValidUsername(formProfile.username)
    )
      return true;
    return false;
  };

  const renderUsernameStatus = () => {
    if (formProfile.username) {
      if (checkIsValidUsername(formProfile.username)) {
        return (
          <Box
            bgColor="green.600"
            p="1"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
          >
            <Text color="white">Username valid</Text>
          </Box>
        );
      } else {
        return (
          <Box
            bgColor="red.600"
            borderBottomLeftRadius="4px"
            borderBottomRightRadius="4px"
            p="1"
          >
            <Text color="white">
              Minimal 3 karakter terdiri dari huruf besar, kecil, angka,
              karakter . _
            </Text>
          </Box>
        );
      }
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoadingSubmit(true);
    console.log('formProfile', formProfile);
    const res = await API_UPDATE_USER_PROFILE_DETAIL(formProfile, user.id);
    if (res.status === 200) {
      toast({
        status: 'success',
        title: 'Berhasil',
        description: 'Berhasil Ubah Profile',
      });
      setUser(adapterUserToFE(res.data.data));
      getUserDetail();
    } else {
      toast({
        status: 'error',
        title: 'Gagal',
        description: 'Gagal Ubah Profile',
      });
    }
    setLoadingSubmit(false);
  };

  const clearForm = () => {
    setFormProfile({
      name: '',
      username: '',
      // email: '',
      // password: '',
      // confirmPassword: '',
      gender: 0,
      type_of_disability: 0,
      role: 2,
      phone: '',
      address: '',
      birthdate: '',
      description: '',
      image: '',
    });
  };

  useEffect(() => {
    getUserDetail();
  }, []);

  return (
    <MainAppLayout>
      <Flex width="700px" justifyContent="center" minHeight="90vh" bg="white">
        <Box w="full" m={5}>
          <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
            {loadingGetUserDetail ? (
              <Spinner />
            ) : (
              <Box>
                {userDetail ? (
                  <Box m={5}>
                    <VStack spacing={4} align="stretch">
                      <Box>
                        <FormLabel htmlFor="name" fontWeight="bold">
                          Full Name
                        </FormLabel>
                        <Input
                          id="name"
                          name="name"
                          type="text"
                          maxWidth="full"
                          height={50}
                          placeholder="Full Name"
                          value={formProfile.name}
                          onChange={onChangeForm}
                        />
                      </Box>
                      <Box>
                        <FormLabel htmlFor="username" fontWeight="bold">
                          Username
                        </FormLabel>
                        <Input
                          id="username"
                          name="username"
                          type="text"
                          maxWidth="full"
                          height={50}
                          placeholder="User Name"
                          value={formProfile.username}
                          onChange={onChangeForm}
                        />
                        {renderUsernameStatus()}
                      </Box>
                      <Box>
                        <FormLabel htmlFor="email" fontWeight="bold">
                          Gender
                        </FormLabel>
                        <Select
                          id="gender"
                          placeholder="Select Gender"
                          name="gender"
                          value={formProfile.gender}
                          onChange={onChangeGender}
                        >
                          <option value={1}>Pria</option>
                          <option value={2}>Wanita</option>
                        </Select>
                      </Box>
                      <Box>
                        <FormLabel
                          htmlFor="type_of_disability"
                          fontWeight="bold"
                        >
                          Disabilitas
                        </FormLabel>
                        <Select
                          id="type_of_disability"
                          name="type_of_disability"
                          placeholder="Select Disabilitas"
                          value={formProfile.type_of_disability}
                          onChange={onChangeDisabilitas}
                          required
                        >
                          <option value={3}>None</option>
                          <option value={1}>Tunanetra</option>
                          <option value={2}>Tunarungu</option>
                        </Select>
                      </Box>
                      <Box>
                        <FormLabel htmlFor="phone" fontWeight="bold">
                          Phone Number
                        </FormLabel>
                        <InputGroup mt={5}>
                          <InputLeftAddon children="+62" />
                          <Input
                            type="tel"
                            placeholder="phone number"
                            name="phone"
                            value={formProfile.phone}
                            onChange={onChangeForm}
                          />
                        </InputGroup>
                      </Box>
                      <Box>
                        <FormLabel htmlFor="birthdate" fontWeight="bold">
                          Tanggal Lahir
                        </FormLabel>
                        <InputGroup mt={5}>
                          <Input
                            type="date"
                            placeholder="Tanggal Lahir"
                            name="birthdate"
                            value={formProfile.birthdate}
                            onChange={onChangeForm}
                          />
                        </InputGroup>
                      </Box>
                      <Box>
                        <FormLabel htmlFor="description" fontWeight="bold">
                          Deskripsi
                        </FormLabel>
                        <Textarea
                          id="description"
                          name="description"
                          type="text"
                          maxWidth="full"
                          rows="4"
                          placeholder="Deskripsi"
                          value={formProfile.description}
                          onChange={onChangeForm}
                        />
                      </Box>
                      <Box>
                        <FormLabel htmlFor="address" fontWeight="bold">
                          Alamat Tempat Tinggal
                        </FormLabel>
                        <Textarea
                          id="address"
                          name="address"
                          type="text"
                          maxWidth="full"
                          rows="2"
                          placeholder="Deskripsi"
                          value={formProfile.address}
                          onChange={onChangeForm}
                        />
                      </Box>
                      <Box>
                        <FormLabel htmlFor="image" fontWeight="bold">
                          Profile image link
                        </FormLabel>
                        <Input
                          id="image"
                          name="image"
                          type="text"
                          maxWidth="full"
                          height={50}
                          placeholder="Profile image link"
                          value={formProfile.image}
                          onChange={onChangeForm}
                        />
                        <Image
                          src={formProfile.image}
                          alt={user.name}
                          w="100px"
                          h="100px"
                          // borderRadius="100%"
                          objectFit="cover"
                          objectPosition="center"
                          fallbackSrc="/user.png"
                        />
                      </Box>
                      <Box>
                        <VStack spacing={3} mt={5}>
                          <Button
                            disabled={checkIsValidRegister()}
                            onClick={handleSubmit}
                            colorScheme="blue"
                            loading={loadingSubmit}
                            variant="outline"
                            width="100%"
                            p={5}
                            type="submit"
                          >
                            Ubah
                          </Button>
                        </VStack>
                      </Box>
                    </VStack>
                  </Box>
                ) : (
                  <Text>Error</Text>
                )}
              </Box>
            )}
          </Box>
        </Box>
      </Flex>
    </MainAppLayout>
  );
}
