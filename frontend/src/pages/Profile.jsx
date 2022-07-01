import { React, useEffect, useState } from 'react';
import {
  Box,
  Flex,
  HStack,
  Image,
  Spinner,
  Text,
  VStack,
} from '@chakra-ui/react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import MainAppLayout from '../components/layout/MainAppLayout';
import { API_GET_USER_DETAIL_BY_ID } from '../api/user';
import useStore from '../provider/zustand/store';
import { GiTeacher } from 'react-icons/gi';
import { FaUserTie } from 'react-icons/fa';
import { BsTelephone } from 'react-icons/bs';
import { RiWomenFill } from 'react-icons/ri';
import { CgGenderMale } from 'react-icons/cg';
import { BsCalendarDate } from 'react-icons/bs';
import { FaHome } from 'react-icons/fa';
import { BsFillInfoCircleFill } from 'react-icons/bs';

export default function Profile() {
  const [userDetail, setUserDetail] = useState();
  const [loadingGetUserDetail, setLoadingGetUserDetail] = useState(false);
  const user = useStore((state) => state.user);

  const getUserDetail = async () => {
    setLoadingGetUserDetail(true);
    const res = await API_GET_USER_DETAIL_BY_ID(user.id);
    if (res.status === 200) {
      setUserDetail(res.data.data);
    }
    setLoadingGetUserDetail(false);
  };

  useEffect(() => {
    getUserDetail();
  }, []);

  return (
    <MainAppLayout>
      <Flex minHeight="90vh" bg="white">
        <Box m={5}>
          <Box as="h1" fontSize="xl" fontWeight="semibold" mb={2}>
            {loadingGetUserDetail ? (
              <Spinner />
            ) : (
              <Box>
                {userDetail ? (
                  <Flex gap="40px">
                    {/* profile image */}
                    <Box>
                      <Image
                        src={userDetail.image}
                        alt={userDetail.username}
                        w="230px"
                        h="230px"
                        borderRadius="100%"
                        objectFit="cover"
                        objectPosition="center"
                        fallbackSrc="/user.png"
                      />
                    </Box>
                    {/* user info */}
                    <Box>
                      <Text mb="1" fontSize="x-large">
                        {userDetail.name}
                      </Text>
                      <Text mb="1" fontSize="20px">
                        @{userDetail.username}
                      </Text>
                      <HStack mb="1">
                        {userDetail.role === 2 ? (
                          <FaUserTie />
                        ) : userDetail.role === 1 ? (
                          <GiTeacher />
                        ) : (
                          ''
                        )}
                        <Text fontSize="16px">
                          {userDetail.role === 2
                            ? 'Siswa'
                            : userDetail.role === 1
                            ? 'Guru'
                            : ''}
                        </Text>
                      </HStack>
                      <HStack mb="1">
                        <BsTelephone />
                        <Text fontSize="16px">{userDetail.phone}</Text>
                      </HStack>
                      <HStack mb="1">
                        {userDetail.gender ? <RiWomenFill /> : <CgGenderMale />}
                        <Text fontSize="16px">
                          {userDetail.gender ? 'Pria' : 'Wanita'}
                        </Text>
                      </HStack>
                      <HStack mb="1">
                        <BsCalendarDate />
                        <Text fontSize="16px">
                          {new Date(userDetail.birthdate).toLocaleDateString(
                            'id'
                          )}
                        </Text>
                      </HStack>
                      <HStack mb="1">
                        <FaHome />
                        <Text fontSize="16px">{userDetail.address}</Text>
                      </HStack>
                      <HStack mb="1">
                        <BsFillInfoCircleFill />
                        <Text fontSize="16px">{userDetail.description}</Text>
                      </HStack>
                    </Box>
                  </Flex>
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
